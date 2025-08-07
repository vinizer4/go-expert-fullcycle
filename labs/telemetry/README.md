# Temperature System by ZIP Code with OpenTelemetry and Zipkin

This project implements a distributed system in Go that queries ZIP codes and returns temperature information, instrumented with **OpenTelemetry (OTEL)** and **Zipkin** for distributed tracing.

## 📋 Overview

The system consists of:

- **Service A**: Receives ZIP code requests and validates the format
- **Service B**: Processes ZIP code, fetches location and temperature
- **OpenTelemetry Collector**: Collects traces from services
- **Zipkin**: Visualizes distributed traces
- **Prometheus**: Collects metrics

## 🏗️ Architecture

```
Client → Service A → Service B → External APIs
                ↓         ↓
            OTEL Collector
                ↓
              Zipkin
```

### Data Flow:
1. **Service A** receives ZIP code via POST
2. Validates format (8 digits)
3. Forwards to **Service B**
4. **Service B** queries:
   - ViaCEP API (location data)
   - OpenWeatherMap API (coordinates + weather)
5. Returns temperatures in Celsius, Fahrenheit, and Kelvin

## 🚀 How to Run

### Prerequisites
- Docker & Docker Compose
- OpenWeatherMap API Key

### 1. Clone and Configure

```bash
git clone <repo-url>
cd go-telemetry
```

### 2. Set Environment Variables

Create a `.env` file:

```bash
OPENWEATHERMAP_API_KEY=your_api_key_here
```

> 📝 **How to get an API Key**: Go to [WeatherAPI](https://openweathermap.org/api) and create a free account.

### 3. Run with Docker Compose

```bash
docker-compose up --build
```

This will start:
- **Service A** at `http://localhost:8080`
- **Service B** at `http://localhost:8081`
- **Zipkin** at `http://localhost:9411`
- **Prometheus** at `http://localhost:9090`
- **OTEL Collector** at `localhost:4317` (gRPC)

## 📊 Testing the System

### Example Request

```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'
```

### Expected Response

```json
{
  "city": "São Paulo",
  "temp_C": 22.5,
  "temp_F": 72.5,
  "temp_K": 295.65
}
```

### Test Scenarios

```bash
# ✅ Valid ZIP code
curl -X POST http://localhost:8080 -H "Content-Type: application/json" -d '{"cep": "01310100"}'

# ❌ Invalid ZIP code (422)
curl -X POST http://localhost:8080 -H "Content-Type: application/json" -d '{"cep": "123"}'

# ❌ ZIP code not found (404)
curl -X POST http://localhost:8080 -H "Content-Type: application/json" -d '{"cep": "00000000"}'
```

## 🔍 Observability with OTEL + Zipkin

### What is OpenTelemetry?
OpenTelemetry is an observability framework that collects traces, metrics, and logs from distributed applications.

### What is Zipkin?
Zipkin is a distributed tracing tool that visualizes how requests flow between services.

### Viewing Traces

1. Access **Zipkin**: http://localhost:9411
2. Click "Run Query" to view traces
3. Click a trace to see details

### Implemented Spans

#### Service A:
- `handle_cep_request`: Main request span
- `validate_cep`: ZIP code format validation
- `forward_to_service_b`: Communication with Service B

#### Service B:
- `handle_temperature_request`: Main span
- `validate_cep`: ZIP code validation
- `fetch_viacep_data`: Fetch ViaCEP data
- `fetch_coordinates`: Fetch coordinates
- `fetch_weather_data`: Fetch weather data

### Span Attributes

Each span contains detailed information:

```
- cep: Queried ZIP code
- city: Found city
- state: State
- geo.latitude/longitude: Coordinates
- temperature.*: Temperatures
- http.status_code: HTTP status
- api.name: External API name
```

## 📈 Metrics with Prometheus

Access: http://localhost:9090

Available metrics:
- HTTP request duration
- Span counters
- Status codes
- External API latency

## 🛠️ Development

### Project Structure

```
├── cmd/
│   ├── service-a/         # Service A
│   └── service-b/         # Service B
├── internal/
│   ├── service-a/         # Service A logic
│   └── service-b/         # Service B logic
├── shared/
│   ├── telemetry/         # OTEL configuration
│   ├── models/            # Shared models
│   └── http-utils/        # HTTP utilities
├── configs/               # Configurations
├── docker-compose.yml     # Orchestration
└── otel-collector-config.yaml
```

### Running Locally (Development)

```bash
# Terminal 1 - OTEL Collector + Zipkin
docker-compose up zipkin otel-collector prometheus

# Terminal 2 - Service B
export OPENWEATHERMAP_API_KEY=your_key
export OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317
go run cmd/service-b/main.go

# Terminal 3 - Service A
export SERVICE_B_URL=http://localhost:8081
export OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317
go run cmd/service-a/main.go
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `OPENWEATHERMAP_API_KEY` | WeatherMap API Key | - |
| `SERVICE_B_URL` | Service B URL | `http://localhost:8081` |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | OTEL endpoint | `otel-collector:4317` |

### Used Ports

| Service | Port | Description |
|---------|------|-------------|
| Service A | 8080 | Main API |
| Service B | 8081 | Internal API |
| Zipkin | 9411 | Web interface |
| Prometheus | 9090 | Metrics |
| OTEL Collector | 4317 | gRPC Receiver |
| OTEL Collector | 4318 | HTTP Receiver |

## 🚨 Troubleshooting

### Common Issues

1. **Error: "OPENWEATHERMAP_API_KEY environment variable is required"**
   - Solution: Set the API key in `.env`

2. **Traces do not appear in Zipkin**
   - Check if all containers are running
   - Confirm OTEL Collector is accessible

3. **Timeouts on external APIs**
   - Check internet connectivity
   - Confirm valid API key

### Useful Logs

```bash
# View OTEL Collector logs
docker-compose logs otel-collector

# View service logs
docker-compose logs service-a service-b
```

## 📚 Used APIs

- **ViaCEP**: https://viacep.com.br/ (ZIP code lookup)
- **OpenWeatherMap**: https://openweathermap.org/ (coordinates + weather)

## 🎯 Observability Features

### ✅ Implemented

- [x] Distributed traces between services
- [x] Detailed spans for each operation
- [x] Context propagation between services
- [x] Automatic HTTP instrumentation
- [x] Custom metrics
- [x] Detailed span attributes
- [x] Error handling with spans
- [x] Visualization in Zipkin

### 🔄 Possible Improvements

- [ ] Correlated structured logs
- [ ] Custom business metrics
- [ ] SLI/SLO-based alerts
- [ ] Smart sampling
- [ ] Grafana dashboard