package llm

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

type DeepSeekClient struct {
	llm         *openai.LLM
	temperature float64
}

// 足し算ツールの構造体を定義
type AdditionTool struct{}

// tools.Tool インターフェースを実装
func (t AdditionTool) Name() string { return "add_numbers" }
func (t AdditionTool) Description() string {
	return "2つの数値を足し合わせるツールです。入力は '1+2' のような形式で受け取ります。"
}
func (t AdditionTool) Call(ctx context.Context, input string) (string, error) {
	// 簡易的なパース（Agentが '1+2' と送ってくることを想定）
	fmt.Printf("--- Tool [add_numbers] called with input: %s ---\n", input)
	nums := strings.Split(input, "+")
	if len(nums) != 2 {
		return "", fmt.Errorf("invalid input format")
	}
	a, _ := strconv.ParseFloat(strings.TrimSpace(nums[0]), 64)
	b, _ := strconv.ParseFloat(strings.TrimSpace(nums[1]), 64)
	return fmt.Sprintf("%f", a+b), nil
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
	// 1. ツールのリスト（構造体を使用）
	allTools := []tools.Tool{
		AdditionTool{},
	}

	// 2. Agentの初期化
	// v0.1.14 で最も安定している ZeroShotReactDescription を使用
	executor, err := agents.Initialize(
		c.llm,
		allTools,
		agents.ZeroShotReactDescription,
		agents.WithMaxIterations(3),
	)
	if err != nil {
		return "", fmt.Errorf("failed to initialize agent: %w", err)
	}

	// 3. 実行
	result, err := chains.Run(ctx, executor, prompt)
	if err != nil {
		return "", fmt.Errorf("agent execution failed: %w", err)
	}

	return result, nil
}