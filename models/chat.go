package models

// ChatRequest チャットリクエストの構造体
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// ChatResponse チャットレスポンスの構造体
type ChatResponse struct {
	Reply string `json:"reply"`
}

// LoginRequest ログインリクエストの構造体
type LoginRequest struct {
	Password string `json:"password" binding:"required"`
}

// LoginResponse ログインレスポンスの構造体
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
