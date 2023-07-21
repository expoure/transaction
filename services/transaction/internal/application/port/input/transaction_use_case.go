package input

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
	"github.com/google/uuid"
)

type TransactionDomainService interface {
	CreateTransactionServices(domain.TransactionDomain) (
		*domain.TransactionDomain, *customized_errors.RestErr)

	ListTransactionsService(
		accoountId uuid.UUID,
	) (*[]domain.TransactionDomain, *customized_errors.RestErr)
}
