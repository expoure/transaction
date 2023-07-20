package request

type AccountRequest struct {
	DocumentNumber string `json:"documentNumber" binding:"required,min=11,max=14"`
}
