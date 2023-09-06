package response

type ApiInfoResp struct {
	QueryPath           *QueryPath    `json:"query_path"`
	EditUid             int           `json:"edit_uid"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	ReqBodyIsJsonSchema bool          `json:"req_body_is_json_schema"`
	ResBodyIsJsonSchema bool          `json:"res_body_is_json_schema"`
	ApiOpened           bool          `json:"api_opened"`
	Index               int           `json:"index"`
	Tag                 []interface{} `json:"tag"`
	Id                  int           `json:"_id"`
	Method              string        `json:"method"`
	CatId               int           `json:"catid"`
	Title               string        `json:"title"`
	Path                string        `json:"path"`
	ProjectId           int           `json:"project_id"`
	ReqParams           []interface{} `json:"req_params"`
	ResBodyType         string        `json:"res_body_type"`
	Uid                 int           `json:"uid"`
	AddTime             int           `json:"add_time"`
	UpTime              int           `json:"up_time"`
	ReqQuery            []interface{} `json:"req_query"`
	ReqHeaders          []*ReqHeader  `json:"req_headers"`
	ReqBodyForm         []interface{} `json:"req_body_form"`
	V                   int           `json:"__v"`
	Desc                string        `json:"desc"`
	Markdown            string        `json:"markdown"`
	ReqBodyOther        string        `json:"req_body_other"`
	ReqBodyType         string        `json:"req_body_type"`
	ResBody             string        `json:"res_body"`
	Username            string        `json:"username"`
}

type QueryPath struct {
	Path   string        `json:"path"`
	Params []interface{} `json:"params"`
}

type ReqHeader struct {
	Required string `json:"required"`
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}
