package core

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Port   string
	
	// OpenAI Configuration
	OpenAIAPIKey string
	
	// Google Calendar Configuration
	GoogleCalendarAPIKey string
	
	// Dropbox Configuration
	DropboxAccessToken string
	
	// Weather Configuration
	OpenWeatherAPIKey string
	
	// Astrology Configuration
	AstrologyAPIKey string
	
	// Health Configuration
	HealthAPIKey string
	
	// Telegram Configuration
	TelegramBotToken string
	
	// Database Configuration
	DatabaseURL string
	
	// Redis Configuration
	RedisURL string
	
	// JWT Secret
	JWTSecret string
	
	// Environment
	Environment string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Port:                 getEnv("PORT", "8000"),
		OpenAIAPIKey:         getEnv("OPENAI_API_KEY", ""),
		GoogleCalendarAPIKey: getEnv("GOOGLE_CALENDAR_API_KEY", ""),
		DropboxAccessToken:   getEnv("DROPBOX_ACCESS_TOKEN", ""),
		OpenWeatherAPIKey:    getEnv("OPENWEATHER_API_KEY", ""),
		AstrologyAPIKey:      getEnv("ASTROLOGY_API_KEY", ""),
		HealthAPIKey:         getEnv("HEALTH_API_KEY", ""),
		TelegramBotToken:     getEnv("TELEGRAM_BOT_TOKEN", ""),
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		RedisURL:             getEnv("REDIS_URL", ""),
		JWTSecret:            getEnv("JWT_SECRET", ""),
		Environment:          getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}