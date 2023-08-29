package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gookit/color"
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
	callback         Callback
}

type Parameters struct {
	Type       string   `json:"type"`
	Items      any      `json:"items,omitempty"`
	Properties any      `json:"properties,omitempty"`
	Required   []string `json:"required"`
}

type DescriptionType struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type CommonResolverBody struct {
	Text string `json:"text"`
}

type Callback func(reply string, replyType string) (resp string)

const (
	TypeObject  = "object"
	TypeString  = "string"
	TypeArray   = "array"
	TypeNumber  = "number"
	TypeBoolean = "bool"
)

const (
	CommonFuncName  = "resolver"
	CommonFuncField = "text"
)

func NewGpt(config config.Config) *Gpt {
	gpt := Gpt{
		Temperature:      0,
		TopP:             0,
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

func (g *Gpt) ChatCompletion(remark string) (reply string) {
	tips := "正在调用大模型..."
	if remark != "" {
		tips = fmt.Sprintf("【%s】大模型调用中...", remark)
	}
	color.Grayln(tips)

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
		color.Errorln(err)

		return ""
	}
	if err != nil {
		color.Redln(err.Error())
		return
	}
	if len(resp.Choices) == 0 {
		color.Redln("大模型返回的结果为空")
		return
	}
	// 函数调用
	if resp.Choices[0].Message.FunctionCall != nil {
		switch resp.Choices[0].Message.FunctionCall.Name {
		// 默认的文本提取器
		case CommonFuncName:
			g.callback = commonCallback
			reply = g.callback(resp.Choices[0].Message.FunctionCall.Arguments, resp.Choices[0].Message.FunctionCall.Name)
		default:
			reply = resp.Choices[0].Message.FunctionCall.Arguments
			if g.callback != nil {
				reply = g.callback(resp.Choices[0].Message.FunctionCall.Arguments, resp.Choices[0].Message.FunctionCall.Name)
			}
		}
		return
	}
	if g.callback != nil {
		reply = g.callback(resp.Choices[0].Message.Content, "")
		return
	}
	reply = resp.Choices[0].Message.Content
	return
}

func (g *Gpt) Resolve(messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, callback Callback) *Gpt {
	g.Messages = messages
	g.Functions = functions
	g.callback = callback
	return g
}

func commonCallback(reply string, replyType string) (resp string) {
	commonBody := &CommonResolverBody{}
	err := json.Unmarshal([]byte(reply), commonBody)
	if err != nil {
		color.Redln(err)
		return
	}
	resp = commonBody.Text
	return
}
