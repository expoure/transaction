package output

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateTransaction(
		transactionDomain domain.TransactionDomain,
	) (*domain.TransactionDomain, *customized_errors.RestErr)

	ListTransactionsByAccountID(
		accoountId uuid.UUID,
	) (*[]domain.TransactionDomain, *customized_errors.RestErr)
}

type TransactionProducer interface {
	TransactionCreated(
		transactionDomain domain.TransactionDomain,
	)
}
