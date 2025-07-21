package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheckResponse represents the health check response
type HealthCheckResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Service   string    `json:"service"`
}

// HealthCheck handles the health check endpoint
func HealthCheck(c *gin.Context) {
	response := HealthCheckResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
		Service:   "cimeika-backend-go",
	}

	c.JSON(http.StatusOK, response)
}