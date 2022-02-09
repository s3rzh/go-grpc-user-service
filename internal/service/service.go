package service

import (
	"context"

	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
	"github.com/s3rzh/go-grpc-user-service/pkg/cache"
)

type UserService interface {
	CreateUser(context.Context, *api.User) (int, error)
	GetUsers(context.Context) (*api.UsersResponse, error)
	DeleteUser(context.Context, *api.UserEmail) (int, error)
}

type Service struct {
	UserService
}

func NewService(rep *repository.Repository, cache cache.Cache) *Service {
	return &Service{
		UserService: NewUserGRPCService(rep, cache),
	}
}
