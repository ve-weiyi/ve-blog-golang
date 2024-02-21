package httpx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// ClientBuilder 用于构建 HTTP 客户端。
type ClientBuilder struct {
	httpClient *http.Client
	timeout    time.Duration
	headers    map[string]string
	params     map[string]string
	body       []byte
}

// NewClientBuilder 创建一个具有默认设置的新 ClientBuilder。
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		httpClient: &http.Client{},
		timeout:    30 * time.Second, // 默认超时时间
		headers:    make(map[string]string),
		params:     make(map[string]string),
	}
}

// WithTimeout 设置 HTTP 请求的超时时间。
func (c *ClientBuilder) WithTimeout(timeout time.Duration) *ClientBuilder {
	c.timeout = timeout
	return c
}

// WithHeaders 设置 HTTP 请求的头部。
func (c *ClientBuilder) WithHeaders(headers map[string]string) *ClientBuilder {
	c.headers = headers
	return c
}

// WithParams 设置 HTTP 请求的参数。
func (c *ClientBuilder) WithParams(params map[string]string) *ClientBuilder {
	c.params = params
	return c
}

// WithBody 设置 HTTP 请求的正文。
func (c *ClientBuilder) WithBody(body []byte) *ClientBuilder {
	c.body = body
	return c
}

// Build 根据构建器的设置创建一个新的 HTTP 客户端。
func (c *ClientBuilder) Build() *http.Client {
	client := &http.Client{
		Timeout: c.timeout,
	}

	return client
}

// DoRequest 执行一个 HTTP 请求。
func (c *ClientBuilder) DoRequest(method, rawURL string) (respBody []byte, err error) {
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

func (c *ClientBuilder) setHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

func (c *ClientBuilder) setQueryParams(req *http.Request) {
	query := req.URL.Query()
	for key, value := range c.params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
}

func (c *ClientBuilder) setBody(req *http.Request) {

	req.Body = io.NopCloser(bytes.NewReader(c.body))
}
