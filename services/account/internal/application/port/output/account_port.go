package output

import (
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
)

type AccountPort interface {
	CreateAccount(
		userDomain domain.AccountDomain,
	) (*domain.AccountDomain, *rest_errors.RestErr)

	FindAccountByDocumentNumber(
		documentNumber string,
	) (*domain.AccountDomain, *rest_errors.RestErr)

	FindAccountByID(
		id uuid.UUID,
	) (*domain.AccountDomain, *rest_errors.RestErr)
}
