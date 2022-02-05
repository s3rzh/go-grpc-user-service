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

// var dbURL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
// "postgres://restservice@localhost/restservice?sslmode=disable
//     conn, _ := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5432/test")
// "postgres://postgres:qwerty@localhost:5432/api?sslmode=disable"

// defer conn.Close(context.Background())

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {

	var dbURL string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
