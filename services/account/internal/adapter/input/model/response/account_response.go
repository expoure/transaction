package response

type AccountResponse struct {
	ID             string  `json:"id"`
	DocumentNumber string  `json:"documentNumber"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
	DeletedAt      *string `json:"deletedAt"`
}
