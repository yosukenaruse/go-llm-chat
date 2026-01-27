package services

import (
	"context"
	"fmt"
	"go-llm-chat/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// ChatService チャットサービスの構造体
type ChatService struct {
	llm *openai.LLM
}

// NewChatService ChatServiceのコンストラクタ
func NewChatService() *ChatService {
	// .envファイルを読み込み
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found")
	}

	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set in environment variables")
	}

	// DeepSeek用のOpenAI互換LLMクライアントを作成
	// DeepSeekはOpenAI互換のAPIを提供しているため、OpenAIクライアントを使用
	llm, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithBaseURL("https://api.deepseek.com/v1"), // DeepSeekのエンドポイント
		openai.WithModel("deepseek-chat"),                  // DeepSeekのモデル
	)
	if err != nil {
		log.Fatalf("Failed to create LLM client: %v", err)
	}

	return &ChatService{
		llm: llm,
	}
}

// GetResponse メッセージに対するレスポンスを取得
func (s *ChatService) GetResponse(req *models.ChatRequest) (*models.ChatResponse, error) {
	ctx := context.Background()

	// LLMにメッセージを送信
	response, err := llms.GenerateFromSinglePrompt(
		ctx,
		s.llm,
		req.Message,
		llms.WithTemperature(0.7), // 応答の創造性を調整（0.0-1.0）
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %w", err)
	}

	return &models.ChatResponse{
		Reply: response,
	}, nil
}
