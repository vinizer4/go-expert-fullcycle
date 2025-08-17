# Go Expert Challenge - Stress Test CLI

Implementation of a CLI in Go to perform stress tests on a web address.

## Architecture

Requests are made concurrently according to the specified amount and are distributed across a pool of workers. Each worker is responsible for making an HTTP request and storing the result in a communication channel. The result is then processed and displayed at the end of execution.

## How to run

### Locally

To run the project locally, you need to have Go installed on your machine. After installation, simply execute the command below, replacing the values of `--url`, `--requests`, and `--concurrency` as desired.

```sh
go run cmd/cli/main.go \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

## Tests

To run the unit tests, simply execute the command below.

```sh
make test
```