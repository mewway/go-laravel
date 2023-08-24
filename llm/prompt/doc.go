// Package prompt
// @Description: 文档类的 prompt，包括生成和解析文档两个部分
package prompt

import (
	"fmt"

	"github.com/mewway/go-laravel/llm"
	"github.com/mewway/go-openai"
)

func DocGenerate(fields string) (messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition) {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf("定义 golang 的 const：%s", fields),
	}, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "生成 golang的const 定义代码块, 需要注释和分组有助于良好阅读,分成多个const组,变量定义使用驼峰形式,首字母大写,不可以有多余的 package 等语法结构,然后调用constResolver工具",
	})

	functions = append(functions, openai.FunctionDefinition{
		Name:        "constResolver",
		Description: "定义 golang的 const",
		Parameters: llm.Parameters{
			Type: llm.TypeObject,
			Properties: map[string]map[string]string{
				"code": {
					"description": "代码块:形如 const (...),不可以有多余的 package 等语法",
					"type":        llm.TypeString,
				},
			},
			Required: []string{"code"},
		},
	})
	return
}

func DocParse(fields string) (messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition) {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf("定义 golang 的 const：%s", fields),
	}, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "生成 golang的const 定义代码块, 需要注释和分组有助于良好阅读,分成多个const组,变量定义使用驼峰形式,首字母大写,不可以有多余的 package 等语法结构,然后调用constResolver工具",
	})

	functions = append(functions, openai.FunctionDefinition{
		Name:        "constResolver",
		Description: "定义 golang的 const",
		Parameters: llm.Parameters{
			Type: llm.TypeObject,
			Properties: map[string]map[string]string{
				"code": {
					"description": "代码块:形如 const (...),不可以有多余的 package 等语法",
					"type":        llm.TypeString,
				},
			},
			Required: []string{"code"},
		},
	})
	return
}
