package internal

type Account struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"-,omitempty"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency,omitempty"`
}

type Payment struct {
	Account     string `json:"account,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	FromAccount string `json:"from_account,omitempty"`
	ToAccount   string `json:"to_account,omitempty"`
	Direction   string `json:"direction,omitempty"`
}
