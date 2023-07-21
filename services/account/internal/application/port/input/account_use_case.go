package input

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/customized_errors"
	"github.com/google/uuid"
)

type AccountDomainService interface {
	CreateAccountServices(domain.AccountDomain) (
		*domain.AccountDomain, *customized_errors.RestErr)

	FindAccountByIDServices(
		id uuid.UUID,
	) (*domain.AccountDomain, *customized_errors.RestErr)

	FindAccountByDocumentNumberServices(
		documentNumber string,
	) (*domain.AccountDomain, *customized_errors.RestErr)

	UpdateAccountBalanceByIDServices(
		id uuid.UUID,
		transactionAmount int64,
	) (*money.Money, *customized_errors.RestErr)
}
