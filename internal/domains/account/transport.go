package account

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/coins-ph/internal"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(s *AccountService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createAccountHandler := kithttp.NewServer(
		makeCreateAccountEndpoint(s),
		decodeCreateAccountRequest,
		encodeResponse,
		opts...,
	)

	getAccountHandler := kithttp.NewServer(
		makeGetAccountEndpoint(s),
		decodeGetAccountRequest,
		encodeResponse,
		opts...,
	)

	getAllAccountHandler := kithttp.NewServer(
		makeGetAllAccountEndpoint(s),
		decodeGetAllAccountRequest,
		encodeResponse,
		opts...,
	)

	addBalanceAccountHandler := kithttp.NewServer(
		makeAddBalanceEndpoint(s),
		decodeAddBalanceRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/wallet/v1/account", createAccountHandler).Methods("POST")
	r.Handle("/wallet/v1/account/all", getAllAccountHandler).Methods("POST")
	r.Handle("/wallet/v1/account/topup", addBalanceAccountHandler).Methods("POST")
	r.Handle("/wallet/v1/account/{id}", getAccountHandler).Methods("GET")
	return r
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return accountRequest{
		ID: body.ID,
	}, nil
}

func decodeAddBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ID      string `json:"id"`
		Balance int64  `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return accountBalanceRequest{
		ID:      body.ID,
		Balance: body.Balance,
	}, nil
}

func decodeGetAllAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Offset int32 `json:"offset"`
		Limit  int32 `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return getAllAccountRequest{
		Offset: body.Offset,
		Limit:  body.Limit,
	}, nil
}

func decodeGetAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return accountRequest{ID: id}, nil
}

var errBadRoute = errors.New("bad route")

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case internal.ErrAccountNotFound:
		w.WriteHeader(http.StatusNotFound)
	case internal.ErrorInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
