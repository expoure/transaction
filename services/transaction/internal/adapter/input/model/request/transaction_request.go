package request

import "github.com/google/uuid"

type AccountRequest struct {
	AccountID       uuid.UUID `json:"accountID" binding:"required,uuid4"`
	OperationTypeID uuid.UUID `json:"operationTypeID" binding:"required,min=1,max=4"`
	Amount          float64   `json:"amount" binding:"required,min=0"`
}
