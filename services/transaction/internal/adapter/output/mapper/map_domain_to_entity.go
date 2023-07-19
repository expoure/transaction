package mapper

import (
	"database/sql"

	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
	"github.com/expoure/pismo/transaction/internal/configuration/database/custom_types"
)

func MapDomainToEntity(
	domain domain.TransactionDomain,
) *entity.TransactionEntity {
	return &entity.TransactionEntity{
		ID:              domain.ID,
		AccountID:       domain.AccountID,
		OperationTypeID: sql.NullInt32{Int32: domain.OperationTypeID, Valid: true},
		EventDate:       domain.EventDate,
		Amount: &custom_types.Money{
			Amount:   domain.Amount.Amount(),
			Currency: domain.Amount.Currency().Code,
		},
	}
}
