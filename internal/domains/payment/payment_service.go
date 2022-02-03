package payment

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/dataservice"
	"github.com/coins-ph/internal/dataservice/db"
)

type PaymentService struct {
	*dataservice.Payment
}

func NewPaymentService(d db.DBTX) *PaymentService{
	account := dataservice.NewPayment(d)
	return &PaymentService{account}
}

func (p *PaymentService) GetPaymentByID(ctx context.Context, id string, offset, limit int) (result []internal.Payment, err error) {
	result, err = p.Payment.GetPaymentByID(ctx, id, offset, limit)
	return
}

func (p *PaymentService) GetAllPayment(ctx context.Context, offset, limit int) (result []internal.Payment, err error) {
	result, err = p.Payment.GetAllPayment(ctx, offset, limit)
	return
}

func (p *PaymentService) Transfer(ctx context.Context, account, toAccount string, amount int64) (result []internal.Payment, err error) {
	result, err = p.Payment.Transfer(ctx, account, toAccount, amount)
	return
}
