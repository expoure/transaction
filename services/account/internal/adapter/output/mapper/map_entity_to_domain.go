package mapper

import (
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
	}

	if entity.DeletedAt.Valid {
		domainConverted.DeletedAt = &entity.DeletedAt.Time
	} else {
		domainConverted.DeletedAt = nil
	}

	return domainConverted
}
