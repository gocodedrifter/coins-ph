package account

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/dataservice"
	"github.com/coins-ph/internal/dataservice/db"
)

type AccountService struct {
	*dataservice.Account
}

func NewAccountService(d db.DBTX) *AccountService {
	account := dataservice.NewAccount(d)
	return &AccountService{account}
}

func (ac *AccountService) Create(context context.Context, req internal.Account) (resp internal.Account, err error) {
	resp, err = ac.Account.Create(context, req)
	return
}

func (ac *AccountService) Get(context context.Context, id string) (resp internal.Account, err error) {
	resp, err = ac.Account.Get(context, id)
	return
}

func (ac *AccountService) AddBalance(context context.Context, id string, balance int64) (resp internal.Account, err error) {
	resp, err = ac.Account.AddBalance(context, id, balance)
	return
}

func (ac *AccountService) GetAllAccounts(context context.Context, offset, limit int32) (resp []internal.Account, err error) {
	resp, err = ac.Account.GetAllAccounts(context, offset, limit)
	return
}
