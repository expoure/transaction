package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/database/custom_types"
	mock_sqlc_repository "github.com/expoure/pismo/transaction/internal/configuration/database/sqlc/mock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_transactionRepositoryImpl_CreateTransaction(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	querier := mock_sqlc_repository.NewMockQuerier(control)

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	tr := &transactionRepositoryImpl{
		queries:  querier,
		connPool: mock,
	}

	t.Run("Success", func(t *testing.T) {
		transactionDomain := domain.TransactionDomain{
			ID:              uuid.New(),
			AccountID:       uuid.New(),
			OperationTypeID: 1,
			EventDate:       time.Now(),
			Amount:          *money.New(-1, "BRL"),
		}

		amount := custom_types.Money{Amount: -1, Currency: "BRL"}
		rows := mock.NewRows([]string{"id", "account_id", "operation_type_id", "event_date", "amount"}).
			AddRow(
				transactionDomain.ID,
				transactionDomain.AccountID,
				"1",
				transactionDomain.EventDate,
				amount,
			)

		mock.ExpectQuery("INSERT INTO transaction").
			WithArgs(
				transactionDomain.AccountID,
				transactionDomain.OperationTypeID,
				transactionDomain.EventDate,
				transactionDomain.Amount.Amount(),
				transactionDomain.Amount.Currency().Code,
			).
			WillReturnRows(rows)

		got, err := tr.CreateTransaction(transactionDomain)
		require.Nil(t, err)
		require.Equal(t, transactionDomain, *got)
	})
}

func Test_transactionRepositoryImpl_ListTransactionsByAccountID(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	querier := mock_sqlc_repository.NewMockQuerier(control)

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	tr := &transactionRepositoryImpl{
		queries:  querier,
		connPool: mock,
	}

	t.Run("Success", func(t *testing.T) {
		transactionOneId := uuid.New()
		transactionTwoId := uuid.New()
		accountID := uuid.New()
		eventDate := time.Now()

		transactions := []domain.TransactionDomain{
			{
				ID:              transactionOneId,
				AccountID:       accountID,
				OperationTypeID: 1,
				EventDate:       eventDate,
				Amount:          *money.New(-1, "BRL"),
			},
			{
				ID:              transactionTwoId,
				AccountID:       accountID,
				OperationTypeID: 4,
				EventDate:       eventDate,
				Amount:          *money.New(1, "BRL"),
			},
		}

		transactionsEntities := []entity.TransactionEntity{
			{
				ID:              transactionOneId,
				AccountID:       accountID,
				OperationTypeID: pgtype.Int4{Int32: 1},
				EventDate:       eventDate,
				Amount:          custom_types.Money{Amount: -1, Currency: "BRL"},
			},
			{
				ID:              transactionTwoId,
				AccountID:       accountID,
				OperationTypeID: pgtype.Int4{Int32: 4},
				EventDate:       eventDate,
				Amount:          custom_types.Money{Amount: 1, Currency: "BRL"},
			},
		}

		querier.EXPECT().ListTransactions(gomock.Any(), gomock.Any()).Return(transactionsEntities, nil)

		got, err := tr.ListTransactionsByAccountID(uuid.New())
		require.Nil(t, err)
		require.Equal(t, transactions, *got)
	})

	t.Run("Error", func(t *testing.T) {
		unknownErr := errors.New("unknown error")
		querier.EXPECT().ListTransactions(gomock.Any(), gomock.Any()).Return(nil, unknownErr)

		_, err := tr.ListTransactionsByAccountID(uuid.New())
		require.NotNil(t, err)
	})
}
