package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/s3rzh/go-grpc-user-service/pkg/api"
)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
	return &UserPostgres{db: db}
}

func (s *UserPostgres) CreateUser(ctx context.Context, u *api.User) (int, error) {
	var userId int

	row := s.db.QueryRow(ctx,
		"INSERT INTO users (age, email) VALUES ($1, $2) RETURNING id",
		u.Age, u.Email)

	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserPostgres) GetUsers(ctx context.Context) (*api.UsersResponse, error) {
	return nil, nil
}

func (s *UserPostgres) DeleteUser(ctx context.Context, e string) error {
	_, err := s.db.Exec(ctx, "DELETE FROM users WHERE email = $1", e)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserPostgres) CheckUserForExists(ctx context.Context, e string) (bool, error) {
	var count int
	row := s.db.QueryRow(ctx,
		"SELECT 1 FROM users WHERE email = $1",
		e)

	err := row.Scan(&count)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
