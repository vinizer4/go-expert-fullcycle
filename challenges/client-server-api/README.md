### README

# Client-Server API Challenge

This project is a solution to the challenge of creating a client-server application in Go. The client requests the current USD to BRL exchange rate from the server, which fetches the data from an external API, saves it to a SQLite database, and returns the exchange rate to the client. The client then saves the exchange rate to a file.

## Project Structure

```
client-server-api/
├── client/
│   └── client.go
├── server/
│   └── server.go
├── main.go
└── go.mod
```

## Requirements

- Go 1.20 or later
- SQLite3

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/project.git
   cd project
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

## Running the Project

1. Start the server:

   ```sh
   go run main.go
   ```

   The server will start and listen on port 8080.

2. The client will automatically run and request the exchange rate from the server, then save it to `cotacao.txt`.

## Project Details

- **Server**:
    - Fetches the USD to BRL exchange rate from `https://economia.awesomeapi.com.br/json/last/USD-BRL`.
    - Saves the exchange rate to a SQLite database with a timeout of 10ms.
    - Returns the exchange rate to the client.

- **Client**:
    - Requests the exchange rate from the server with a timeout of 300ms.
    - Saves the exchange rate to `cotacao.txt` in the format `Dólar: {valor}`.

## Notes

- Ensure that the server is running before the client makes a request.
- The server and client use context timeouts to handle delays and ensure timely responses.
