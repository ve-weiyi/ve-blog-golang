package httpx

import (
	"net/http"
	"testing"
)

func TestNewClientBuilder(t *testing.T) {
	// Create a new HTTP httpClient with options using Builder pattern
	resp, err := NewClientBuilder(
		http.MethodGet,
		"https://baidu.com",
	).
		WithHeaders(map[string]string{"Content-Type": "application/json"}).
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}).
		WithBody([]byte(`{"key": "value"}`)).
		Build().
		DoRequest()

	t.Log(string(resp), err)
}

func TestNewClientOptions(t *testing.T) {
	// Create a new HTTP httpClient with options using Option pattern
	resp, err := NewClient(
		http.MethodGet,
		"https://baidu.com",
		WithHeaders(map[string]string{"Content-Type": "application/json"}),
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}),
		WithBody([]byte(`{"key": "value"}`)),
	).DoRequest()

	t.Log(string(resp), err)
}
