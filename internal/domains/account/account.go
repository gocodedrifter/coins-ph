package account

import (
	"context"
	"github.com/coins-ph/internal"
)

type Service interface {
	Create(context context.Context, req internal.Account) (resp internal.Account, err error)
	Get(context context.Context, id string) (resp internal.Account, err error)
	AddBalance(context context.Context, id string, balance int64) (resp internal.Account, err error)
	GetAllAccounts(context context.Context, offset, limit int32) (resp []internal.Account, err error)
}
