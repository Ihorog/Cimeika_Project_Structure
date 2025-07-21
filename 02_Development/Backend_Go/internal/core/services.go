package core

import (
	"cimeika-backend/internal/integrations"
	"cimeika-backend/internal/services"
)

// Services holds all application services
type Services struct {
	Config       *Config
	AIService    *services.AIService
	CalendarService *services.CalendarService
	FileService  *services.FileService
	WeatherService *services.WeatherService
	HealthService *services.HealthService
	TelegramService *services.TelegramService
	AstrologyService *services.AstrologyService
}

// NewServices creates a new services instance
func NewServices(config *Config) *Services {
	// Initialize integrations
	openAIClient := integrations.NewOpenAIClient(config.OpenAIAPIKey)
	googleClient := integrations.NewGoogleClient(config.GoogleCalendarAPIKey)
	dropboxClient := integrations.NewDropboxClient(config.DropboxAccessToken)
	weatherClient := integrations.NewWeatherClient(config.OpenWeatherAPIKey)
	healthClient := integrations.NewHealthClient(config.HealthAPIKey)
	telegramClient := integrations.NewTelegramClient(config.TelegramBotToken)
	astrologyClient := integrations.NewAstrologyClient(config.AstrologyAPIKey)

	// Initialize services
	aiService := services.NewAIService(openAIClient)
	calendarService := services.NewCalendarService(googleClient)
	fileService := services.NewFileService(dropboxClient)
	weatherService := services.NewWeatherService(weatherClient)
	healthService := services.NewHealthService(healthClient)
	telegramService := services.NewTelegramService(telegramClient)
	astrologyService := services.NewAstrologyService(astrologyClient)

	return &Services{
		Config:          config,
		AIService:       aiService,
		CalendarService: calendarService,
		FileService:     fileService,
		WeatherService:  weatherService,
		HealthService:   healthService,
		TelegramService: telegramService,
		AstrologyService: astrologyService,
	}
}