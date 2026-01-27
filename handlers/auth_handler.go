package handlers

import (
	"go-llm-chat/models"
	"go-llm-chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler 認証ハンドラーの構造体
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler AuthHandlerのコンストラクタ
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// HandleLogin ログインリクエストを処理
func (h *AuthHandler) HandleLogin(c *gin.Context) {
	var req models.LoginRequest

	// リクエストボディをバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.LoginResponse{
			Success: false,
			Message: "Invalid request format",
		})
		return
	}

	// パスワードを検証
	err := h.authService.ValidatePassword(req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.LoginResponse{
			Success: false,
			Message: "Invalid password",
		})
		return
	}

	// 認証成功時、Cookieに認証トークンを設定
	c.SetCookie(
		"auth_token",        // Cookie名
		"authenticated",     // Cookie値
		3600*24,            // 有効期限（秒）- 24時間
		"/",                // パス
		"",                 // ドメイン（空文字の場合は現在のドメイン）
		false,              // Secure（HTTPSのみ）
		true,               // HttpOnly（JavaScriptからアクセス不可）
	)

	// 成功レスポンスを返却
	c.JSON(http.StatusOK, models.LoginResponse{
		Success: true,
		Message: "Login successful",
	})
}

// HandleLogout ログアウトリクエストを処理
func (h *AuthHandler) HandleLogout(c *gin.Context) {
	// Cookieを削除（有効期限を過去に設定）
	c.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successful",
	})
}

// AuthMiddleware 認証ミドルウェア
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cookieから認証トークンを取得
		authToken, err := c.Cookie("auth_token")

		// トークンが存在しない、または値が正しくない場合
		if err != nil || authToken != "authenticated" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
