package request

type AccountRequest struct {
	DocumentNumber string `json:"documentNumber" binding:"required,min=11,max=14"`
	// Password       string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	// Name           string `json:"name" binding:"required,min=4,max=100"`
	// Age            int8   `json:"age" binding:"required,min=1,max=140"`
}
