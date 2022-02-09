package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {

	var dbURL string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
