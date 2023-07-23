package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	mock_repository "github.com/expoure/pismo/transaction/internal/adapter/output/repository/mock"
	"github.com/expoure/pismo/transaction/internal/application/constants"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"go.uber.org/mock/gomock"
)

func Test_transactionDomainService_ListTransactionsService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockTransactionRepository(control)

	t.Run("It lists two transactions", func(t *testing.T) {
		accountId := uuid.New()
		want := []domain.TransactionDomain{
			{
				AccountID:       accountId,
				OperationTypeID: 1,
				EventDate:       time.Now(),
				Amount:          *money.New(100, "BRL"),
			},
			{
				AccountID:       accountId,
				OperationTypeID: 2,
				EventDate:       time.Now(),
				Amount:          *money.New(200, "BRL"),
			},
		}

		repo.EXPECT().ListTransactionsByAccountID(gomock.Eq(accountId)).Return(&want, nil)

		td := &transactionDomainService{
			repository: repo,
		}

		got, err := td.ListTransactionsService(accountId)

		require.Nil(t, err)
		require.Equal(t, &want, got)
		require.Len(t, *got, len(want))
	})

	t.Run("It returns a empty list of transactions", func(t *testing.T) {
		accountId := uuid.New()

		repo.EXPECT().ListTransactionsByAccountID(gomock.Eq(accountId)).Return(&[]domain.TransactionDomain{}, nil)

		td := &transactionDomainService{
			repository: repo,
		}

		got, err := td.ListTransactionsService(accountId)

		require.Nil(t, err)
		require.Equal(t, &[]domain.TransactionDomain{}, got)
		require.Len(t, *got, 0)
	})

	t.Run("It errors when trying to list transactions", func(t *testing.T) {
		unknownErr := errors.New("unknown error")
		repo.EXPECT().ListTransactionsByAccountID(gomock.Any()).
			Return(nil, &unknownErr)

		td := &transactionDomainService{
			repository: repo,
		}

		_, err := td.ListTransactionsService(uuid.New())

		require.Error(t, err)
		require.ErrorContains(t, err, constants.ErrWasNotPossibleToListTransactions)
	})

}
