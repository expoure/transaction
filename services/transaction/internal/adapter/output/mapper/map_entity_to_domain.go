package mapper

import (
	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
)

func MapEntityToDomain(
	entity entity.TransactionEntity,
) *domain.TransactionDomain {
	domainConverted := &domain.TransactionDomain{
		ID:              entity.ID,
		AccountID:       entity.AccountID,
		OperationTypeID: entity.OperationTypeID.Int32,
		EventDate:       entity.EventDate,
		Amount:          *money.New(entity.Amount.Amount, entity.Amount.Currency),
	}

	return domainConverted
}
