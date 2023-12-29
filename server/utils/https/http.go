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

type HttpBuilder struct {
	//method  string            // GET POST PUT DELETE
	baseUrl string                 // http://www.baidu.com
	params  url.Values             // ?a=1&b=2
	headers map[string]string      // map[Content-Type:application/x-www-form-urlencoded]
	data    map[string]interface{} // {"a":1,"b":2}
}

func NewHttpBuilder(baseUrl string) *HttpBuilder {
	uv, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		log.Println(err)
		return nil
	}
	urls := strings.SplitN(uv.String(), "?", 2)

	builder := &HttpBuilder{
		baseUrl: urls[0],
		params:  uv.Query(),
		headers: map[string]string{},
		data:    map[string]interface{}{},
	}

	return builder
}

func (h *HttpBuilder) DoRequest(method string) (respBody []byte, err error) {

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

func (h *HttpBuilder) GetUrl() string {
	if h.baseUrl == "" {
		return ""
	}
	if len(h.params) == 0 {
		return h.baseUrl
	}
	return h.baseUrl + "?" + h.params.Encode()
}

func (h *HttpBuilder) GetBody() string {
	str, err := json.Marshal(h.data)
	if err != nil {
		return ""
	}

	return string(str)
}

// 请求头
func (h *HttpBuilder) AddHeader(key string, value interface{}) *HttpBuilder {
	if key == "" {
		return h
	}
	h.headers[key] = fmt.Sprint(value)
	return h
}

// 查询参数 https://www.baidu.com?a=1&b=2
func (h *HttpBuilder) AddParam(key string, value interface{}) *HttpBuilder {
	if key == "" {
		return h
	}
	h.params.Add(key, fmt.Sprint(value))
	return h
}

// 请求体 {"a":1,"b":2}
func (h *HttpBuilder) AddData(key string, value interface{}) *HttpBuilder {

	h.data[key] = value
	return h
}

// 请求体 {"a":1,"b":2}
func (h *HttpBuilder) AddBody(obj interface{}) *HttpBuilder {
	m, err := structToMap(obj)
	if err != nil {
		return h
	}

	for k, v := range m {
		h.data[k] = v
	}
	return h
}

func (h *HttpBuilder) Post() ([]byte, error) {
	return h.DoRequest("POST")
}

func (h *HttpBuilder) Get() ([]byte, error) {
	return h.DoRequest("GET")
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
