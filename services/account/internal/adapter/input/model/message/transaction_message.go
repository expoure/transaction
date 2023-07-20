package message

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// algo errado aqui
type TransactionMessage struct {
	ID              uuid.UUID   `json:"id"`
	AccountID       uuid.UUID   `json:"accountId"`
	OperationTypeID int         `json:"operationTypeId"`
	EventDate       time.Time   `json:"eventDate"`
	Amount          json.Number `json:"amount"`
}
