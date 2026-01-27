package handlers

import (
	"context"
	"go-llm-chat/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-llm-chat/usecase"
)

// ChatHandler チャットハンドラーの構造体
type ChatHandler struct {
	chatUsecase usecase.ChatUsecase
}

// NewChatHandler ChatHandlerのコンストラクタ
func NewChatHandler(chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		chatUsecase: chatUsecase,
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

	// usecaseを使用してレスポンスを取得
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	response, err := h.chatUsecase.GetResponse(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get response",
		})
		return
	}

	// レスポンスを返却
	c.JSON(http.StatusOK, response)
}
