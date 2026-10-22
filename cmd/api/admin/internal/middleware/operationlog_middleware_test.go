package middleware

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-openapi/loads"
)

func TestCappedBuffer(t *testing.T) {
	b := &cappedBuffer{limit: 5}
	if n, err := b.Write([]byte("123456789")); err != nil || n != 9 {
		t.Fatalf("Write() = (%d, %v), want (9, nil)", n, err)
	}
	if got := b.String(); got != "12345" {
		t.Errorf("String() = %q, want %q", got, "12345")
	}
}

func TestGetAPISpecMatchesParameterizedPath(t *testing.T) {
	doc, err := loads.Analyzed(json.RawMessage(`{
		"swagger": "2.0",
		"info": {"title": "test", "version": "1.0"},
		"paths": {
			"/admin-api/v1/articles/{id}": {
				"get": {"summary": "get article", "tags": ["文章"]}
			}
		}
	}`), "")
	if err != nil {
		t.Fatalf("load spec: %v", err)
	}

	m := NewOperationLogMiddleware(doc.Spec(), nil, nil)
	op := m.getApiSpec("/admin-api/v1/articles/42", http.MethodGet)
	if op == nil {
		t.Fatal("getApiSpec() = nil, want operation")
	}
	if op.Summary != "get article" {
		t.Errorf("Summary = %q, want %q", op.Summary, "get article")
	}
}

func TestResponseRecorderFlushAndHijack(t *testing.T) {
	writer := &flushWriter{}
	rec := &responseRecorder{
		ResponseWriter: writer,
		body:           &cappedBuffer{limit: 10},
	}

	rec.Flush()
	if !writer.flushed {
		t.Error("Flush() did not reach the underlying writer")
	}

	if _, _, err := rec.Hijack(); err != http.ErrNotSupported {
		t.Errorf("Hijack() err = %v, want %v", err, http.ErrNotSupported)
	}
}

type flushWriter struct {
	flushed bool
}

func (w *flushWriter) Header() http.Header         { return http.Header{} }
func (w *flushWriter) Write(p []byte) (int, error) { return len(p), nil }
func (w *flushWriter) WriteHeader(int)             {}
func (w *flushWriter) Flush()                      { w.flushed = true }
