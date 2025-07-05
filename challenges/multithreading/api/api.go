package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type ViaCEPAddress struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type BrasilAPIAddress struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func convertBrasilAPIToAddress(b BrasilAPIAddress) Address {
	return Address{
		Cep:        b.Cep,
		Logradouro: b.Street,
		Bairro:     b.Neighborhood,
		Localidade: b.City,
		Uf:         b.State,
	}
}

func convertViaCEPToAddress(v ViaCEPAddress) Address {
	return Address(v)
}

type apiConverterFunc func([]byte) (Address, error)

var apiConverters = map[string]apiConverterFunc{
	"BrasilAPI": convertBrasilAPIResponse,
	"ViaCEP":    convertViaCEPResponse,
}

func convertBrasilAPIResponse(body []byte) (Address, error) {
	var brasilAPIAddress BrasilAPIAddress
	err := json.Unmarshal(body, &brasilAPIAddress)
	if err != nil {
		return Address{}, err
	}
	return convertBrasilAPIToAddress(brasilAPIAddress), nil
}

func convertViaCEPResponse(body []byte) (Address, error) {
	var viaCEPAddress ViaCEPAddress
	err := json.Unmarshal(body, &viaCEPAddress)
	if err != nil {
		return Address{}, err
	}
	return convertViaCEPToAddress(viaCEPAddress), nil
}

func FetchAddressFromAPI(ctx context.Context, url string, name string, ch chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Erro ao criar requisição em %s :: %v", name, err)
		return
	}

	startTime := time.Now()
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Erro ao buscar endereço em %s :: %v", name, err)
		return
	}
	defer res.Body.Close()
	duration := time.Since(startTime)

	if res.StatusCode != http.StatusOK {
		ch <- fmt.Sprintf("Busca em %s, com erro na resposta :: %s | %d bytes - %s", name, res.Status, res.ContentLength, duration.String())
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- fmt.Sprintf("Erro lendo resposta da API(%s) :: %v", name, err)
		return
	}

	converter, exists := apiConverters[name]
	if !exists {
		ch <- fmt.Sprintf("Conversor para a API %s não encontrado.", name)
		return
	}

	address, err := converter(body)
	if err != nil {
		ch <- fmt.Sprintf("Erro na conversão da resposta da API %s :: %v", name, err)
		return
	}

	if err := json.Unmarshal(body, &address); err != nil {
		ch <- fmt.Sprintf("Erro na conversão da resposta da API(%s) :: %v", name, err)
		return
	}
	ch <- fmt.Sprintf("API %s :: resposta (%s) - %+v", name, duration.String(), address)
}
