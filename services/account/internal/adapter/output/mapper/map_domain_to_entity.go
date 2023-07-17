package mapper

import (
	"database/sql"

	"github.com/expoure/pismo/account/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/account/internal/application/domain"
)

func MapDomainToEntity(
	domain domain.AccountDomain,
) *entity.AccountEntity {
	return &entity.AccountEntity{
		ID:             domain.ID,
		DocumentNumber: domain.DocumentNumber,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      sql.NullTime{Time: *domain.DeletedAt},
	}
}
