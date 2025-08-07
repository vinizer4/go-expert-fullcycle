package main

import (
	"cep-weather/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// isValidZipCode checks if the given zipCode is a valid 8-digit number.
func isValidZipCode(zipCode string) bool {
	return regexp.MustCompile(`^\d{8}$`).MatchString(zipCode)
}

// fetchCityFromCEP retrieves the city name for a given CEP (Brazilian postal code).
// It returns the city name, an error if any occurred, and an HTTP status code.
func fetchCityFromCEP(
	cep string,
	httpFunc func(string, string) (model.ZipCodeResponse, error, int),
) (string, error, int) {
	if !isValidZipCode(cep) {
		return "", errors.New("invalid zipcode"), http.StatusUnprocessableEntity
	}
	uri := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep)
	apiResponse, err, status := httpFunc(uri, http.MethodGet)
	if err != nil {
		return "", err, status
	}
	if apiResponse.City == "" || status == http.StatusNotFound {
		return "", errors.New("can not find zipcode"), http.StatusNotFound
	}
	return apiResponse.City, nil, status
}

// fetchWeather retrieves the current temperature for a given city.
// It returns the temperature in Celsius, an error if any occurred, and an HTTP status code.
func fetchWeather(
	city string,
	httpFunc func(string, string) (model.WeatherResponse, error, int),
) (float64, error, int) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return 0, errors.New("no API key set"), http.StatusBadRequest
	}
	encodedCity := url.QueryEscape(city)
	uri := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=pt", apiKey, encodedCity)
	apiResponse, err, status := httpFunc(uri, http.MethodGet)
	if err != nil {
		return 0, err, status
	}
	return apiResponse.Current.TemperatureCelsius, nil, status
}

// makeHTTPRequest performs an HTTP request and unmarshals the response into the specified type T.
// It returns the unmarshaled response, an error if any occurred, and an HTTP status code.
func makeHTTPRequest[T any](uri string, method string) (T, error, int) {
	var result T
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return result, fmt.Errorf("error creating request: %w", err), http.StatusInternalServerError
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("error sending request: %w", err), http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode), resp.StatusCode
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("error parsing response: %w", err), http.StatusInternalServerError
	}
	return result, nil, http.StatusOK
}

// handleRequest is the main HTTP handler for the CEP weather API.
// It processes the incoming request, fetches the city and weather data, and returns the response.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	city, err, status := fetchCityFromCEP(cep, makeHTTPRequest[model.ZipCodeResponse])
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	temperature, err, status := fetchWeather(city, makeHTTPRequest[model.WeatherResponse])
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	tempResponse := model.TemperatureData{
		Celsius:    model.Float64Marshal(temperature),
		Fahrenheit: model.Float64Marshal(temperature*1.8 + 32),
		Kelvin:     model.Float64Marshal(temperature + 273),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempResponse)
}

func main() {
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
