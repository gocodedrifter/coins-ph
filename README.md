# COINS-PH MICROSERVICE

## Introduction


The Project developed use Domain Driven Design with Go Kit.

## Project Structure

Below are the structure details.

- [X] `build/`: defines the code used for creating infrastructure as well as docker containers.
    - [X] `wallet/Dockerfile`
- [X] `cmd/`
    - [X] `internal/postgresql.go`: connection to postgresql database
    - [X] `wallet/main.go`: main class to run the api
- [X] `db/`
    - [X] `migrations/`: contains database migrations.
    - [X] `seeds/`: contains file meant to populate basic database values.
- [X] `internal/`: defines the _core domain_.
    - [X] `dataservice/`: a concrete _repository_ used by the domain, for this project using `postgresql`
    - [X] `domains/`
        - [X] `account/`: all operation related to account
        - [X] `payment/`: all operation related to payment

## How To Run The Application using Docker 

Below are the steps to run the application
* Run 
    * `docker-compose build wallet`
    * `docker-compose run wallet migrate -path /api/migrations/ -database postgres://postgres:password@db-coins:5432?sslmode=disable up`

## API Documentation

[API ](./doc/API.md#L15-L22)
