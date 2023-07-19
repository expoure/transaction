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
	Amount          *money.Money // talvez amount não seja o melhor nome
	EventDate       time.Time
}

type OperationType struct {
	ID          int32
	Description string
}
