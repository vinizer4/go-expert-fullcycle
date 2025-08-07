package main

import (
	"cep-weather/model"
	"os"
	"testing"
)

// Mock for ZipCodeResponse
func mockZipCodeHTTP(uri string, method string) (model.ZipCodeResponse, error, int) {
	switch uri {
	case "https://viacep.com.br/ws/12345678/json":
		return model.ZipCodeResponse{City: "São Paulo"}, nil, 200
	case "https://viacep.com.br/ws/00000000/json":
		return model.ZipCodeResponse{}, nil, 404
	default:
		return model.ZipCodeResponse{}, nil, 422
	}
}

// Mock for WeatherResponse
func mockWeatherHTTP(uri string, method string) (model.WeatherResponse, error, int) {
	return model.WeatherResponse{Current: struct {
		TemperatureCelsius float64 `json:"temp_c"`
	}(struct{ TemperatureCelsius float64 }{TemperatureCelsius: 25.0})}, nil, 200
}

func TestIsValidZipCode(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345678", true},
		{"abcdefgh", false},
		{"1234567", false},
		{"123456789", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := isValidZipCode(tt.input); got != tt.expected {
			t.Errorf("isValidZipCode(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestFetchCityFromCEP(t *testing.T) {
	city, err, status := fetchCityFromCEP("12345678", mockZipCodeHTTP)
	if city != "São Paulo" || err != nil || status != 200 {
		t.Errorf("fetchCityFromCEP valid failed: got (%v, %v, %v)", city, err, status)
	}
	_, err, status = fetchCityFromCEP("00000000", mockZipCodeHTTP)
	if err == nil || status != 404 {
		t.Errorf("fetchCityFromCEP not found failed: got (%v, %v)", err, status)
	}
	_, err, status = fetchCityFromCEP("invalid", mockZipCodeHTTP)
	if err == nil || status != 422 {
		t.Errorf("fetchCityFromCEP invalid failed: got (%v, %v)", err, status)
	}
}

func TestFetchWeather(t *testing.T) {
	os.Setenv("API_KEY", "testkey")
	temp, err, status := fetchWeather("São Paulo", mockWeatherHTTP)
	if temp != 25.0 || err != nil || status != 200 {
		t.Errorf("fetchWeather valid failed: got (%v, %v, %v)", temp, err, status)
	}
	os.Unsetenv("API_KEY")
	_, err, status = fetchWeather("São Paulo", mockWeatherHTTP)
	if err == nil || status != 400 {
		t.Errorf("fetchWeather no API key failed: got (%v, %v)", err, status)
	}
}
