package request

type TransactionRequest struct {
	AccountID       string  `json:"accountID" binding:"required,uuid4"`
	OperationTypeID int32   `json:"operationTypeID" binding:"required,min=1,max=4"`
	Amount          float64 `json:"amount" binding:"required,min=0"`
}
