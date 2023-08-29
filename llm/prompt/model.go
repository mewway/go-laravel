// Package prompt
// @Description: 模型层的 prompt 包括模型的批量初始化、批量迁移、存量变更归档
package prompt

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/llm"
	"github.com/mewway/go-openai"
)

func ConstResolver(fields string) (messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, callback llm.Callback) {
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

func TableCheckResolver(ddl, countMap string) (messages []openai.ChatCompletionMessage, functions []openai.FunctionDefinition, callback llm.Callback) {
	messages = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("现在存在一个表的建表语句为：%s，对应字段名和去重后的数量 map 的 json 为： %s, 根据要求调用工具并返回", ddl, countMap),
		},
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `# 你现在是一个 mysql 专家，你会根据我给的建表语句判断字段类型是否使用恰当，索引设置是否正确，并且按照你认为的优先级划分推荐修改的等级，并给出修改的 SQL 语句
# 我会向你提供建表的 DDL 语句，并且会以 map 的形式向你提供每一个字段的去重后数据行数的数量（记住这不是数据的长度），Id 的唯一值即为表的总行数
# 你需要处理分析并按照要求返回：
	1. 你要根据注释说明判断字段是否存在枚举值，必须在字段存在枚举值的时候才会判断：枚举值的数量和我给定的去重数量是否匹配，不匹配要说明
	2. 如果字段类型的选用不恰当，你应该说明并给出修正的建议
	3. 如果索引重复、或者索引定义在了效果不佳的字段上，你也需要明确说明，并支出修正建议
	4. 你要能根据字段命名的含义，判断其长度配置和字段类型选用是否正确
	5. 判断索引缺失的场景
	6. 判断默认值的配置问题，判断时间字段的选用是否存在风险或者异常
	7. 你的推荐修改等级为
		1. Error: 错误（急切需要修改，不修改会导致错误发生）这个级别会返回 SQL 修正的建议
		2. Warning: 警告（最好按照建议修改），这个级别会返回 SQL 修正的建议
	8. 你的修正建议要尽可能言简意赅，并且按照Error、Warning的紧急程度进行排序后返回
	9. 针对字段命名拼写错误给出修正建议，不要使用 mysql 中的不常用的数据类型如： enum、blob等
    10. 每次仅返回你认为最为优先级较高的不超过 5 条的修改建议

# 你的返回需要符合以下风格：
	- 【Info】字段【foo】使用了 varchar(32) 类型，建议结合实际场景选用更大的长度
	- 【Warning】字段【url】使用了  text 类型，修改为varchar(256)，【SQL】ALTER TABLE xxxx Modify xxxx
	- 【Error】索引【foo】 与 【bar】重复, 建议删除索引【foo】 【SQL】xxx
	- 【Error】索引【foo】  效果不佳, 字段去重后数量过低，建议修改字段或结合实际场景判断
最后按照建议的数量, 调用resolver工具，返回到结果 results数组中, 一个建议一条`,
		},
	}

	functions = append(functions, openai.FunctionDefinition{
		Name:        "tableChecker",
		Description: "根据给定数据库信息返回优化的建议",
		Parameters: llm.Parameters{
			Type: llm.TypeObject,
			Properties: map[string]any{
				"results": map[string]any{
					"type": llm.TypeArray,
					"items": map[string]any{
						"type": llm.TypeObject,
						"properties": map[string]any{
							"level": map[string]any{
								"description": "推荐修改等级",
								"type":        llm.TypeString,
							},
							"suggest": map[string]any{
								"description": "修正建议",
								"type":        llm.TypeString,
							},
							"sql": map[string]any{
								"description": "建议修正的 SQL 语句（如果有）",
								"type":        llm.TypeString,
							},
						},
					},
				},
			},
			Required: []string{"level", "suggest"},
		},
	})

	callback = tableCheckerHandler
	return
}

type TableCheckResponseBody struct {
	Results []*Suggestion `json:"results"`
}

type Suggestion struct {
	Level   string `json:"level"`
	Sql     string `json:"sql"`
	Suggest string `json:"suggest"`
}

func tableCheckerHandler(reply string, replyType string) (resp string) {
	r := new(TableCheckResponseBody)
	err := json.Unmarshal([]byte(reply), r)
	if err != nil {
		return ""
	}
	pattern := "【%s】%s 【SQL】%s"
	for _, v := range r.Results {
		switch strings.ToLower(v.Level) {
		case "error":
			color.Errorln(fmt.Sprintf(pattern, v.Level, v.Suggest, v.Sql))
		case "warning":
			color.Warnln(fmt.Sprintf(pattern, v.Level, v.Suggest, v.Sql))
		default:
			color.Infoln(fmt.Sprintf(pattern, v.Level, v.Suggest, v.Sql))
		}
	}
	return
}
