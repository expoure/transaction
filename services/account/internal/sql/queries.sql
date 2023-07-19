-- name: FindAccountById :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: FindAccountByDocumentNumber :one
SELECT * FROM account
WHERE document_number = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO account (document_number, balance) VALUES ($1, (0,'BRL')) RETURNING *;

-- name: FindAccountBalanceById :one
SELECT balance FROM account
WHERE id = $1 LIMIT 1;
