package usecase

import (
	"context"
	"go-llm-chat/models"
)

// LLMClient は外部LLMへの依存を抽象化する。
// usecase層は具体実装（DeepSeek/OpenAI互換など）を知らない。
type LLMClient interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

// ChatUsecase はチャットのユースケース。
type ChatUsecase interface {
	GetResponse(ctx context.Context, req *models.ChatRequest) (*models.ChatResponse, error)
}

// ChatInteractor は ChatUsecase の実装。
type ChatInteractor struct {
	llm LLMClient
}

func NewChatInteractor(llm LLMClient) *ChatInteractor {
	return &ChatInteractor{llm: llm}
}

func (u *ChatInteractor) GetResponse(ctx context.Context, req *models.ChatRequest) (*models.ChatResponse, error) {
	reply, err := u.llm.Generate(ctx, req.Message)
	if err != nil {
		return nil, err
	}
	return &models.ChatResponse{Reply: reply}, nil
}
