package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxPool interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}
