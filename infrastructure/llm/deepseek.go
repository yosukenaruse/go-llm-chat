package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// DeepSeekClient は DeepSeek(OpenAI互換API) を使う LLM クライアント。
// usecase.LLMClient を満たす。
type DeepSeekClient struct {
	llm         *openai.LLM
	temperature float64
}

func NewDeepSeekClient(apiKey string, temperature float64) (*DeepSeekClient, error) {
	if temperature < 0.0 || temperature > 1.0 {
		return nil, fmt.Errorf("temperature must be between 0.0 and 1.0: %v", temperature)
	}

	llm, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
		openai.WithModel("deepseek-chat"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create LLM client: %w", err)
	}

	return &DeepSeekClient{
		llm:         llm,
		temperature: temperature,
	}, nil
}

func (c *DeepSeekClient) Generate(ctx context.Context, prompt string) (string, error) {
	resp, err := llms.GenerateFromSinglePrompt(
		ctx,
		c.llm,
		prompt,
		llms.WithTemperature(c.temperature),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}
	return resp, nil
}
