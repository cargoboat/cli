package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// ManagementClient 管理客户端
type ManagementClient struct {
	addr, username, password string
	//log                      log.Logger
	httpClient *http.Client
	baseURL    *url.URL
}

// NewManagementClient ...
func NewManagementClient(baseURL, username, password string) *ManagementClient {
	client := &ManagementClient{
		httpClient: &http.Client{},
		//log:        log,
		username: username,
		password: password,
	}
	u, _ := url.Parse(baseURL)
	client.baseURL = u

	return client
}

// urlJoin url 拼接
func (m *ManagementClient) urlJoin(uri string) string {
	u := *m.baseURL
	u.Path = path.Join(u.Path, uri)
	return u.String()
}

// do 执行HTTP请求
func (m *ManagementClient) do(req *http.Request) (response *http.Response, err error) {
	req.SetBasicAuth(m.username, m.password)
	response, err = m.httpClient.Do(req)
	return
}

// configItemModel ...
type configItemModel struct {
	Group string `json:"group"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetValue ...
func (m *ManagementClient) SetValue(group, key, value string) (err error) {
	config := configItemModel{
		Group: group,
		Key:   key,
		Value: value,
	}
	var jsonByte []byte
	jsonByte, err = json.Marshal(config)
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, m.urlJoin("/configs"), bytes.NewBuffer(jsonByte))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	var resp *http.Response
	resp, err = m.do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("设置配置错误")
	}
	return nil
}
func (m *ManagementClient) GetConfigList(group string) (value map[string]string, err error) {
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s?prefix=%s", m.urlJoin("/configs"), group), nil)
	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = m.do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("resp status code:%d", resp.StatusCode))
		return
	}
	defer resp.Body.Close()
	jsonDecode := json.NewDecoder(resp.Body)
	value = make(map[string]string)
	err = jsonDecode.Decode(&value)
	if err != nil {
		return
	}
	return
}

// GetKeysList ...
func (m *ManagementClient) GetKeysList(group string) (list []string, err error) {
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s?prefix=%s", m.urlJoin("/keys"), group), nil)
	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = m.do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("resp status code:%d", resp.StatusCode))
		return
	}
	defer resp.Body.Close()
	jsonDecode := json.NewDecoder(resp.Body)
	list = make([]string, 0)
	err = jsonDecode.Decode(&list)
	if err != nil {
		return
	}
	return
}

// GetKeysList ...
func (m *ManagementClient) Delete(key string) (err error) {
	var req *http.Request
	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s?key=%s", m.urlJoin("/configs"), key), nil)
	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = m.do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, rerr := ioutil.ReadAll(resp.Body)
		if rerr != nil {
			err = rerr
			return
		}
		err = errors.New(string(body))
		return
	}
	return
}

// configItem ...
type configItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// configResult ...
type configResult struct {
	Version int64        `json:"version"`
	Configs []configItem `json:"configs"`
}
