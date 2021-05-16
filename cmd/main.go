package main

import (
	"context"
	"flag"
	"fmt"
	"go-learning-server/internal/dao"
	"go-learning-server/internal/server/http"
	"go-learning-server/internal/service"
	"go-learning-server/pkg/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	config.Set(conf)
	fmt.Println("go-learning-server start")

	apiServer := InitializeAllInstance()

	go apiServer.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			defer cancel()

			fmt.Println("go-learning-server exit")
			apiServer.Stop(ctx)
			time.Sleep(time.Second)
			return

		case syscall.SIGHUP:

		default:
			return
		}
	}
}

func InitializeAllInstance() *http.APIServer {
	// TODO: use wire
	return http.NewAPIServer(service.NewService(dao.NewDao()))
}
