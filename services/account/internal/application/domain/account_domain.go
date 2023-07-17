package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountDomain struct {
	ID             uuid.UUID
	DocumentNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
