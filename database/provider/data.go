package provider

type DataProvider interface {
	SetServer(server string)
	Authorization() *Authorization
	Alias() map[string]string
	Api(api string, args ...[]any)
	GetRequest() BaseRequest
	Execute() BaseResponse
	SetRelation(...any)
}

type Authorization struct {
	Headers map[string]string
	Query   map[string]string
	Json    map[string]string
	Body    map[string]string
}

type BaseResponse interface {
}

type BaseRequest interface {
}

var AliasApiMap = map[string]string{}
var ApiProviderMap = map[string]*DataProvider{}

type DataProviderImpl struct {
	Name   string
	Server string
	DataProvider
}
