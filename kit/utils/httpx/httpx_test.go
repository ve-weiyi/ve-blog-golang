package httpx

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClientBuilder(t *testing.T) {
	// Create a new HTTP httpClient with options using Builder pattern
	resp, err := NewClientBuilder().
		WithTimeout(10 * time.Second).
		WithHeaders(map[string]string{"Content-Type": "application/json"}).
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}).
		WithBody([]byte(`{"key": "value"}`)).
		WithMethod(http.MethodGet).
		WithURL("https://baidu.com").
		DoRequest()

	t.Log(string(resp), err)
}

func TestNewClientOptions(t *testing.T) {
	// Create a new HTTP httpClient with options using Option pattern
	resp, err := NewClient(
		WithTimeout(10*time.Second),
		WithHeaders(map[string]string{"Content-Type": "application/json"}),
		WithParams(map[string]string{"param1": "value1", "param2": "value2"}),
		WithBody([]byte(`{"key": "value"}`)),
		WithMethod(http.MethodGet),
		WithURL("https://baidu.com"),
	).DoRequest()

	t.Log(string(resp), err)
}
