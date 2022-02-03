package repository

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/repository/db"
)

type Payment struct {
	q *db.Queries
}

func NewPayment(d db.DBTX) *Payment {
	return &Payment{
		q: db.New(d),
	}
}

func (p *Payment) GetPaymentByID(ctx context.Context, id string, offset, limit int) (result []internal.Payment, err error) {
	payments, err := p.q.ListPaymentsById(ctx, db.ListPaymentsByIdParams{
		Account: newNullString(id),
		Page:    int32(offset),
		Number:  int32(limit),
	})

	if err != nil {
		return nil, err
	}

	for _, pay := range payments {
		result = append(result, internal.Payment{
			Account:     pay.AccountID.String,
			Amount:      pay.Amount.(int64),
			FromAccount: pay.FromAccountID.String,
			ToAccount:   pay.ToAccountID.String,
			Direction:   pay.Direction.String,
		})
	}

	return result, nil
}

func (p *Payment) GetAllPayment(ctx context.Context, offset, limit int) (result []internal.Payment, err error) {
	payments, err := p.q.ListPayments(ctx, db.ListPaymentsParams{
		Page:   int32(offset),
		Number: int32(limit),
	})

	if err != nil {
		return nil, err
	}

	for _, pay := range payments {
		result = append(result, internal.Payment{
			Account:     pay.AccountID.String,
			Amount:      pay.Amount.(int64),
			FromAccount: pay.FromAccountID.String,
			ToAccount:   pay.ToAccountID.String,
			Direction:   pay.Direction.String,
		})
	}

	return result, nil
}

func (p *Payment) Transfer(ctx context.Context, account, toAccount string, amount int64) (result []internal.Payment, err error) {
	payments, err := p.q.Payment(ctx, db.PaymentParams{
		Account:   account,
		ToAccount: toAccount,
		Amount:    newNullInt64(amount),
	})

	if err != nil {
		return nil, err
	}

	for _, pay := range payments {
		result = append(result, internal.Payment{
			Account:     pay.Account,
			Amount:      pay.Amount.Int64,
			FromAccount: pay.FromAccount,
			ToAccount:   pay.ToAccount,
			Direction:   pay.Direction,
		})
	}

	return
}
