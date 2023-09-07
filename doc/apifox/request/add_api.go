package request

// POST https://api.apifox.cn/api/v1/api-details?locale=zh-CN 新增
// PUT https://api.apifox.cn/api/v1/api-details/108537235?locale=zh-CN 更新
// DELETE https://api.apifox.cn/api/v1/api-details/108537235?locale=zh-CN 删除
/*
新增和更新的参数如下
path: /api/test/1
method: get
name: 1.test
folderId: 0
status: developing
serverId:
responsibleId: 746451
tags: ["dfa","dtt"]
description: 这是 markdown 文档说明
operationId:
sourceUrl:
responses: [{"code":200,"contentType":"json","tempId":"1693978912662","name":"成功","jsonSchema":{"type":"object","properties":{}}}]
responseExamples: []
codeSamples: []
commonParameters: {"query":[],"body":[],"cookie":[],"header":[{"name":"X-Authorization-Cx"},{"name":"X-Platform-Id"}]}
customApiFields: {}
commonResponseStatus: {}
responseId: 0
type: http
parameters: {"query":[{"required":false,"description":"","type":"string","id":"XT2pYYcfY8","example":"1","enable":true,"name":" query"}]}
requestBody: {"type":"multipart/form-data","parameters":[{"required":false,"description":"","type":"string","id":"lhsTdvSQyD","example":"b","enable":true,"name":"a"}],"jsonSchema":{"type":"object","properties":{}}}
responseChildren: ["TEMP.1693978912662"]
auth: {}
advancedSettings: {}
inheritPostProcessors: {}
inheritPreProcessors: {}
preProcessors: []
postProcessors: []
*/

type SaveApiReq struct {
	Path                  string                 `json:"path"`
	Method                string                 `json:"method"`
	Name                  string                 `json:"name"`
	FolderId              int                    `json:"folder_id"`
	Status                string                 `json:"status"`
	ServerId              string                 `json:"server_id"`
	ResponsibleId         int                    `json:"responsible_id"`
	Tags                  []string               `json:"tags"`
	Description           string                 `json:"description"`
	OperationId           string                 `json:"operation_id"`
	SourceUrl             string                 `json:"source_url"`
	Responses             []Response             `json:"responses"`
	ResponseExamples      []string               `json:"response_examples"`
	CodeSamples           []string               `json:"code_samples"`
	CommonParameters      CommonParams           `json:"common_parameters"`
	CustomApiFields       map[string]interface{} `json:"custom_api_fields"`
	CommonResponseStatus  interface{}            `json:"common_response_status"`
	ResponseId            int                    `json:"response_id"`
	Type                  string                 `json:"type"`
	Parameters            Params                 `json:"parameters"`
	RequestBody           ReqBody                `json:"request_body"`
	ResponseChildren      []string               `json:"response_children"`
	Auth                  interface{}            `json:"auth"`
	AdvancedSettings      interface{}            `json:"advanced_settings"`
	InheritPostProcessors bool                   `json:"inherit_post_processors"`
	InheritPreProcessors  bool                   `json:"inherit_pre_processors"`
	PreProcessors         []string               `json:"pre_processors"`
	PostProcessors        []string               `json:"post_processors"`
}

type Response struct {
	Code        int        `json:"code"`
	ContentType string     `json:"content_type"`
	TempId      string     `json:"temp_id"`
	Name        string     `json:"name"`
	JsonSchema  JsonSchema `json:"json_schema"`
}

type JsonSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type CommonParams struct {
	Query  []interface{} `json:"query"`
	Body   []interface{} `json:"body"`
	Cookie []interface{} `json:"cookie"`
	Header []struct {
		Name string `json:"name"`
	} `json:"header"`
}

type Params struct {
	Query []Param `json:"query"`
}

type Param struct {
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Id          string `json:"id"`
	Example     string `json:"example"`
	Enable      bool   `json:"enable"`
	Name        string `json:"name"`
}

type ReqBody struct {
	Type       string     `json:"type"`
	Parameters []Param    `json:"parameters"`
	JsonSchema JsonSchema `json:"json_schema"`
}
