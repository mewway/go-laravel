package response

type ApiListResp struct {
	List  []*ApiItem
	Total int
	Count int
}

type ApiItem struct {
	EditUid   int           `json:"edit_uid"`
	Status    string        `json:"status"`
	ApiOpened bool          `json:"api_opened"`
	Tag       []interface{} `json:"tag"`
	Id        int           `json:"_id"`
	Method    string        `json:"method"`
	CatId     int           `json:"catid"`
	Title     string        `json:"title"`
	Path      string        `json:"path"`
	ProjectId int           `json:"project_id"`
	Uid       int           `json:"uid"`
	AddTime   int           `json:"add_time"`
}
