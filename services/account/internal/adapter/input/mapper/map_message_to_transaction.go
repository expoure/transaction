package mapper

import (
	"encoding/json"

	"github.com/expoure/pismo/account/internal/adapter/input/model/message"
)

func MapMessageToTransaction(
	msgValue []byte,
) *message.TransactionMessage {
	var transactionMessage message.TransactionMessage
	json.Unmarshal(msgValue, &transactionMessage)
	return &transactionMessage
}
