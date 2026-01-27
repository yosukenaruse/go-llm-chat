package handlers

import (
	"go-llm-chat/models"
	"go-llm-chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatHandler チャットハンドラーの構造体
type ChatHandler struct {
	chatService *services.ChatService
}

// NewChatHandler ChatHandlerのコンストラクタ
func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

// HandleChat チャットリクエストを処理
func (h *ChatHandler) HandleChat(c *gin.Context) {
	var req models.ChatRequest

	// リクエストボディをバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// ChatServiceを使用してレスポンスを取得
	response, err := h.chatService.GetResponse(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get response",
		})
		return
	}

	// レスポンスを返却
	c.JSON(http.StatusOK, response)
}
