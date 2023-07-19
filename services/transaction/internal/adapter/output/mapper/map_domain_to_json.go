package mapper

import (
	"encoding/json"

	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
)

func MapDomainToJson(
	domain domain.TransactionDomain,
) *[]byte {

	t := entity.TransactionJson{
		ID:              domain.ID.String(),
		AccountID:       domain.AccountID.String(),
		OperationTypeID: domain.OperationTypeID,
		EventDate:       domain.EventDate.String(),
		Amount:          domain.Amount.Amount(),
	}

	transactionJson, _ := json.Marshal(t)
	return &transactionJson
}
