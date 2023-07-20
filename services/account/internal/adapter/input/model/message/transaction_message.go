package message

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

type TransactionMessage struct {
	AccountID uuid.UUID   `json:"accountId"`
	Amount    json.Number `json:"amount"`
}
