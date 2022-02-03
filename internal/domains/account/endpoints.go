package account

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/go-kit/kit/endpoint"
)

type accountRequest struct {
	ID string `json:"id,omitempty"`
}

type accountResponse struct {
	Err     error            `json:"error,omitempty"`
	Account internal.Account `json:"account,omitempty"`
}

func (r accountResponse) error() error { return r.Err }

func makeCreateAccountEndpoint(s *AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(accountRequest)
		resp, err := s.Create(ctx, internal.Account{ID: req.ID})
		return accountResponse{
			Err:     err,
			Account: resp,
		}, err
	}
}

func makeGetAccountEndpoint(s *AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(accountRequest)
		acc, err := s.Get(ctx, req.ID)
		return accountResponse{
			Err:     err,
			Account: acc,
		}, err
	}
}

type accountBalanceRequest struct {
	ID      string `json:"id,omitempty"`
	Balance int64  `json:"balance,omitempty"`
}

func makeAddBalanceEndpoint(s *AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(accountBalanceRequest)
		acc, err := s.AddBalance(ctx, req.ID, req.Balance)
		return accountResponse{
			Err:     err,
			Account: acc,
		}, nil
	}
}

type getAllAccountRequest struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}

type getAllAccountResponse struct {
	Err     error              `json:"error,omitempty"`
	Account []internal.Account `json:"accounts,omitempty"`
}

func makeGetAllAccountEndpoint(s *AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAllAccountRequest)
		acc, err := s.GetAllAccounts(ctx, req.Offset, req.Limit)
		return getAllAccountResponse{
			Err:     err,
			Account: acc,
		}, err
	}
}
