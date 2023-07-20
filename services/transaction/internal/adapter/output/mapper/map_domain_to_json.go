package mapper

import (
	"encoding/json"
	"strconv"

	"github.com/expoure/pismo/transaction/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/transaction/internal/application/domain"
)

func MapDomainToEventJson(
	domain domain.TransactionDomain,
) *[]byte {

	stringAmount := strconv.Itoa(int(domain.Amount.Amount()))
	t := entity.TransactionEventJson{
		AccountID: domain.AccountID.String(),
		Amount:    stringAmount,
	}

	transactionJson, _ := json.Marshal(t)
	return &transactionJson
}
