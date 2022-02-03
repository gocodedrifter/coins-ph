package main

import (
	"flag"
	"fmt"
	"github.com/coins-ph/cmd/internal"
	"github.com/coins-ph/internal/domains/account"
	"github.com/coins-ph/internal/domains/payment"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort              = "8080"
)

func main() {

	godotenv.Load("../../.env")
	var (
		addr  = envString("PORT", defaultPort)

		httpAddr          = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	pool, err := internal.NewPostgreSQL()
	if err != nil {
		fmt.Println("error due to ", err)
	}
	accountService := account.NewAccountService(pool)
	paymentService := payment.NewPaymentService(pool)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	mux.Handle("/wallet/v1/account", account.MakeHandler(accountService, httpLogger))
	mux.Handle("/wallet/v1/payment", payment.MakeHandler(paymentService, httpLogger))
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
