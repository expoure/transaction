package domain

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type AccountDomain struct {
	ID             uuid.UUID
	DocumentNumber string
	Balance        *money.Money
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
