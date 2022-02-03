package payment

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

func MakeHandler(s *PaymentService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getAllPaymentByID := kithttp.NewServer(
		makePaymentByIdEndpoint(s),
		decodeGetPaymentByIDRequest,
		encodeResponse,
		opts...,
	)

	transferAmount := kithttp.NewServer(
		makeTransferAmountEndpoint(s),
		decodeTransferAmoountRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/wallet/v1/payment", getAllPaymentByID).Methods("POST")
	r.Handle("/wallet/v1/payment/transfer", transferAmount).Methods("POST")
	return r
}

func decodeGetPaymentByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ID     string `json:"id"`
		Offset int    `json:"offset"`
		Limit  int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return getPaymentByIdRequest{
		ID:     body.ID,
		Offset: body.Offset,
		Limit:  body.Limit,
	}, nil
}

func decodeTransferAmoountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Account   string `json:"account"`
		ToAccount string `json:"to_account"`
		Amount    int64  `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return transferAmountRequest{
		Account:   body.Account,
		ToAccount: body.ToAccount,
		Amount:    body.Amount,
	}, nil
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
