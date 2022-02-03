package internal

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	// Initialize "pgx".
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewPostgreSQL instantiates the PostgreSQL database using configuration defined in environment variables.
func NewPostgreSQL() (*pgxpool.Pool, error) {

	// XXX: We will revisit this code in future episodes replacing it with another solution
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseSSLMode := os.Getenv("DATABASE_SSLMODE")
	// XXX: -

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
