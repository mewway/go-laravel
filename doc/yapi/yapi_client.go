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
	c := new(http.Client)
	return &Client{
		Token:      t,
		Server:     s,
		HttpClient: c,
	}
}

func (c Client) QueryApiList() *response.ApiListResp {
	api := "/api/interface/list"
	resp, err := c.HttpClient.Get(fmt.Sprintf("%s/%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/")))
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
	data := new(response.ApiListResp)
	err = json.Unmarshal(body, data)
	if err != nil {
		color.Errorln("Request body parse failed: " + err.Error())
		return nil
	}
	return data
}

func (c Client) QueryApiInfo() *response.ApiInfoResp {
	api := "/api/interface/get"
	resp, err := c.HttpClient.Get(fmt.Sprintf("%s/%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/")))
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
	s := fmt.Sprintf("%s/%s", strings.TrimRight(c.Server, "/"), strings.TrimLeft(api, "/"))
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := c.HttpClient.Post(s, "application/json", bytes.NewBuffer(d))
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
	return nil
}
