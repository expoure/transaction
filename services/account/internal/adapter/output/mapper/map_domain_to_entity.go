package mapper

import (
	"github.com/expoure/pismo/account/internal/adapter/output/model/entity"
	"github.com/expoure/pismo/account/internal/application/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapDomainToEntity(
	domain domain.AccountDomain,
) *entity.AccountEntity {
	return &entity.AccountEntity{
		ID:             domain.ID,
		DocumentNumber: domain.DocumentNumber,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      pgtype.Timestamptz{Time: *domain.DeletedAt, Valid: true},
	}
}
