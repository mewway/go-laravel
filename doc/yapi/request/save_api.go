package request

type SaveApiReq struct {
	Token        string        `json:"token"`
	ReqQuery     []interface{} `json:"req_query"`
	ReqHeaders   []*ReqHeader  `json:"req_headers"`
	ReqBodyForm  []interface{} `json:"req_body_form"`
	Title        string        `json:"title"`
	CatId        string        `json:"catid"`
	Path         string        `json:"path"`
	Status       string        `json:"status"`
	ResBodyType  string        `json:"res_body_type"`
	ResBody      string        `json:"res_body"`
	SwitchNotice bool          `json:"switch_notice"`
	Message      string        `json:"message"`
	Desc         string        `json:"desc"`
	Method       string        `json:"method"`
	ReqParams    []interface{} `json:"req_params"`
	Id           string        `json:"id"`
}

type ReqHeader struct {
	Name string `json:"name"`
}
