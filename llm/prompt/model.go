// Package prompt
// @Description: 模型层的 prompt 包括模型的批量初始化、批量迁移、存量变更归档
package prompt

import (
	"fmt"

	"github.com/mewway/go-laravel/llm"
	"github.com/mewway/go-openai"
)

func ConstResolver(fields string) (messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition) {
	messages = []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `# 你现在是一个 golang 的编程助手，你负责在模型初始化的时候定义常量
# 我会给你所有的数据库字段名和它的注释说明,每一个字段独立一行，格式如： 【字段名】注释说明
# 你要按照下列步骤处理并返回：
	1. 判断每一行数据是否存在枚举值，至少存在两个枚举你才处理，不存在就跳过
	2. 存在枚举值，你会在常量定义时会根据来源的语言转换成对应的英文
	3. 转换的格式需要按照格式：字段名 + 枚举名 = 枚举值 // 注释说明
	4. 不同的字段名的常量定义一定要记得分组定义成多个 const (...) 的结构
	5. 常量的命名需要按照驼峰格式，并且首字母大写
	6. 返回的数据只能包含 const 定义代码，不可以有多余的 package 等语法结构`,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("定义 golang 的 const：%s", fields),
		},
	}

	functions = append(functions, openai.FunctionDefinition{
		Name:        llm.CommonFuncName,
		Description: "定义 golang的 const",
		Parameters: llm.Parameters{
			Type: llm.TypeObject,
			Properties: map[string]map[string]string{
				llm.CommonFuncField: {
					"description": "代码块:形如 const (...),不可以有多余的 package 等语法",
					"type":        llm.TypeString,
				},
			},
			Required: []string{llm.CommonFuncField},
		},
	})
	return
}
