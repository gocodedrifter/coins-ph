package payment

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/go-kit/kit/endpoint"
)

type getPaymentByIdRequest struct {
	ID     string `json:"id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type getPaymentResponse struct {
	Err  error              `json:"error,omitempty"`
	Data []internal.Payment `json:"data,omitempty"`
}

func (r getPaymentResponse) error() error { return r.Err }

func makePaymentByIdEndpoint(s *PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPaymentByIdRequest)
		resp, err := s.GetPaymentByID(ctx, req.ID, req.Offset, req.Limit)
		return getPaymentResponse{
			Err:  err,
			Data: resp,
		}, err
	}
}

type transferAmountRequest struct {
	Account   string `json:"account,omitempty"`
	ToAccount string `json:"to_account,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
}

func makeTransferAmountEndpoint(s *PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transferAmountRequest)
		resp, err := s.Transfer(ctx, req.Account, req.ToAccount, req.Amount)
		return getPaymentResponse{
			Err:  err,
			Data: resp,
		}, err
	}
}
