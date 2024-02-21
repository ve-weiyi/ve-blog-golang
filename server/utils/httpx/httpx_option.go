package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Client 是一个HTTP客户端
type Client struct {
	httpClient *http.Client
	timeout    time.Duration

	//method  string
	//rawURL  string                 // http://www.baidu.com
	params  map[string]string      // ?a=1&b=2
	headers map[string]string      // map[Content-Type:application/x-www-form-urlencoded]
	body    map[string]interface{} // {"a":1,"b":2}
}

// NewClient 创建一个新的HTTP客户端
func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{},
		timeout:    30 * time.Second,
		params:     make(map[string]string),
		headers:    make(map[string]string),
		body:       make(map[string]interface{}),
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// 发送请求
func (c *Client) DoRequest(method string, rawURL string) (respBody []byte, err error) {
	uv, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest(method, uv.String(), nil)
	if err != nil {
		return nil, err
	}

	// 设置头部
	c.setHeaders(req)

	// 设置查询参数
	c.setQueryParams(req)

	// 设置请求体
	c.setBody(req)

	log.Println("requestUrl:", req.URL.String())
	// 执行请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request fail. url:%v, code:%d,err:%s", req.URL.String(), resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) setHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

func (c *Client) setQueryParams(req *http.Request) {
	query := req.URL.Query()
	for key, value := range c.params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
}

func (c *Client) setBody(req *http.Request) {
	str, err := json.Marshal(c.body)
	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(str))
}

func (c *Client) EncodeURL(rawURL string) string {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return rawURL
	}

	c.setHeaders(req)
	c.setQueryParams(req)
	c.setBody(req)

	return req.URL.String()
}

// 发送GET请求
func (c *Client) Get(url string) (respBody []byte, err error) {
	return c.DoRequest(http.MethodGet, url)

}

// 发送POST请求
func (c *Client) Post(url string) (respBody []byte, err error) {
	return c.DoRequest(http.MethodPost, url)

}
