package service

import (
	"context"

	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserGRPCService struct {
	rep *repository.Repository
}

func NewUserGRPCService(r *repository.Repository) *UserGRPCService {
	return &UserGRPCService{rep: r}
}

func (s *UserGRPCService) CreateUser(ctx context.Context, u *api.User) (*api.UserResponse, error) {
	// check for cache

	_, err := s.rep.CreateUser(ctx, u)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func (s *UserGRPCService) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	return nil, nil
}

func (s *UserGRPCService) DeleteUser(ctx context.Context, e *api.UserEmail) (*api.UserResponse, error) {
	return nil, nil
}
