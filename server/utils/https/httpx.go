package https

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpConnector struct {
	//method  string            // GET POST PUT DELETE
	baseUrl string                 // http://www.baidu.com
	params  url.Values             // ?a=1&b=2
	headers map[string]string      // map[Content-Type:application/x-www-form-urlencoded]
	data    map[string]interface{} // {"a":1,"b":2}
}

func NewHttpConnector(baseUrl string, opts ...Option) *HttpConnector {
	uv, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		log.Println(err)
		return nil
	}
	urls := strings.SplitN(uv.String(), "?", 2)

	builder := &HttpConnector{
		baseUrl: urls[0],
		params:  uv.Query(),
		headers: map[string]string{},
		data:    map[string]interface{}{},
	}

	return builder
}

func (h *HttpConnector) DoRequest(method string) (respBody []byte, err error) {

	requestUrl := h.GetUrl()
	headers := h.headers
	data := h.GetBody()

	req, err := http.NewRequest(method, requestUrl, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept-Charset", "UTF-8")
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request fail. url:%v,code:%d,err:%s", requestUrl, resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (h *HttpConnector) GetUrl() string {
	if h.baseUrl == "" {
		return ""
	}
	if len(h.params) == 0 {
		return h.baseUrl
	}
	return h.baseUrl + "?" + h.params.Encode()
}

func (h *HttpConnector) GetBody() string {
	str, err := json.Marshal(h.data)
	if err != nil {
		return ""
	}

	return string(str)
}

// 请求头
func (h *HttpConnector) AddHeader(key string, value interface{}) *HttpConnector {
	if key == "" {
		return h
	}
	h.headers[key] = fmt.Sprint(value)
	return h
}

// 查询参数 https://www.baidu.com?a=1&b=2
func (h *HttpConnector) AddParam(key string, value interface{}) *HttpConnector {
	if key == "" {
		return h
	}
	h.params.Add(key, fmt.Sprint(value))
	return h
}

// 请求体 {"a":1,"b":2}
func (h *HttpConnector) AddData(key string, value interface{}) *HttpConnector {

	h.data[key] = value
	return h
}

// 请求体 {"a":1,"b":2}
func (h *HttpConnector) AddBody(obj interface{}) *HttpConnector {
	m, err := structToMap(obj)
	if err != nil {
		return h
	}

	for k, v := range m {
		h.data[k] = v
	}
	return h
}

func (h *HttpConnector) Post() ([]byte, error) {
	return h.DoRequest("POST")
}

func (h *HttpConnector) Get() ([]byte, error) {
	return h.DoRequest("GET")
}
