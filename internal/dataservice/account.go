package dataservice

import (
	"context"
	"errors"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/dataservice/db"
	"log"
)

type Account struct {
	q *db.Queries
}

func NewAccount(d db.DBTX) *Account {
	return &Account{
		q: db.New(d),
	}
}

func (a *Account) Create(ctx context.Context, params internal.Account) (internal.Account, error) {

	row, err := a.q.GetAccount(ctx, params.ID)
	if len(row.ID) > 0 {
		return internal.Account{}, internal.ErrAccountAlreadyExist
	}

	newID, err := a.q.CreateAccount(ctx, db.CreateAccountParams{
		ID:   params.ID,
		Name: newNullString(params.Name),
	})

	if err != nil {
		return internal.Account{}, internal.ErrAccountCreation
	}

	_, err = a.q.TopupAccountBalance(ctx, db.TopupAccountBalanceParams{
		ID:       newID,
		Balance:  newNullInt64(0),
		Currency: newNullString("USD"),
	})

	if err != nil {
		log.Println(err)
	}

	return internal.Account{ID: newID, Name: params.Name, Balance: params.Balance, Currency: "USD"}, nil
}

func (a *Account) Get(ctx context.Context, id string) (internal.Account, error) {

	row, _ := a.q.GetAccount(ctx, id)
	if len(row.ID) <= 0 {
		return internal.Account{}, internal.ErrAccountNotFound
	}

	return internal.Account{
		ID:       row.ID,
		Balance:  row.Balance.Int64,
		Currency: row.Currency.String,
	}, nil
}

func (a *Account) AddBalance(ctx context.Context, id string, balance int64) (internal.Account, error) {
	if balance <= 0 {
		return internal.Account{}, errors.New("balance must be greater than zero")
	}
	row, err := a.q.TopupAccountBalance(ctx, db.TopupAccountBalanceParams{
		ID:       id,
		Balance:  newNullInt64(balance),
		Currency: newNullString("USD"),
	})

	if err != nil {
		return internal.Account{}, internal.ErrorCodeUnknown
	}

	return internal.Account{
		ID:       row.ID,
		Balance:  row.Balance.Int64,
		Currency: row.Currency.String,
	}, nil
}

func (a *Account) GetAllAccounts(ctx context.Context, offset, limit int32) (accounts []internal.Account, err error) {
	acc, err := a.q.ListAccounts(ctx, db.ListAccountsParams{
		Page:   limit,
		Number: offset,
	})

	for _, obj := range acc {
		accounts = append(accounts, internal.Account{
			ID:       obj.ID,
			Balance:  obj.Balance.Int64,
			Currency: obj.Currency.String,
		})
	}

	return
}
