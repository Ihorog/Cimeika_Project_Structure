# Cimeika Backend - Go Implementation

This is the Go implementation of the Cimeika backend API, providing an alternative to the Python Flask backend.

## Features

- **RESTful API** using Gin framework
- **AI Assistant** with multiple personas (Ci, ПоДія, Настрій, Маля, Казкар)
- **External API Integrations** (OpenAI, Google Calendar, Dropbox, Weather, Astrology, Health)
- **Telegram Bot** integration
- **Docker** support with multi-stage builds
- **Health checks** and monitoring
- **CORS** and middleware support

## Project Structure

```
Backend_Go/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/        # HTTP request handlers
│   │   ├── middleware/      # Middleware components
│   │   └── routes.go        # Route definitions
│   ├── core/
│   │   ├── config.go        # Configuration management
│   │   └── services.go      # Service initialization
│   ├── integrations/        # External API clients
│   ├── models/              # Data models
│   └── services/            # Business logic services
├── pkg/                     # Shared utilities (if needed)
├── Dockerfile              # Docker configuration
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── .env.example            # Environment variables template
└── README.md               # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker (optional)

### Installation

1. Clone the repository and navigate to the Go backend directory:
   ```bash
   cd 02_Development/Backend_Go
   ```

2. Copy the environment template and configure your API keys:
   ```bash
   cp .env.example .env
   # Edit .env with your actual API keys
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

### Running Locally

1. Start the server:
   ```bash
   go run cmd/main.go
   ```

2. The API will be available at `http://localhost:8000`

3. Check health status:
   ```bash
   curl http://localhost:8000/health
   ```

### Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t cimeika-backend-go .
   ```

2. Run the container:
   ```bash
   docker run -p 8000:8000 --env-file .env cimeika-backend-go
   ```

### API Endpoints

#### Health Check
- `GET /health` - Service health status

#### AI Assistant
- `POST /api/v1/ai/chat` - Chat with AI assistant
- `POST /api/v1/ai/transform` - Transform between personas

#### Calendar Management
- `GET /api/v1/calendar/events` - Get calendar events
- `POST /api/v1/calendar/events` - Create new event
- `PUT /api/v1/calendar/events/:id` - Update event
- `DELETE /api/v1/calendar/events/:id` - Delete event

#### File Management
- `POST /api/v1/files/upload` - Upload file
- `GET /api/v1/files/download/:id` - Download file
- `DELETE /api/v1/files/:id` - Delete file

#### Weather
- `GET /api/v1/weather/current` - Get current weather
- `GET /api/v1/weather/forecast` - Get weather forecast

#### Health & Wellness
- `GET /api/v1/health/metrics` - Get health metrics
- `POST /api/v1/health/metrics` - Save health metrics

#### Astrology
- `GET /api/v1/astrology/horoscope` - Get horoscope
- `GET /api/v1/astrology/compatibility` - Get compatibility

#### Telegram Bot
- `POST /api/v1/telegram/webhook` - Telegram webhook
- `POST /api/v1/telegram/send-message` - Send Telegram message

## AI Personas

The system supports multiple AI personas, each with specialized capabilities:

- **Ci** - Central assistant for general tasks
- **ПоДія** - Event and calendar management specialist
- **Настрій** - Mood and wellness management specialist  
- **Маля** - Creative activities with children specialist
- **Казкар** - Storytelling specialist

## Environment Variables

See `.env.example` for all required environment variables. Key variables include:

- `OPENAI_API_KEY` - OpenAI API key for AI features
- `GOOGLE_CALENDAR_API_KEY` - Google Calendar integration
- `TELEGRAM_BOT_TOKEN` - Telegram bot integration
- `PORT` - Server port (default: 8000)

## Development

### Adding New Features

1. Create models in `internal/models/`
2. Add integrations in `internal/integrations/`
3. Implement services in `internal/services/`
4. Create handlers in `internal/api/handlers/`
5. Add routes in `internal/api/routes.go`

### Testing

```bash
go test ./...
```

### Building for Production

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go
```

## Contributing

1. Follow Go conventions and best practices
2. Add tests for new features
3. Update documentation as needed
4. Ensure Docker builds work correctly

## License

Part of the Cimeika Project Structure.