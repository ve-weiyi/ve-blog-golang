package httpx

import (
	"encoding/json"
	"time"
)

// Option 用于配置Client的选项
type Option func(*Client)

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

// WithHeader 设置请求头
func WithHeader(key string, value string) Option {
	return func(c *Client) {
		c.headers[key] = value
	}
}

// WithParam 设置查询参数
func WithParam(key string, value string) Option {
	return func(c *Client) {
		c.params[key] = value
	}
}

// WithData 设置请求体
func WithData(key string, value interface{}) Option {
	return func(c *Client) {
		c.body[key] = value
	}
}

// // WithBody 设置请求体
func WithBody(obj interface{}) Option {
	return func(c *Client) {
		m, err := structToMap(obj)
		if err != nil {
			return
		}

		for k, v := range m {
			c.body[k] = v
		}
	}
}

// WithHeaders 设置请求头
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) {
		for key, value := range headers {
			c.headers[key] = value
		}
	}
}

// WithParams 设置查询参数
func WithParams(params map[string]string) Option {
	return func(c *Client) {
		for key, value := range params {
			c.params[key] = value
		}
	}
}

func structToMap(obj interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func jsonToMap(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
