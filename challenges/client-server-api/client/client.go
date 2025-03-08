package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func FetchCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func SaveCotacaoToFile(cotacao *Cotacao) error {
	content := fmt.Sprintf("Dólar: %s", cotacao.Bid)
	return ioutil.WriteFile("cotacao.txt", []byte(content), 0644)
}

func RunClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotacao, err := FetchCotacao(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = SaveCotacaoToFile(cotacao)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cotação salva com sucesso!")
}
