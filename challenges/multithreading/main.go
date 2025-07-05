package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"multithreading/api"
	"multithreading/cep"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Erro ao carregar arquivos do sistema :: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Informe um CEP por favor.")
		return
	}

	cepArg := os.Args[1]
	isValid, err := cep.IsValid(cepArg)
	if err != nil {
		fmt.Printf("Erro enquanto validava o CEP :: %v\n", err)
		return
	}
	if !isValid {
		fmt.Println("CEP inválido.")
		return
	}

	apiAURL := os.Getenv("BRASIL_API_URL")
	apiBURL := os.Getenv("VIACEP_API_URL")
	if apiAURL == "" || apiBURL == "" {
		fmt.Println("As variáveis de ambiente BRASIL_API_URL e VIACEP_API_URL não foram devidamente configuradas.")
		return
	}

	apiA := fmt.Sprintf(apiAURL, cepArg)
	apiB := fmt.Sprintf(apiBURL, cepArg)

	ch := make(chan string, 2)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go api.FetchAddressFromAPI(ctx, apiA, "BrasilAPI", ch)
	go api.FetchAddressFromAPI(ctx, apiB, "ViaCEP", ch)

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Timeout: nenhuma API respondeu a tempo.")
	}
}
