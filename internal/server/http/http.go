package http

import (
	"context"
	"go-learning-server/internal/service"
	"go-learning-server/pkg/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Ping(ctx context.Context) (err error)
	Close()
	GetUser(context.Context, *service.GetUserReq) (*service.GetUserRes, error)
}

type APIServer struct {
	httpServer *http.Server
	Handler    Handler
}

func NewAPIServer(h Handler) *APIServer {
	r := gin.New()

	r.Use(Default())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	cfg := config.Get()

	a := &APIServer{
		httpServer: &http.Server{
			Addr:    cfg.Addr,
			Handler: r,
		},
		Handler: h,
	}

	r.GET("/user", a.GetUser)

	return a
}

func Default() gin.HandlerFunc {
	conf := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowAllOrigins = true
	return cors.New(conf)
}

func (s *APIServer) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *APIServer) Stop(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

type GetUserRequest struct {
	ID int64 `json:"id"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}

func (a *APIServer) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// DTO to DO
	res, err := a.Handler.GetUser(c, &service.GetUserReq{ID: req.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, GetUserResponse{Name: res.Name})
}
