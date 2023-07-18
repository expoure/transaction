package mapper

import (
	"github.com/expoure/pismo/account/internal/adapter/input/model/response"
	"github.com/expoure/pismo/account/internal/application/domain"
)

func MapDomainToResponse(
	accountDomain *domain.AccountDomain,
) response.AccountResponse {
	response := response.AccountResponse{
		ID:             accountDomain.ID.String(),
		DocumentNumber: accountDomain.DocumentNumber,
		Balance: response.BalanceResponse{
			Amount:          accountDomain.Balance.Display(),
			Currency:        accountDomain.Balance.Currency().Code,
			SmallUnitAmount: accountDomain.Balance.Amount(),
		},
		CreatedAt: accountDomain.CreatedAt.String(),
		UpdatedAt: accountDomain.UpdatedAt.String(),
	}

	if accountDomain.DeletedAt != nil {
		deletedAt := accountDomain.DeletedAt.String()
		response.DeletedAt = &deletedAt
	} else {
		response.DeletedAt = nil
	}

	return response
}
