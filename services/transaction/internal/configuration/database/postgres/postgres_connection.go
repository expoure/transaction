package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SSLMode string

const (
	SslDisable SSLMode = "disable"
)

func NewPostgresConnection(ctx context.Context) (*pgxpool.Pool, error) {
	var (
		HOST     = os.Getenv("POSTGRES_HOST")
		PORT     = os.Getenv("POSTGRES_PORT")
		USER     = os.Getenv("POSTGRES_USER")
		PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		DATABASE = os.Getenv("POSTGRES_DB")
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DATABASE, SslDisable)
	db, err := pgxpool.New(context.Background(), strConn)

	if err != nil {
		panic(err)
	}

	return db, err
}
