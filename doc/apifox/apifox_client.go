package apifox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/doc/apifox/request"
	"github.com/mewway/go-laravel/doc/apifox/response"
	"github.com/mewway/go-laravel/facades"
)

type Client struct {
	Authorization string
	ProjectId     string
	Locale        string
	Server        string
	HttpClient    *http.Client
}

func NewClient() *Client {
	conf := facades.Config()
	bearer := conf.GetString("doc.apifox.authorization")
	s := conf.GetString("doc.apifox.server")
	pId := conf.GetString("doc.apifox.project_id")
	l := conf.GetString("doc.apifox.locale", "zh-CN")
	if bearer == "" || s == "" || pId == "" {
		color.Errorln("Config doc.apifox.authorization or doc.apifox.server or apifox.project_id can't not be empty")
		return nil
	}
	c := new(http.Client)
	return &Client{
		Authorization: bearer,
		Server:        s,
		HttpClient:    c,
		ProjectId:     pId,
		Locale:        l,
	}
}

func (c Client) QueryApiList() *response.ApiListResp {
	api := "/api/v1/api-details"
	s := fmt.Sprintf("%s/%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"))
	q := url.Values{}
	q.Set("locale", c.Locale)
	s = s + "?" + q.Encode()

	req, _ := http.NewRequest("GET", s, nil)
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("X-Project-Id", c.ProjectId)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		color.Errorln("Request failed:" + err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Errorln("Request body read failed: " + err.Error())
		return nil
	}
	r := new(response.ApiListResp)
	err = json.Unmarshal(body, r)
	if err != nil {
		color.Errorln("Request body parse failed: " + err.Error())
		return nil
	}
	return r
}

func (c Client) QueryApiInfo(id int) *response.ApiInfo {
	lst := c.QueryApiList()
	for _, v := range lst.Data {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (c Client) SaveApi(req request.SaveApiReq) (err error) {
	// 新增
	api := "/api/v1/api-details"
	// 更新
	api = "/api/v1/api-details/%s"

	_ = fmt.Sprintf("%s/%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"))

	return nil
}
