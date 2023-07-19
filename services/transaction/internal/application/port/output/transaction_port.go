package output

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
)

type TransactionPort interface {
	CreateTransaction(
		transactionDomain domain.TransactionDomain,
	) (*domain.TransactionDomain, *rest_errors.RestErr)
}
