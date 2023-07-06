package https

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
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

func (h *HttpBuilder) DoRequest(method string) (string, int) {

	requestUrl := h.GetUrl()
	headers := h.headers
	data := h.GetDatas()

	client := &http.Client{}
	req, err := http.NewRequest(method, requestUrl, strings.NewReader(data))
	if err != nil {
		log.Println(err)
	}

	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept-Charset", "UTF-8")
	}
	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body), resp.StatusCode
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

func (h *HttpBuilder) GetDatas() string {
	str, err := jsoniter.MarshalToString(h.data)
	if err != nil {
		return ""
	}

	return str
}

func (h *HttpBuilder) AddParam(key string, value interface{}) *HttpBuilder {
	if key == "" {
		return h
	}
	h.params.Add(key, fmt.Sprint(value))
	return h
}

func (h *HttpBuilder) AddHeader(key string, value interface{}) *HttpBuilder {
	if key == "" {
		return h
	}
	h.headers[key] = fmt.Sprint(value)
	return h
}

func (h *HttpBuilder) AddData(key string, value interface{}) *HttpBuilder {

	h.data[key] = value
	return h
}

func (h *HttpBuilder) Post() (string, int) {
	return h.DoRequest("POST")
}

func (h *HttpBuilder) Get() (string, int) {
	return h.DoRequest("GET")
}
