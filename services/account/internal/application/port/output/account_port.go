package output

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/google/uuid"
)

type AccountRepositoryPort interface {
	CreateAccount(
		accountDomain domain.AccountDomain,
	) (*domain.AccountDomain, *error)

	UpdateAccountBalanceByID(
		id uuid.UUID,
		transactionAmount int64,
	) (*money.Money, *error)

	FindAccountBalanceByID(
		id uuid.UUID,
	) (*money.Money, *error)

	FindAccountByDocumentNumber(
		documentNumber string,
	) (*domain.AccountDomain, *error)

	FindAccountByID(
		id uuid.UUID,
	) (*domain.AccountDomain, *error)
}
