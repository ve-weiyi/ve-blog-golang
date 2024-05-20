package httpx

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func TestNewClientBuilder(t *testing.T) {
	// Create a new HTTP httpClient with options using Builder pattern
	resp, err := NewClientBuilder().
		WithTimeout(10*time.Second).
		WithHeaders(map[string]string{"Content-Type": "application/json"}).
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}).
		WithBody([]byte(`{"key": "value"}`)).
		DoRequest("GET", "https://baidu.com")

	t.Log(string(resp), err)
}

func TestNewClientOptions(t *testing.T) {
	// Create a new HTTP httpClient with options using Option pattern
	resp, err := NewClient(
		WithTimeout(10*time.Second),
		WithHeaders(map[string]string{"Content-Type": "application/json"}),
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}),
		WithBody([]byte(`{"key": "value"}`)),
	).DoRequest("GET", "https://baidu.com")

	t.Log(string(resp), err)
}

func TestNewClient(t *testing.T) {
	client := &http.Client{}

	start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2024, 4, 9, 0, 0, 0, 0, time.UTC).Unix() - 1
	request, err := http.NewRequest("POST", "http://120.79.96.124:8340"+"/v1/business/get/sn_mapping", strings.NewReader(jsonconv.ObjectToJson(map[string]any{
		"start_time": start,
		"end_time":   end,
	})))

	request.Header.Set("appid", "ak-acc-sn—mapping")
	request.Header.Set("secret", "acc-sn-1a6d25374f64")

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(respBody))
}
