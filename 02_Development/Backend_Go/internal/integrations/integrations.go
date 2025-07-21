package integrations

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// OpenAIClient wraps the OpenAI API client
type OpenAIClient struct {
	client *openai.Client
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient(apiKey string) *OpenAIClient {
	client := openai.NewClient(apiKey)
	return &OpenAIClient{client: client}
}

// Chat sends a chat message to OpenAI
func (c *OpenAIClient) Chat(systemPrompt, userMessage string) (string, error) {
	if c.client == nil {
		return "OpenAI client not configured", nil
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userMessage,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to get chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// GoogleClient wraps Google API client
type GoogleClient struct {
	apiKey string
}

// NewGoogleClient creates a new Google client
func NewGoogleClient(apiKey string) *GoogleClient {
	return &GoogleClient{apiKey: apiKey}
}

// DropboxClient wraps Dropbox API client
type DropboxClient struct {
	accessToken string
}

// NewDropboxClient creates a new Dropbox client
func NewDropboxClient(accessToken string) *DropboxClient {
	return &DropboxClient{accessToken: accessToken}
}

// WeatherClient wraps weather API client
type WeatherClient struct {
	apiKey string
}

// NewWeatherClient creates a new weather client
func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{apiKey: apiKey}
}

// HealthClient wraps health API client
type HealthClient struct {
	apiKey string
}

// NewHealthClient creates a new health client
func NewHealthClient(apiKey string) *HealthClient {
	return &HealthClient{apiKey: apiKey}
}

// TelegramClient wraps Telegram bot API client
type TelegramClient struct {
	botToken string
}

// NewTelegramClient creates a new Telegram client
func NewTelegramClient(botToken string) *TelegramClient {
	return &TelegramClient{botToken: botToken}
}

// AstrologyClient wraps astrology API client
type AstrologyClient struct {
	apiKey string
}

// NewAstrologyClient creates a new astrology client
func NewAstrologyClient(apiKey string) *AstrologyClient {
	return &AstrologyClient{apiKey: apiKey}
}