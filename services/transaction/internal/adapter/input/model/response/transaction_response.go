package response

type TransactionResponse struct {
	ID              string  `json:"id"`
	OperationTypeID int32   `json:"operationTypeId"`
	AccountID       string  `json:"accountId"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"eventDate"`
}
