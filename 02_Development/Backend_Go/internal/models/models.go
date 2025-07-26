package models

import "time"

// Event represents a calendar event
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// File represents an uploaded file
type File struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	URL         string    `json:"url"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

// WeatherData represents weather information
type WeatherData struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Timestamp   time.Time `json:"timestamp"`
}

// HealthMetrics represents health and wellness data
type HealthMetrics struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Type        string    `json:"type"` // "mood", "energy", "sleep", etc.
	Value       float64   `json:"value"`
	Notes       string    `json:"notes,omitempty"`
	RecordedAt  time.Time `json:"recorded_at"`
}

// Horoscope represents astrological data
type Horoscope struct {
	Sign        string `json:"sign"`
	Date        string `json:"date"`
	Prediction  string `json:"prediction"`
	LuckyNumber int    `json:"lucky_number,omitempty"`
	LuckyColor  string `json:"lucky_color,omitempty"`
}

// TelegramMessage represents a Telegram message
type TelegramMessage struct {
	ChatID   int64  `json:"chat_id"`
	Text     string `json:"text"`
	Username string `json:"username,omitempty"`
	SentAt   time.Time `json:"sent_at"`
}

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatContext represents conversation context
type ChatContext struct {
	UserID      string    `json:"user_id"`
	Persona     string    `json:"persona"`
	Messages    []string  `json:"messages"`
	LastUpdated time.Time `json:"last_updated"`
}