package input

import (
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/expoure/pismo/account/internal/configuration/rest_errors"
	"github.com/google/uuid"
)

type AccountDomainService interface {
	CreateAccountServices(domain.AccountDomain) (
		*domain.AccountDomain, *rest_errors.RestErr)

	FindAccountByIDServices(
		id uuid.UUID,
	) (*domain.AccountDomain, *rest_errors.RestErr)

	FindAccountByDocumentNumberServices(
		documentNumber string,
	) (*domain.AccountDomain, *rest_errors.RestErr)
}
