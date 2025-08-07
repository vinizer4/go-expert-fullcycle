Here is your README translated to English:

# Weather API by ZIP Code

This project is part of the Full Cycle post-graduate program and fulfills the challenge of developing a Go system that receives a ZIP code (CEP), identifies the city, and returns the current weather.

## Challenge Objective

Develop a Go system that receives a ZIP code, identifies the city, and returns the current weather (temperature in Celsius, Fahrenheit, and Kelvin). The system must be published on Google Cloud Run.

## Challenge Requirements

- The system must receive a valid 8-digit ZIP code
- It must search for the ZIP code, find the location name, and return the temperatures formatted in Celsius, Fahrenheit, and Kelvin
- Appropriate responses for the following scenarios:
	- Success: HTTP 200 code, with body
  ```json
  {
    "temp_C": 27.2,
    "temp_F": 80.9,
    "temp_K": 300.2
  }
  ```
	- Invalid ZIP code: HTTP 422 code, message "invalid zipcode"
	- ZIP code not found: HTTP 404 code, message "can not find zipcode"
- Deployment on Google Cloud Run

## Project Structure

```
.
├── cmd
│   └── main.go
├── README.md
├── Dockerfile
├── .gitignore
├── model
│   ├── via_cep.go
│   └── weather.go
└── go.mod
```

## Features

- Fetches city information from a ZIP code using the ViaCEP API
- Retrieves weather data for the city using the WeatherAPI
- Converts temperature to Celsius, Fahrenheit, and Kelvin

## Prerequisites

- Go 1.22 or higher
- Docker (for containerization)

## Environment Variables

- `PORT`: The port on which the server will run (default: 8080)
- `API_KEY`: Your WeatherAPI key

## Build and Run

### Local Development

1. Set the required environment variables
2. Run the application:

```bash
go run cmd/main.go
```

### Using Docker

1. Build the Docker image:

```bash
docker build -t api-clima-cep .
```

2. Run the container:

```bash
docker run -p 8080:8080 -e PORT=8080 -e API_KEY=your_api_key api-clima-cep
```

## API Usage

Send a GET request to the root endpoint with the `cep` query parameter:

```
GET /?cep=12345678
```

The API will return a JSON response with temperature data in Celsius, Fahrenheit, and Kelvin.

## Deployment

Deployment is performed automatically on Google Cloud Run via a GitHub Actions workflow. The workflow is triggered on every push to the `master` branch.

### Deployment Configuration

1. Set the following secrets in your GitHub repository:
	- `GCP_SA_KEY`: Google Cloud service account key in JSON format
	- `WEATHER_API_KEY`: WeatherAPI key

2. The workflow performs the following steps:
	- Authenticates with Google Cloud
	- Sets up the Google Cloud SDK
	- Creates a repository in Artifact Registry (if it does not exist)
	- Builds and pushes the Docker image to Artifact Registry
	- Deploys the image to Cloud Run
	- Cleans up old images from Artifact Registry

3. After deployment, the Cloud Run service URL is displayed in the workflow logs.

## Accessing the Deployed Service

The deployed service on Cloud Run is configured to allow unauthenticated access.

### Service URL

The service is available at: https://cep-weather-api-iiuuzoq6ia-rj.a.run.app

Example of request: https://cep-weather-api-iiuuzoq6ia-rj.a.run.app/?cep=76190000