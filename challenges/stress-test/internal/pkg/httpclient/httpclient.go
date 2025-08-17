package httpclient

import (
	"net/http"
	"time"
)

type HttpClientInterface interface {
	Get(endpoint string) *HttpClientResponse
}

type HttpClientResponse struct {
	StatusCode *int
	Duration   time.Duration
	Error      error
}

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c HttpClient) Get(addr string) *HttpClientResponse {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return &HttpClientResponse{
			Error: err,
		}
	}

	client := &http.Client{}

	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return &HttpClientResponse{
			StatusCode: &resp.StatusCode,
			Error:      err,
		}
	}

	return &HttpClientResponse{
		StatusCode: &resp.StatusCode,
		Duration:   time.Since(start),
	}
}
