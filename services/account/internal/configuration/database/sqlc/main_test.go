package sqlc

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var TestQueries *Queries
var TestDB *sql.DB

func TestMain(m *testing.M) {
	var (
		HOST     = "localhost"
		PORT     = "5440"
		USER     = "test"
		PASSWORD = "test"
		DATABASE = "test"
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)
	conn, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Can not ping database", err)
	}
	TestQueries = New(conn)
	TestDB = conn

	os.Exit(m.Run())
}
