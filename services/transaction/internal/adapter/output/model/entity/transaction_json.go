package entity

type TransactionJson struct {
	ID              string `json:"id"`
	AccountID       string `json:"accountId"`
	OperationTypeID int32  `json:"operationTypeId"`
	EventDate       string `json:"eventDate"`
	Amount          int64  `json:"amount"`
}
