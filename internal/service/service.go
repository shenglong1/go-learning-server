package service

import (
	"context"
)

type DAOInterface interface {
	Close()
	Ping(ctx context.Context) (err error)
	GetUser(context.Context, *GetUserReq) (*GetUserRes, error)
}

// Service service.
type Service struct {
	dao DAOInterface
}

// New new a service and return.
func NewService(dao DAOInterface) *Service {
	return &Service{
		dao: dao,
	}
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}

type GetUserReq struct {
	ID int64 `json:"id"`
}

type GetUserRes struct {
	Name string `json:"name"`
}

func (s *Service) GetUser(ctx context.Context, r *GetUserReq) (*GetUserRes, error) {
	return s.dao.GetUser(ctx, r)
}
