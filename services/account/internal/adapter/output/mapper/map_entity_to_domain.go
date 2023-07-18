package mapper

import (
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/expoure/pismo/account/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/account/internal/application/domain"
)

func MapEntityToDomain(
	entity entity.AccountEntity,
) *domain.AccountDomain {
	domainConverted := &domain.AccountDomain{
		ID:             entity.ID,
		DocumentNumber: entity.DocumentNumber,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		Balance:        mapBalance(entity.Balance),
	}

	if entity.DeletedAt.Valid {
		domainConverted.DeletedAt = &entity.DeletedAt.Time
	} else {
		domainConverted.DeletedAt = nil
	}

	return domainConverted
}

func mapBalance(entityBalance string) *money.Money {
	compositeString := strings.Trim(entityBalance, "()")

	values := strings.Split(compositeString, ",")

	amount, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		// ajustar isso aqui: colocar o log
		panic(err)
	}
	currency := strings.Trim(values[1], "'")

	return money.New(amount, currency)
}
