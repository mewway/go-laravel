package response

import "time"

type ApiListResp struct {
	Success bool       `json:"success"`
	Data    []*ApiInfo `json:"data"`
}

type ApiInfo struct {
	Id                   int           `json:"id"`
	Name                 string        `json:"name"`
	Type                 string        `json:"type"`
	ServerId             string        `json:"serverId"`
	PreProcessors        []interface{} `json:"preProcessors"`
	PostProcessors       []interface{} `json:"postProcessors"`
	InheritPreProcessors struct {
	} `json:"inheritPreProcessors"`
	InheritPostProcessors struct {
	} `json:"inheritPostProcessors"`
	Description string        `json:"description"`
	OperationId string        `json:"operationId"`
	SourceUrl   string        `json:"sourceUrl"`
	Method      string        `json:"method"`
	Path        string        `json:"path"`
	Tags        []interface{} `json:"tags"`
	Status      int           `json:"status"`
	RequestBody struct {
		Type        string         `json:"type"`
		Parameters  []interface{}  `json:"parameters"`
		JsonSchema  map[string]any `json:"jsonSchema,omitempty"`
		Description string         `json:"description,omitempty"`
		Example     string         `json:"example,omitempty"`
	} `json:"requestBody"`
	Parameters struct {
		Path  []interface{} `json:"path,omitempty"`
		Query []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Required    bool   `json:"required"`
			Description string `json:"description"`
			Example     string `json:"example,omitempty"`
			Type        string `json:"type"`
			Enable      bool   `json:"enable,omitempty"`
		} `json:"query,omitempty"`
		Cookie []interface{} `json:"cookie,omitempty"`
		Header []struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			Required bool   `json:"required"`
			Example  string `json:"example"`
			Type     string `json:"type"`
			Enable   bool   `json:"enable,omitempty"`
		} `json:"header,omitempty"`
	} `json:"parameters"`
	CommonParameters struct {
		Query  []interface{} `json:"query,omitempty"`
		Body   []interface{} `json:"body,omitempty"`
		Cookie []interface{} `json:"cookie,omitempty"`
		Header []struct {
			Name string `json:"name"`
		} `json:"header,omitempty"`
	} `json:"commonParameters"`
	Auth struct {
	} `json:"auth"`
	Responses []struct {
		Id            int            `json:"id"`
		CreatedAt     time.Time      `json:"createdAt"`
		UpdatedAt     time.Time      `json:"updatedAt"`
		DeletedAt     interface{}    `json:"deletedAt"`
		ApiDetailId   int            `json:"apiDetailId"`
		Name          string         `json:"name"`
		Code          int            `json:"code"`
		ContentType   string         `json:"contentType"`
		JsonSchema    map[string]any `json:"jsonSchema"`
		DefaultEnable bool           `json:"defaultEnable"`
		ProjectId     int            `json:"projectId"`
		Ordering      int            `json:"ordering"`
	} `json:"responses"`
	ResponseExamples []struct {
		Id          int         `json:"id"`
		CreatedAt   time.Time   `json:"createdAt"`
		UpdatedAt   time.Time   `json:"updatedAt"`
		DeletedAt   interface{} `json:"deletedAt"`
		ApiDetailId int         `json:"apiDetailId"`
		Name        string      `json:"name"`
		ResponseId  int         `json:"responseId"`
		Data        string      `json:"data"`
		Ordering    int         `json:"ordering"`
	} `json:"responseExamples"`
	CodeSamples          []interface{} `json:"codeSamples"`
	ProjectId            int           `json:"projectId"`
	FolderId             int           `json:"folderId"`
	Ordering             int           `json:"ordering"`
	ResponsibleId        int           `json:"responsibleId"`
	CommonResponseStatus struct {
	} `json:"commonResponseStatus"`
	AdvancedSettings struct {
		IsDefaultUrlEncoding int `json:"isDefaultUrlEncoding,omitempty"`
	} `json:"advancedSettings"`
	CustomApiFields any `json:"customApiFields"`
	MockScript      struct {
	} `json:"mockScript"`
	CreatedAt        time.Time   `json:"createdAt"`
	UpdatedAt        time.Time   `json:"updatedAt"`
	DeletedAt        interface{} `json:"deletedAt"`
	CreatorId        int         `json:"creatorId"`
	EditorId         int         `json:"editorId"`
	ResponseChildren []string    `json:"responseChildren"`
}
