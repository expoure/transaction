package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/expoure/pismo/transaction/internal/configuration/database/custom_types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestQueries_ListTransactions(t *testing.T) {
	ctx := context.TODO()
	q := &Queries{
		db: TestQueries.db,
	}

	t.Run("It list no transactions", func(t *testing.T) {
		got, err := q.ListTransactions(ctx, uuid.New())

		require.NoError(t, err)
		require.Len(t, got, len([]Transaction{}))

	})

	t.Run("It list one transaction", func(t *testing.T) {
		transactionToCreate := Transaction{
			AccountID:       uuid.New(),
			OperationTypeID: pgtype.Int4{Int32: 1, Valid: true},
			EventDate:       time.Now(),
			Amount: &custom_types.Money{
				Amount:   100,
				Currency: "USD",
			},
		}

		err := insertTransactionHelper(transactionToCreate)

		got, err := q.ListTransactions(ctx, transactionToCreate.AccountID)

		require.NoError(t, err)
		require.Len(t, got, 1)

	})
}

func insertTransactionHelper(transaction Transaction) error {
	_, err := TestQueries.db.Exec(
		context.TODO(),
		`INSERT INTO transaction (
			account_id,
			operation_type_id,
			event_date,
			amount
		) 
		VALUES ($1, $2, $3, ($4, $5)) `,
		transaction.AccountID,
		transaction.OperationTypeID.Int32,
		transaction.EventDate,
		transaction.Amount.Amount,
		transaction.Amount.Currency,
	)

	return err
}
