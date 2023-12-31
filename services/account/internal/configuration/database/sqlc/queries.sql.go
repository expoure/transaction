// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: queries.sql

package sqlc

import (
	"context"

	custom_types "github.com/expoure/pismo/account/internal/configuration/database/custom_types"
	uuid "github.com/google/uuid"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (document_number, balance) VALUES ($1, (0,'BRL')) RETURNING id, document_number, balance, created_at, updated_at, deleted_at
`

func (q *Queries) CreateAccount(ctx context.Context, documentNumber string) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, documentNumber)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findAccountBalanceById = `-- name: FindAccountBalanceById :one
SELECT balance FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindAccountBalanceById(ctx context.Context, id uuid.UUID) (*custom_types.Money, error) {
	row := q.db.QueryRow(ctx, findAccountBalanceById, id)
	var balance *custom_types.Money
	err := row.Scan(&balance)
	return balance, err
}

const findAccountByDocumentNumber = `-- name: FindAccountByDocumentNumber :one
SELECT id, document_number, balance, created_at, updated_at, deleted_at FROM account
WHERE document_number = $1 LIMIT 1
`

func (q *Queries) FindAccountByDocumentNumber(ctx context.Context, documentNumber string) (Account, error) {
	row := q.db.QueryRow(ctx, findAccountByDocumentNumber, documentNumber)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findAccountById = `-- name: FindAccountById :one
SELECT id, document_number, balance, created_at, updated_at, deleted_at FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindAccountById(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRow(ctx, findAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.DocumentNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
