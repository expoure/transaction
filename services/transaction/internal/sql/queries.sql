-- name: ListTransactions :many
SELECT * FROM transaction WHERE account_id = $1 ORDER BY event_date DESC;
