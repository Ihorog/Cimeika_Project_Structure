package handlers

import (
	"cimeika-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatHandler handles AI chat requests
type ChatHandler struct {
	aiService *services.AIService
}

// NewChatHandler creates a new chat handler
func NewChatHandler(aiService *services.AIService) *ChatHandler {
	return &ChatHandler{aiService: aiService}
}

// ChatRequest represents a chat request
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
	Persona string `json:"persona,omitempty"`
	Context string `json:"context,omitempty"`
}

// ChatResponse represents a chat response
type ChatResponse struct {
	Response string `json:"response"`
	Persona  string `json:"persona"`
}

// Handle processes chat requests
func (h *ChatHandler) Handle(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default persona to "ci" if not specified
	if req.Persona == "" {
		req.Persona = "ci"
	}

	response, err := h.aiService.Chat(req.Message, req.Persona, req.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process chat request"})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{
		Response: response,
		Persona:  req.Persona,
	})
}

// TransformHandler handles persona transformation requests
type TransformHandler struct {
	aiService *services.AIService
}

// NewTransformHandler creates a new transform handler
func NewTransformHandler(aiService *services.AIService) *TransformHandler {
	return &TransformHandler{aiService: aiService}
}

// TransformRequest represents a transformation request
type TransformRequest struct {
	FromPersona string `json:"from_persona" binding:"required"`
	ToPersona   string `json:"to_persona" binding:"required"`
	Context     string `json:"context,omitempty"`
}

// TransformResponse represents a transformation response
type TransformResponse struct {
	Message     string `json:"message"`
	NewPersona  string `json:"new_persona"`
	Transformed bool   `json:"transformed"`
}

// Handle processes transformation requests
func (h *TransformHandler) Handle(c *gin.Context) {
	var req TransformRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.aiService.Transform(req.FromPersona, req.ToPersona, req.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process transformation"})
		return
	}

	c.JSON(http.StatusOK, TransformResponse{
		Message:     message,
		NewPersona:  req.ToPersona,
		Transformed: true,
	})
}