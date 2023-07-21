package sqlc

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TestQueries *Queries
var TestDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var (
		HOST     = "localhost"
		PORT     = "5440"
		USER     = "test"
		PASSWORD = "test"
		DATABASE = "test"
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	conn, err := pgxpool.New(context.Background(), strConn)

	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Can not ping database", err)
	}
	TestQueries = New(conn)
	TestDB = conn

	os.Exit(m.Run())
}
