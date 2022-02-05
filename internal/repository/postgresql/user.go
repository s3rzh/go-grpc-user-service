package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
	return &UserPostgres{db: db}
}

func (s *UserPostgres) CreateUser(ctx context.Context, u *api.User) (*api.UserResponse, error) {
	age := 22
	email := "123@g.com"
	var id int

	row := s.db.QueryRow(context.Background(),
		"INSERT INTO users (age, email) VALUES ($1, $2) RETURNING id",
		age, email)

	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	fmt.Println(id)

	return nil, nil
}

func (s *UserPostgres) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	return nil, nil
}

func (s *UserPostgres) DeleteUser(ctx context.Context, e *api.UserEmail) (*api.UserResponse, error) {
	return nil, nil
}
