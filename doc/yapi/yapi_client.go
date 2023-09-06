package yapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/doc/yapi/request"
	"github.com/mewway/go-laravel/doc/yapi/response"
	"github.com/mewway/go-laravel/facades"
)

type Client struct {
	Token      string
	Server     string
	HttpClient *http.Client
}

func NewClient() *Client {
	conf := facades.Config()
	t := conf.GetString("doc.yapi.token")
	s := conf.GetString("doc.yapi.server")
	if t == "" || s == "" {
		color.Errorln("Config doc.yapi.token or doc.yapi.server can't not be empty")
		return nil
	}
	c := new(http.Client)
	return &Client{
		Token:      t,
		Server:     s,
		HttpClient: c,
	}
}

func (c Client) QueryApiList() *response.ApiListResp {
	api := "/api/interface/list"
	s := fmt.Sprintf("%s/%s?token=%s&page=1&limit=1000", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"), c.Token)
	resp, err := c.HttpClient.Get(s)
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
	r := &struct {
		ErrCode int                   `json:"errcode"`
		ErrMsg  string                `json:"errmsg"`
		Data    *response.ApiListResp `json:"data"`
	}{}
	err = json.Unmarshal(body, r)
	if err != nil {
		color.Errorln("Request body parse failed: " + err.Error())
		return nil
	}
	return r.Data
}

func (c Client) QueryApiInfo() *response.ApiInfoResp {
	api := "/api/interface/get"
	s := fmt.Sprintf("%s/%s?token=%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"), c.Token)
	resp, err := c.HttpClient.Get(s)
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
	data := new(response.ApiInfoResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		color.Errorln("Request body parse failed: " + err.Error())
		return nil
	}
	return data
}

func (c Client) SaveApi(req *request.SaveApiReq) (err error) {
	api := "/api/interface/save"
	s := fmt.Sprintf("%s/%s?token=%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"), c.Token)
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := c.HttpClient.Post(s, "application/json", bytes.NewBuffer(d))
	if err != nil {
		color.Errorln("Request failed:" + err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Errorln("Request body read failed: " + err.Error())
		return err
	}
	data := new(response.ApiInfoResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		color.Errorln("Request body parse failed: " + err.Error())
		return err
	}
	return nil
}
