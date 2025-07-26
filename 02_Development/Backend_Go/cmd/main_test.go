package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cimeika-backend/internal/api"
	"cimeika-backend/internal/core"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create test configuration
	config := &core.Config{
		Port: "8080",
	}

	// Create test services
	services := core.NewServices(config)

	// Setup router
	router := gin.New()
	api.SetupRoutes(router, services)

	// Create test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])
	assert.Equal(t, "1.0.0", response["version"])
	assert.Equal(t, "cimeika-backend-go", response["service"])
}

func TestAPIRoutes(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create test configuration
	config := &core.Config{
		Port: "8080",
	}

	// Create test services
	services := core.NewServices(config)

	// Setup router
	router := gin.New()
	api.SetupRoutes(router, services)

	testCases := []struct {
		method   string
		endpoint string
		expected int
	}{
		{"GET", "/api/v1/weather/current", http.StatusOK},
		{"GET", "/api/v1/calendar/events", http.StatusOK},
		{"GET", "/api/v1/health/metrics", http.StatusOK},
		{"GET", "/api/v1/astrology/horoscope", http.StatusOK},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tc.method, tc.endpoint, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.expected, w.Code, "Failed for %s %s", tc.method, tc.endpoint)
	}
}