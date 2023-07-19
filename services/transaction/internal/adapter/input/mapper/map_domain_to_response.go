package mapper

import (
	"github.com/expoure/pismo/transaction/internal/adapter/input/model/response"
	"github.com/expoure/pismo/transaction/internal/application/domain"
)

func MapDomainToResponse(
	accountDomain *domain.TransactionDomain,
) response.TransactionResponse {
	response := response.TransactionResponse{
		ID:              accountDomain.ID.String(),
		AccountID:       accountDomain.AccountID.String(),
		OperationTypeID: accountDomain.OperationTypeID,
		EventDate:       accountDomain.EventDate.String(),
		Amount:          accountDomain.Amount.AsMajorUnits(),
	}

	return response
}
