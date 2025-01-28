package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/busca-cep", BuscaCEP)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/busca-cep" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		http.Error(writer, "cep is required", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Buscando CEP"))
}
