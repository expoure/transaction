package output

import (
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/response_api"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateTransaction(
		transactionDomain domain.TransactionDomain,
	) (*domain.TransactionDomain, *error)

	ListTransactionsByAccountID(
		accoountId uuid.UUID,
	) (*[]domain.TransactionDomain, *error)
}

type TransactionProducer interface {
	TransactionCreated(
		transactionDomain domain.TransactionDomain,
	)
}

type AccountHttpClient interface {
	GetAccount(
		id string,
	) (*response_api.Account, error)
}
