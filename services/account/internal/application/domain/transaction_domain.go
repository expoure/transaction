package domain

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type TransactionDomain struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	OperationTypeID int32
	Amount          *money.Money
	EventDate       time.Time
}
