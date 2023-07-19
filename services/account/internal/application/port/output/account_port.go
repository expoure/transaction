package output

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
)

type AccountPort interface {
	CreateAccount(
		accountDomain domain.AccountDomain,
	) (*domain.AccountDomain, *rest_errors.RestErr)

	UpdateAccountBalanceByID(
		id uuid.UUID,
		transactionAmount int64,
	) *rest_errors.RestErr

	FindAccountBalanceByID(
		id uuid.UUID,
	) (*money.Money, *rest_errors.RestErr)

	FindAccountByDocumentNumber(
		documentNumber string,
	) (*domain.AccountDomain, *rest_errors.RestErr)

	FindAccountByID(
		id uuid.UUID,
	) (*domain.AccountDomain, *rest_errors.RestErr)
}
