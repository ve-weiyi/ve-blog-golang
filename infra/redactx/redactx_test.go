package redactx_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ve-weiyi/blog-cloud/infra/redactx"
)

func TestMarshal(t *testing.T) {
	input := map[string]any{
		"username": "alice",
		"password": "plain-text",
		"nested": map[string]any{
			"old_password": "old",
			"mobile":       "13800000000",
		},
		"list": []any{
			map[string]any{"token": "abc"},
			map[string]any{"code": "123456"},
		},
	}

	raw := redactx.Marshal(input)

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got["password"] != "***" {
		t.Errorf("password = %v, want ***", got["password"])
	}
	if got["username"] != "alice" {
		t.Errorf("username = %v, want alice", got["username"])
	}

	nested := got["nested"].(map[string]any)
	if nested["old_password"] != "***" || nested["mobile"] != "13800000000" {
		t.Errorf("nested = %v", nested)
	}

	list := got["list"].([]any)
	if entry := list[0].(map[string]any); entry["token"] != "***" {
		t.Errorf("token = %v, want ***", entry["token"])
	}
	if entry := list[1].(map[string]any); entry["code"] != "***" {
		t.Errorf("code = %v, want ***", entry["code"])
	}

	if strings.Contains(string(raw), "plain-text") || strings.Contains(string(raw), "123456") {
		t.Fatalf("敏感值仍出现在日志中: %s", raw)
	}
}

func TestJSONKeepsInvalidInput(t *testing.T) {
	const raw = "not-json"
	if got := redactx.JSON(raw); got != raw {
		t.Fatalf("JSON(%q) = %q, want unchanged", raw, got)
	}
}

func TestJSONKeepsResponseCode(t *testing.T) {
	const raw = `{"code":"SUCCESS","password":"plain","data":{"code":"SUCCESS"}}`
	got := redactx.JSON(raw)

	if !strings.Contains(got, `"code":"SUCCESS"`) {
		t.Fatalf("response code removed: %s", got)
	}
	if strings.Contains(got, "plain") {
		t.Fatalf("password not redacted: %s", got)
	}
}
