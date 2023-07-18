package response

type AccountResponse struct {
	ID             string          `json:"id"`
	DocumentNumber string          `json:"documentNumber"`
	Balance        BalanceResponse `json:"balance"`
	CreatedAt      string          `json:"createdAt"`
	UpdatedAt      string          `json:"updatedAt"`
	DeletedAt      *string         `json:"deletedAt"`
}

type BalanceResponse struct {
	SmallUnitAmount int64  `json:"smallUnitAmount"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
}
