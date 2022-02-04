package dataservice_test

import (
	"context"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/dataservice"
	"testing"
)

func TestPayment_Transfer(t *testing.T) {
	t.Parallel()

	t.Run("Transfer: OK", func(t *testing.T) {
		t.Parallel()

		pool := newDB(t)
		store := dataservice.NewAccount(pool)
		acc1, err := store.Create(context.Background(),
			internal.Account{
				ID:   "dbbjb541",
				Name: "william",
			})
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		acc2, err := store.Create(context.Background(),
			internal.Account{
				ID:   "dbbkb541",
				Name: "john",
			})
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		store.AddBalance(context.Background(), acc1.ID, 12000)
		store.AddBalance(context.Background(), acc2.ID, 2500)

		payment := dataservice.NewPayment(pool)
		payment.Transfer(context.Background(), acc1.ID, acc2.ID, 2500)

		acc1, err = store.Get(context.Background(), acc1.ID)
		acc2, err = store.Get(context.Background(), acc2.ID)
		assertEqual(t, acc1.Balance, int64(9500), "")
		assertEqual(t, acc2.Balance, int64(5000), "")
	})

}
