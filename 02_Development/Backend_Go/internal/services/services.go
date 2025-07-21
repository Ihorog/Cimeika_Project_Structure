package services

import (
	"cimeika-backend/internal/integrations"
	"fmt"
)

// AIService handles AI-related operations
type AIService struct {
	openAIClient *integrations.OpenAIClient
	personas     map[string]string
}

// NewAIService creates a new AI service
func NewAIService(openAIClient *integrations.OpenAIClient) *AIService {
	personas := map[string]string{
		"ci": `You are Ci, the central assistant of the Cimeika platform. 
Your role is to be an intelligent helper that assists users in organizing their life through planning, mood management, and creativity. 
Respond in Ukrainian language. Be friendly, supportive, and adaptive to the user's needs.`,
		
		"podiya": `You are ПоДія, a specialized persona for event and calendar management. 
You help users organize their schedule, plan events, and manage their time effectively. 
Respond in Ukrainian language. Be organized, proactive, and detail-oriented.`,
		
		"nastriy": `You are Настрій, a specialized persona for mood and wellness management. 
You help users understand their emotions, provide support, and suggest wellness activities. 
Respond in Ukrainian language. Be empathetic, calming, and encouraging.`,
		
		"malya": `You are Маля, a specialized persona for creative activities with children. 
You help with games, art projects, and educational activities for kids. 
Respond in Ukrainian language. Be playful, imaginative, and child-friendly.`,
		
		"kazkar": `You are Казкар, a specialized persona for storytelling. 
You create and tell stories, help with creative writing, and engage in narrative activities. 
Respond in Ukrainian language. Be creative, engaging, and narrative-focused.`,
	}
	
	return &AIService{
		openAIClient: openAIClient,
		personas:     personas,
	}
}

// Chat processes a chat message with the specified persona
func (s *AIService) Chat(message, persona, context string) (string, error) {
	systemPrompt, exists := s.personas[persona]
	if !exists {
		systemPrompt = s.personas["ci"] // Default to Ci
	}
	
	// Add context if provided
	if context != "" {
		systemPrompt += "\n\nAdditional context: " + context
	}
	
	return s.openAIClient.Chat(systemPrompt, message)
}

// Transform handles persona transformation
func (s *AIService) Transform(fromPersona, toPersona, context string) (string, error) {
	transformMessage := fmt.Sprintf("Transforming from %s to %s persona.", fromPersona, toPersona)
	
	if context != "" {
		transformMessage += " Context: " + context
	}
	
	// In a real implementation, this would handle the actual transformation logic
	return transformMessage, nil
}

// CalendarService handles calendar-related operations
type CalendarService struct {
	googleClient *integrations.GoogleClient
}

// NewCalendarService creates a new calendar service
func NewCalendarService(googleClient *integrations.GoogleClient) *CalendarService {
	return &CalendarService{googleClient: googleClient}
}

// FileService handles file-related operations
type FileService struct {
	dropboxClient *integrations.DropboxClient
}

// NewFileService creates a new file service
func NewFileService(dropboxClient *integrations.DropboxClient) *FileService {
	return &FileService{dropboxClient: dropboxClient}
}

// WeatherService handles weather-related operations
type WeatherService struct {
	weatherClient *integrations.WeatherClient
}

// NewWeatherService creates a new weather service
func NewWeatherService(weatherClient *integrations.WeatherClient) *WeatherService {
	return &WeatherService{weatherClient: weatherClient}
}

// HealthService handles health-related operations
type HealthService struct {
	healthClient *integrations.HealthClient
}

// NewHealthService creates a new health service
func NewHealthService(healthClient *integrations.HealthClient) *HealthService {
	return &HealthService{healthClient: healthClient}
}

// TelegramService handles Telegram-related operations
type TelegramService struct {
	telegramClient *integrations.TelegramClient
}

// NewTelegramService creates a new telegram service
func NewTelegramService(telegramClient *integrations.TelegramClient) *TelegramService {
	return &TelegramService{telegramClient: telegramClient}
}

// AstrologyService handles astrology-related operations
type AstrologyService struct {
	astrologyClient *integrations.AstrologyClient
}

// NewAstrologyService creates a new astrology service
func NewAstrologyService(astrologyClient *integrations.AstrologyClient) *AstrologyService {
	return &AstrologyService{astrologyClient: astrologyClient}
}