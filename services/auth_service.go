package services

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// AuthService 認証サービスの構造体
type AuthService struct {
	password string
}

// NewAuthService AuthServiceのコンストラクタ
func NewAuthService() *AuthService {
	// .envファイルを読み込み
	godotenv.Load()

	// 環境変数からパスワードを取得
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "default_password" // デフォルト値（本番環境では使用しない）
	}

	return &AuthService{
		password: password,
	}
}

// ValidatePassword パスワードを検証
func (s *AuthService) ValidatePassword(inputPassword string) error {
	if inputPassword != s.password {
		return errors.New("invalid password")
	}
	return nil
}
