package output

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateTransaction(
		transactionDomain domain.TransactionDomain,
	) (*domain.TransactionDomain, *rest_errors.RestErr)

	ListTransactionsByAccountID(
		accoountId uuid.UUID,
	) (*[]domain.TransactionDomain, *rest_errors.RestErr)
}

type TransactionProducer interface {
	TransactionCreated(
		transactionDomain domain.TransactionDomain,
	)
}
