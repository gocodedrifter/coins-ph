package internal

type Account struct {
	ID       string
	Name     string
	Balance  int64
	Currency string
}

type Payment struct {
	Account     string
	Amount      int64
	FromAccount string
	ToAccount   string
	Direction   string
}
