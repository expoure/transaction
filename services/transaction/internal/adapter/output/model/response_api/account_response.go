package response_api

type Account struct {
	ID      string  `json:"id"`
	Balance Balance `json:"balance"`
}

type Balance struct {
	SmallUnitAmount int64  `json:"smallUnitAmount"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
}
