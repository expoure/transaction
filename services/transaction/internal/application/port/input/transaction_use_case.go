package input

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
	"github.com/google/uuid"
)

type TransactionDomainService interface {
	CreateTransactionServices(domain.TransactionDomain) (
		*domain.TransactionDomain, *rest_errors.RestErr)

	ListTransactionsService(
		accoountId uuid.UUID,
	) (*[]domain.TransactionDomain, *rest_errors.RestErr)
}
