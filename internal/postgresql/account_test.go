package postgresql_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/coins-ph/internal"
	"github.com/coins-ph/internal/postgresql"
	migrate "github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"
)

func TestAccount_Create(t *testing.T) {
	t.Parallel()

	t.Run("Create: OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewAccount(newDB(t))
		acc, err := store.Create(context.Background(),
			internal.Account{
				ID: "dbbjb541",
				Name: "arip",
			})
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if acc.ID == "" {
			t.Fatalf("expected valid record, got empty value")
		}

		acc, err = store.Get(context.Background(), "dbbjb541")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if acc.Balance > 0 || acc.Balance < 0 {
			t.Fatalf("expected valid record, got balance")
		}

		assertEqual(t, acc.Currency, "USD", "")
		assertEqual(t, acc.ID, "dbbjb541", "")
	})

}

func TestAccount_TopupBalance(t *testing.T) {
	t.Parallel()

	t.Run("TopupBalance : OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewAccount(newDB(t))
		acc, err := store.Create(context.Background(),
			internal.Account{
				ID: "dbbjb541",
				Name: "arip",
			})
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if acc.ID == "" {
			t.Fatalf("expected valid record, got empty value")
		}

		store.AddBalance(context.Background(), acc.ID, 5000)
		acc, err = store.AddBalance(context.Background(), acc.ID, 7550)

		assertEqual(t, acc.Balance, int64(12550), "")
	})
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	switch v := a.(type) {
	case int64:
		if v == b.(int64) {
			return
		}
	case string:
		if v == b.(string) {
			return
		}
	}
	if a.(int64) == b.(int64) {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func newDB(tb testing.TB) *pgxpool.Pool {
	tb.Helper()

	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("username", "password"),
		Path:   "coins",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	//-

	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf("Couldn't connect to docker: %s", err)
	}

	pool.MaxWait = 10 * time.Second

	pw, _ := dsn.User.Password()

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.5-alpine",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", dsn.User.Username()),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", pw),
			fmt.Sprintf("POSTGRES_DB=%s", dsn.Path),
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		tb.Fatalf("Couldn't start resource: %s", err)
	}

	_ = resource.Expire(60)

	tb.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			tb.Fatalf("Couldn't purge container: %v", err)
		}
	})

	dsn.Host = fmt.Sprintf("%s:5432", resource.Container.NetworkSettings.IPAddress)
	if runtime.GOOS == "darwin" { // MacOS-specific
		dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		tb.Fatalf("Couldn't open DB: %s", err)
	}

	defer db.Close()

	if err := pool.Retry(func() (err error) {
		return db.Ping()
	}); err != nil {
		tb.Fatalf("Couldn't ping DB: %s", err)
	}

	//-

	instance, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		tb.Fatalf("Couldn't migrate (1): %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migrations/", "postgres", instance)
	if err != nil {
		tb.Fatalf("Couldn't migrate (2): %s", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		tb.Fatalf("Couldnt' migrate (3): %s", err)
	}

	//-

	dbpool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		tb.Fatalf("Couldn't open DB Pool: %s", err)
	}

	tb.Cleanup(func() {
		dbpool.Close()
	})

	return dbpool
}
