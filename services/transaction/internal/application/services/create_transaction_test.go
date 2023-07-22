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

func Test_transactionDomainService_CreateTransactionServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repo := mock_repository.NewMockTransactionRepository(control)
	producer := mock_repository.NewMockTransactionProducer(control)

	td := &transactionDomainService{
		repository: repo,
		producer:   producer,
	}

	successTransactionTests := []struct {
		name        string
		transaction domain.TransactionDomain
	}{
		{
			name:        "It create an outbound transaction",
			transaction: *getTransactionDomain(money.New(-100, "BRL"), 1),
		},
		{
			name:        "It create a inbound transaction",
			transaction: *getTransactionDomain(money.New(100, "BRL"), 4),
		},
	}

	for _, tt := range successTransactionTests {
		t.Run(tt.name, func(t *testing.T) {
			repo.EXPECT().CreateTransaction(gomock.Eq(tt.transaction)).Return(&tt.transaction, nil)
			producer.EXPECT().TransactionCreated(gomock.Eq(tt.transaction))

			got, err := td.CreateTransactionServices(tt.transaction)
			require.Nil(t, err)
			require.Equal(t, tt.transaction, *got)
		})
	}

	errorTransactionTests := []struct {
		name        string
		transaction domain.TransactionDomain
		errExpected string
	}{
		{
			name:        "It errors when trying to create a transaction of type 'COMPRA A VISTA' with positive value",
			transaction: *getTransactionDomain(money.New(100, "BRL"), 1),
			errExpected: constants.InvalidOutboundOperation,
		},
		{
			name:        "It errors when trying to create a transaction of type 'COMPRA PARCELADA' with positive value",
			transaction: *getTransactionDomain(money.New(100, "BRL"), 2),
			errExpected: constants.InvalidOutboundOperation,
		},
		{
			name:        "It errors when trying to create a transaction of type 'SAQUE' with positive value",
			transaction: *getTransactionDomain(money.New(100, "BRL"), 3),
			errExpected: constants.InvalidOutboundOperation,
		},
		{
			name:        "It errors when trying to create a transaction of type 'PAGAMENTO' with negative value",
			transaction: *getTransactionDomain(money.New(-100, "BRL"), 4),
			errExpected: constants.InvalidInboundOperation,
		},
		{
			name:        "It errors when trying to create a transaction of invalid type",
			transaction: *getTransactionDomain(money.New(-100, "BRL"), 5),
			errExpected: constants.InvalidOperation,
		},
	}

	for _, transationTest := range errorTransactionTests {
		t.Run(transationTest.name, func(t *testing.T) {
			_, err := td.CreateTransactionServices(transationTest.transaction)
			require.Error(t, err)
			require.ErrorContains(t, err, transationTest.errExpected)
		})
	}

	t.Run("Error when trying to create a transaction", func(t *testing.T) {
		unknownErr := errors.New("unknown error")
		repo.EXPECT().CreateTransaction(gomock.Any()).Return(nil, &unknownErr)

		_, err := td.CreateTransactionServices(*getTransactionDomain(money.New(-100, "BRL"), 1))
		require.Error(t, err)
		require.ErrorContains(t, err, constants.ErrWasNotPossibleToCreateTransaction)
	})

}

func getTransactionDomain(amount *money.Money, operationTypeID int32) *domain.TransactionDomain {
	return &domain.TransactionDomain{
		AccountID:       uuid.New(),
		OperationTypeID: operationTypeID,
		EventDate:       time.Now(),
		Amount:          amount,
	}
}
