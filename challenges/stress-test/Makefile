build:
	@go build -ldflags="-w -s" -o bin/st cmd/cli/main.go

test:
	@./scripts/test.sh

clean:
	@rm -rf ./bin ./tmp ./coverage.txt
