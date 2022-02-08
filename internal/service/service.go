package service

import (
	"context"

	"github.com/s3rzh/go-grpc-user-service/internal/repository"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserService interface {
	CreateUser(context.Context, *api.User) (int, error)
	GetUsers(context.Context) (*api.UsersResponse, error)
	DeleteUser(context.Context, *api.UserEmail) (int, error)
}

type Service struct {
	UserService
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		UserService: NewUserGRPCService(rep),
	}
}
