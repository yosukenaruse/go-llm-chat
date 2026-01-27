package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	APIKey   string
	Password string
}

func LoadEnv() Env {
	// ローカル開発向け：.env があれば読み込む（本番は環境変数を使う）
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	apiKey := os.Getenv("API_KEY")
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "default_password" // 本番環境では必ず上書きする
	}

	return Env{
		APIKey:   apiKey,
		Password: password,
	}
}
