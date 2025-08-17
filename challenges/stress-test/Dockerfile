FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/st cmd/cli/main.go

# ----------------------------

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bin/st .

ENTRYPOINT [ "./st" ]
