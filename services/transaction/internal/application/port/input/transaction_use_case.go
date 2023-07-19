package input

import (
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/rest_errors"
)

type AccountDomainService interface {
	CreateTransactionServices(domain.TransactionDomain) (
		*domain.TransactionDomain, *rest_errors.RestErr)
}
