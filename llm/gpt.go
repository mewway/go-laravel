package llm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-openai"
)

type Gpt struct {
	client           *openai.Client
	Messages         []openai.ChatCompletionMessage
	Functions        []openai.FunctionDefinition
	Stream           bool
	Temperature      float32
	TopP             float32
	PresencePenalty  float32
	FrequencyPenalty float32
	Model            string
}

func NewGpt(config config.Config) *Gpt {
	gpt := Gpt{
		Temperature:      0.5,
		TopP:             1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		Model:            config.GetString("gpt.model", openai.GPT4),
	}
	server := config.GetString("gpt.server", "http://chat.cds8.cn/api/openai/v1")
	// 传入 token 会导致配置的自定义服务地址失效
	token := config.GetString("gpt.api_key", "")

	conf := openai.ClientConfig{
		BaseURL:              server,
		OrgID:                "",
		APIType:              openai.APITypeOpenAI,
		APIVersion:           "",
		AzureModelMapperFunc: nil,
		HTTPClient:           &http.Client{},
		EmptyMessagesLimit:   300,
	}
	if token != "" {
		conf = openai.DefaultConfig(token)
	}
	gpt.client = openai.NewClientWithConfig(conf)
	return &gpt
}

func (g Gpt) ChatCompletion() {
	var resp, err = g.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:            g.Model,
		Messages:         g.Messages,
		Functions:        g.Functions,
		Temperature:      g.Temperature,
		FrequencyPenalty: g.FrequencyPenalty,
		PresencePenalty:  g.PresencePenalty,
		Stream:           g.Stream,
		TopP:             g.TopP,
	})
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
