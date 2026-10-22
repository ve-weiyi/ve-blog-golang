package swaggerx

import (
	"encoding/json"
	"testing"

	"github.com/go-openapi/spec"
)

func TestOperations(t *testing.T) {
	const raw = `{
	  "swagger": "2.0",
	  "info": {"title": "test", "version": "1.0"},
	  "paths": {
	    "/comments": {
	      "get": {"summary": "查询评论", "tags": ["评论管理"]}
	    },
	    "/comments/{id}": {
	      "delete": {"summary": "删除评论", "tags": ["评论管理"]}
	    },
	    "/comments/batch-review": {
	      "patch": {"summary": "批量审核评论", "tags": ["评论管理"]}
	    },
	    "/comments/{id}/review": {
	      "patch": {"summary": "审核评论", "tags": ["评论管理"]}
	    }
	  }
	}`

	var sp spec.Swagger
	if err := json.Unmarshal([]byte(raw), &sp); err != nil {
		t.Fatalf("unmarshal swagger: %v", err)
	}

	got := Operations(&sp)

	want := []Operation{
		{Path: "/comments", Method: MethodGet, Name: "查询评论", Tags: []string{"评论管理"}},
		{Path: "/comments/batch-review", Method: MethodPatch, Name: "批量审核评论", Tags: []string{"评论管理"}},
		{Path: "/comments/{id}", Method: MethodDelete, Name: "删除评论", Tags: []string{"评论管理"}},
		{Path: "/comments/{id}/review", Method: MethodPatch, Name: "审核评论", Tags: []string{"评论管理"}},
	}

	if len(got) != len(want) {
		t.Fatalf("operation count = %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i].Path != w.Path || got[i].Method != w.Method || got[i].Name != w.Name {
			t.Errorf("operations[%d] = {%q, %q, %q}, want {%q, %q, %q}",
				i, got[i].Path, got[i].Method, got[i].Name, w.Path, w.Method, w.Name)
		}
		if len(got[i].Tags) != 1 || got[i].Tags[0] != w.Tags[0] {
			t.Errorf("operations[%d].Tags = %v, want %v", i, got[i].Tags, w.Tags)
		}
	}
}
