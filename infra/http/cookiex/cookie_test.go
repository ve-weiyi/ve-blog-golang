package cookiex

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeviceID(t *testing.T) {
	tests := []struct {
		name   string
		cookie *http.Cookie
		want   string
	}{
		{"valid", &http.Cookie{Name: DeviceIDName, Value: "d41d8cd98f00b204e9800998ecf8427e"}, "d41d8cd98f00b204e9800998ecf8427e"},
		{"uppercase", &http.Cookie{Name: DeviceIDName, Value: "D41D8CD98F00B204E9800998ECF8427E"}, ""},
		{"wrong length", &http.Cookie{Name: DeviceIDName, Value: "abc"}, ""},
		{"not hex", &http.Cookie{Name: DeviceIDName, Value: "z41d8cd98f00b204e9800998ecf8427e"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.AddCookie(tt.cookie)
			if got := DeviceID(req); got != tt.want {
				t.Fatalf("DeviceID() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSetDeviceID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	SetDeviceID(rec, req, "d41d8cd98f00b204e9800998ecf8427e")

	resp := rec.Result()
	defer resp.Body.Close()
	cookies := resp.Cookies()
	if len(cookies) != 1 {
		t.Fatalf("cookies = %d, want 1", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != DeviceIDName || cookie.Value != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Fatalf("cookie = %+v", cookie)
	}
	if !cookie.HttpOnly {
		t.Fatal("cookie should be HttpOnly")
	}
	if cookie.SameSite != http.SameSiteLaxMode {
		t.Fatalf("SameSite = %v, want Lax", cookie.SameSite)
	}
	if cookie.Path != "/" {
		t.Fatalf("Path = %q, want /", cookie.Path)
	}
}

func TestSetDeviceIDRejectsInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	SetDeviceID(rec, req, "not-a-device-id")

	if cookies := rec.Result().Cookies(); len(cookies) != 0 {
		t.Fatalf("cookies = %v, want none", cookies)
	}
}
