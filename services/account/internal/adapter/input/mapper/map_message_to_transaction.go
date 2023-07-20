package mapper

import (
	"encoding/json"
	"fmt"

	"github.com/expoure/pismo/account/internal/adapter/input/model/message"
)

func MapMessageToTransaction(
	msgValue []byte,
) *message.TransactionMessage {
	var transactionMessage message.TransactionMessage
	json.Unmarshal(msgValue, &transactionMessage)
	fmt.Println((transactionMessage))
	return &transactionMessage
	// return &domain.TransactionDomain{
	// 	ID:              transactionMessage.,
	// }
}
