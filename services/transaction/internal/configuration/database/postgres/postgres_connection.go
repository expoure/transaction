package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type SSLMode string

const (
	SslDisable SSLMode = "disable"
)

func NewPostgresConnection(ctx context.Context) (*sql.DB, error) {
	var (
		HOST     = os.Getenv("POSTGRES_HOST")
		PORT     = os.Getenv("POSTGRES_PORT")
		USER     = os.Getenv("POSTGRES_USER")
		PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		DATABASE = os.Getenv("POSTGRES_DB")
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DATABASE, SslDisable)
	db, err := sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}

	return db, err
}
