all: check-env
	docker-compose up --build -d

down:
	docker-compose down

check-env:
	@grep -q '^AUCTION_INTERVAL=' ./cmd/auction/.env || (echo "AUCTION_INTERVAL is not set in .env file"; exit 1)

test:
	go test -v ./internal/infra/database/auction
