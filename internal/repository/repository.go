package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/s3rzh/go-grpc-user-service/internal/repository/postgresql"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserRepository interface {
	CreateUser(context.Context, *api.User) (*api.UserResponse, error)
	GetUsers(context.Context) (*api.UsersResponse, error)
	DeleteUser(context.Context, *api.UserEmail) (*api.UserResponse, error)
}

type Repository struct {
	UserRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		postgresql.NewUserPostgres(db),
	}
}
