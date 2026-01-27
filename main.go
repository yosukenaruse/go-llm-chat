package main

import (
	"go-llm-chat/handlers"
	"go-llm-chat/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルーターを初期化
	r := gin.Default()

	// 静的ファイルの配信設定
	r.Static("/static", "./static")

	// サービスとハンドラーを初期化
	chatService := services.NewChatService()
	chatHandler := handlers.NewChatHandler(chatService)

	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	// Hello Worldエンドポイント
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// 認証エンドポイント（認証不要）
	r.POST("/login", authHandler.HandleLogin)
	r.POST("/logout", authHandler.HandleLogout)

	// 認証が必要なエンドポイント
	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		// チャットエンドポイント（認証が必要）
		protected.POST("/chat", chatHandler.HandleChat)
	}

	// ポート番号を環境変数から取得（Render対応）
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}

	// サーバーを起動
	r.Run(":" + port)
}
