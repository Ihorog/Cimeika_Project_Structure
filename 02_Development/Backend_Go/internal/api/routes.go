package api

import (
	"cimeika-backend/internal/core"
	"cimeika-backend/internal/api/handlers"
	"cimeika-backend/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, services *core.Services) {
	// Add middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	// Health check endpoint
	router.GET("/health", handlers.HealthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// AI Assistant routes
		ai := v1.Group("/ai")
		{
			ai.POST("/chat", handlers.NewChatHandler(services.AIService).Handle)
			ai.POST("/transform", handlers.NewTransformHandler(services.AIService).Handle)
		}

		// Calendar routes
		calendar := v1.Group("/calendar")
		{
			calendar.GET("/events", handlers.NewCalendarHandler(services.CalendarService).GetEvents)
			calendar.POST("/events", handlers.NewCalendarHandler(services.CalendarService).CreateEvent)
			calendar.PUT("/events/:id", handlers.NewCalendarHandler(services.CalendarService).UpdateEvent)
			calendar.DELETE("/events/:id", handlers.NewCalendarHandler(services.CalendarService).DeleteEvent)
		}

		// File management routes
		files := v1.Group("/files")
		{
			files.POST("/upload", handlers.NewFileHandler(services.FileService).Upload)
			files.GET("/download/:id", handlers.NewFileHandler(services.FileService).Download)
			files.DELETE("/:id", handlers.NewFileHandler(services.FileService).Delete)
		}

		// Weather routes
		weather := v1.Group("/weather")
		{
			weather.GET("/current", handlers.NewWeatherHandler(services.WeatherService).GetCurrent)
			weather.GET("/forecast", handlers.NewWeatherHandler(services.WeatherService).GetForecast)
		}

		// Health & wellness routes
		health := v1.Group("/health")
		{
			health.GET("/metrics", handlers.NewHealthHandler(services.HealthService).GetMetrics)
			health.POST("/metrics", handlers.NewHealthHandler(services.HealthService).SaveMetrics)
		}

		// Astrology routes
		astrology := v1.Group("/astrology")
		{
			astrology.GET("/horoscope", handlers.NewAstrologyHandler(services.AstrologyService).GetHoroscope)
			astrology.GET("/compatibility", handlers.NewAstrologyHandler(services.AstrologyService).GetCompatibility)
		}

		// Telegram bot routes
		telegram := v1.Group("/telegram")
		{
			telegram.POST("/webhook", handlers.NewTelegramHandler(services.TelegramService).HandleWebhook)
			telegram.POST("/send-message", handlers.NewTelegramHandler(services.TelegramService).SendMessage)
		}
	}
}