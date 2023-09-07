package response

type ApiItem struct {
	Key      string     `json:"key"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Children []*ApiItem `json:"children"`
	Folder   *Folder    `json:"folder,omitempty"`
	Api      *Api       `json:"api,omitempty"`
}

type Folder struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	DocId    int    `json:"docId"`
	ParentId int    `json:"parentId"`
	Type     string `json:"type"`
}

type Api struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	Path            string           `json:"path"`
	ResponsibleId   int              `json:"responsibleId"`
	Tags            []interface{}    `json:"tags"`
	FolderId        int              `json:"folderId"`
	Type            string           `json:"type,omitempty"`
	Method          string           `json:"method,omitempty"`
	CustomApiFields *CustomApiFields `json:"customApiFields,omitempty"`
}

type CustomApiFields struct {
}
