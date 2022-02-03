package postgresql

import (
	"context"

	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserPostgres struct {
}

func NewUserPostgres() *UserPostgres {
	return &UserPostgres{}
}

func (s *UserPostgres) CreateUser(ctx context.Context, u *api.User) (*api.UserResponse, error) {
	return nil, nil
}

func (s *UserPostgres) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	return nil, nil
}

func (s *UserPostgres) DeleteUser(ctx context.Context, e *api.UserEmail) (*api.UserResponse, error) {
	return nil, nil
}
