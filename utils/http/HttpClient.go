package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient 结构体封装了HTTP请求的一些基本信息
type HTTPClient struct {
	BaseURL   string
	Client    *http.Client
	HeaderMap http.Header
}

// NewHTTPClient 创建一个新的HTTPClient实例
func NewHTTPClient(baseURL string) *HTTPClient {
	client := &http.Client{
		Timeout: time.Second * 30, // 设置请求超时时间为30秒
	}
	return &HTTPClient{
		BaseURL:   baseURL,
		Client:    client,
		HeaderMap: make(http.Header),
	}
}

// AddHeader 添加HTTP头信息
func (c *HTTPClient) AddHeader(key, value string) {
	c.HeaderMap.Set(key, value)
}

// DoGet 发起一个GET请求
func (c *HTTPClient) DoGet(path string, query url.Values) ([]byte, error) {
	reqURL := c.BaseURL + path
	if query != nil && len(query) > 0 {
		reqURL += "?" + query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	// 添加全局Header
	req.Header = c.HeaderMap.Clone()

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	return body, nil
}

// DoPostJSON DoPost 发起一个POST请求，body为JSON格式
func (c *HTTPClient) DoPostJSON(path string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", "application/json")
	req.Header = c.HeaderMap.Clone()

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("HTTP POST request failed with status code %d", resp.StatusCode)
	}

	return body, nil
}

// DoPostFormData 发起一个POST请求，body为form-data格式
func (c *HTTPClient) DoPostFormData(path string, data url.Values) ([]byte, error) {
	// 将数据转换为form-data格式
	formData := data.Encode()
	body := strings.NewReader(formData)

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}

	// 设置Content-Type
	req.Header = c.HeaderMap.Clone()

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return bodyBytes, fmt.Errorf("HTTP POST request failed with status code %d", resp.StatusCode)
	}

	return bodyBytes, nil
}
