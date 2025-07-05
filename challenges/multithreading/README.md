# Go Multithreading

> [!IMPORTANT]  
> To run the project contained in this repository, you need to have Go installed on your computer. For more information, visit <https://go.dev/>.

## GoLang Post GoExpert Challenge - Multithreading

This project is part of the Post GoExpert challenge, covering knowledge in Go Routines, channels, contexts, and error handling.

The challenge consists of delivering a system in Go that performs a search for ZIP codes (CEP) in two distinct APIs and returns the result from the API that responds the fastest. The APIs to be used are as follows:

```plaintext
A - https://brasilapi.com.br/api/cep/v1/01153000 + cep
B - http://viacep.com.br/ws/ + cep + /json
```

The following requirements must be met:

- Both requests must be made simultaneously;
- Accept the response from the fastest API and discard the slower response;
- Limit the response time to 1 second. Otherwise, a timeout error must be displayed.

## Extras Added to the Challenge

For best practices, I decided to add some extra points to the exercise, such as:

- ZIP code validation:
  > Validates if the information has the length of a ZIP code (8) and consists only of digits.

- Loading URLs from environment variables:
  > Instead of hardcoding the API URLs directly, I made them configurable via environment variables. This way, in case of a change or API replacement, as long as the contract is maintained, there is no need to modify the code.

### Running the System

To run the system, simply execute the command below, always passing a valid ZIP code as an argument.

```shell
 go run main.go 76195000
```

In the terminal window, you should see a message similar to the example below:

```shell
 go run main.go 21831430
API BrasilAPI :: response (86.366792ms) - {Cep:21831430 Logradouro:Rua Júlio Conceição Complemento: Bairro:Senador Camará Localidade:Rio de Janeiro Uf:RJ}
```