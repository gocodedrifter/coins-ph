FROM golang:1.16.3-alpine3.13 AS builder

RUN apk --no-cache add \
      alpine-sdk && \
    rm -rf /var/cache/apk/*

WORKDIR /build/

COPY . .
WORKDIR /build/cmd/wallet
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/wallet

RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo -ldflags "-s -w" -tags 'postgres' \
      github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

#-

FROM alpine:3.13 AS certificates

RUN apk --no-cache add ca-certificates

#-

FROM scratch

WORKDIR /api/
ENV PATH=/api/bin/:$PATH

COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /go/bin/wallet /go/bin/wallet
COPY --from=builder /go/bin/migrate ./bin/migrate
COPY --from=builder /build/.env .
COPY --from=builder /build/db/ .

ENV PORT 8080
EXPOSE 8080
CMD ["migrate", "-path", "/api/migration", "-database", "postgres://postgres:password@db-coins:5432?sslmode=disable"]
ENTRYPOINT ["/go/bin/wallet", "-env", "/api/.env"]
