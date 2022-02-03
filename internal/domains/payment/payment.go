package payment

import (
	"context"
	"github.com/coins-ph/internal"
)

type Service interface {
	GetPaymentByID(ctx context.Context, id string, offset, limit int) (result []internal.Payment, err error)
	GetAllPayment(ctx context.Context, offset, limit int) (result []internal.Payment, err error)
	Transfer(ctx context.Context, account, toAccount string, amount int64) (result []internal.Payment, err error)
}
