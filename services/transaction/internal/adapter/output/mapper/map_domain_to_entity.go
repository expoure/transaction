package mapper

import (
	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/database/custom_types"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapDomainToEntity(
	domain domain.TransactionDomain,
) *entity.TransactionEntity {
	return &entity.TransactionEntity{
		ID:              domain.ID,
		AccountID:       domain.AccountID,
		OperationTypeID: pgtype.Int4{Int32: domain.OperationTypeID, Valid: true},
		EventDate:       domain.EventDate,
		Amount: custom_types.Money{
			Amount:   domain.Amount.Amount(),
			Currency: domain.Amount.Currency().Code,
		},
	}
}
