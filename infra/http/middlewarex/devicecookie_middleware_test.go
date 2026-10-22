package middlewarex

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/blog-cloud/infra/http/cookiex"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

const testDeviceID = "d41d8cd98f00b204e9800998ecf8427e"

func TestDeviceCookieMiddleware_CookieWins(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/home", nil)
	req.AddCookie(&http.Cookie{Name: cookiex.DeviceIDName, Value: testDeviceID})
	req.Header.Set(headerx.HeaderXDeviceId, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	req = req.WithContext(metadata.NewOutgoingContext(req.Context(), metadata.MD{}))
	rec := httptest.NewRecorder()

	var got string
	handler := NewDeviceCookieMiddleware().Handle(func(_ http.ResponseWriter, r *http.Request) {
		md, _ := metadata.FromOutgoingContext(r.Context())
		if values := md.Get(headerx.HeaderXDeviceId); len(values) > 0 {
			got = values[0]
		}
	})
	handler(rec, req)

	if got != testDeviceID {
		t.Fatalf("metadata device id = %q, want cookie value %q", got, testDeviceID)
	}
	if cookies := rec.Result().Cookies(); len(cookies) != 0 {
		t.Fatalf("Set-Cookie = %v, want none", cookies)
	}
}

func TestDeviceCookieMiddleware_HeaderOnlyPassesThrough(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/home", nil)
	req.Header.Set(headerx.HeaderXDeviceId, testDeviceID)
	req = req.WithContext(metadata.NewOutgoingContext(req.Context(), metadata.Pairs(
		headerx.HeaderXDeviceId, testDeviceID,
	)))
	rec := httptest.NewRecorder()

	var got string
	handler := NewDeviceCookieMiddleware().Handle(func(_ http.ResponseWriter, r *http.Request) {
		md, _ := metadata.FromOutgoingContext(r.Context())
		if values := md.Get(headerx.HeaderXDeviceId); len(values) > 0 {
			got = values[0]
		}
	})
	handler(rec, req)

	if got != testDeviceID {
		t.Fatalf("metadata device id = %q, want header value %q", got, testDeviceID)
	}
	if cookies := rec.Result().Cookies(); len(cookies) != 0 {
		t.Fatalf("Set-Cookie = %v, want none", cookies)
	}
}

func TestDeviceCookieMiddleware_NoIdentity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/home", nil)
	req = req.WithContext(metadata.NewOutgoingContext(req.Context(), metadata.MD{}))
	rec := httptest.NewRecorder()

	called := false
	handler := NewDeviceCookieMiddleware().Handle(func(_ http.ResponseWriter, _ *http.Request) {
		called = true
	})
	handler(rec, req)

	if !called {
		t.Fatal("next handler not called")
	}
	if cookies := rec.Result().Cookies(); len(cookies) != 0 {
		t.Fatalf("Set-Cookie = %v, want none", cookies)
	}
}
