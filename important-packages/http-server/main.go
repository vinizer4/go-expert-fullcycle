package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/find-cep", FindCEP)
	http.ListenAndServe(":8080", nil)
}

func FindCEP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/find-cep" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		http.Error(writer, "cep is required", http.StatusBadRequest)
		return
	}
	cep, error := findCep(cepParam)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(cep)
	if err != nil {
		http.Error(writer, "cep is required", http.StatusBadRequest)
		return
	}
}

func findCep(cep string) (*ViaCEP, error) {
	response, error := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	var viacep ViaCEP
	error = json.Unmarshal(body, &viacep)
	if error != nil {
		return nil, error
	}
	return &viacep, nil
}
