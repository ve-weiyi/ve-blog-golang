package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/http/cookiex"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

const testGuestID = "0123456789abcdef0123456789abcdef"

type fakeTokenService struct {
	claims *tokenx.Claims
	err    error
}

func (f *fakeTokenService) Issue(context.Context, tokenx.IssueRequest) (*tokenx.TokenPair, error) {
	return nil, nil
}

func (f *fakeTokenService) Validate(context.Context, string) (*tokenx.Claims, error) {
	return f.claims, f.err
}

func (f *fakeTokenService) Refresh(context.Context, tokenx.RefreshRequest) (*tokenx.TokenPair, error) {
	return nil, nil
}

func (f *fakeTokenService) Revoke(context.Context, string, string) error { return nil }

func (f *fakeTokenService) RevokeDevice(context.Context, string, string) error { return nil }

func (f *fakeTokenService) RevokeAll(context.Context, string) error { return nil }

func (f *fakeTokenService) ListSessions(context.Context, string) ([]tokenx.SessionInfo, error) {
	return nil, nil
}

func TestUserAuthMiddleware_RequiresDeviceBinding(t *testing.T) {
	cases := []struct {
		name         string
		headerDevice string
		cookieDevice string
		claimsDevice string
		wantNext     bool
	}{
		{"header matches", "d", "", "d", true},
		{"cookie matches", "", testGuestID, testGuestID, true},
		{"missing device", "", "", "d", false},
		{"wrong device", "x", "", "d", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := &fakeTokenService{claims: &tokenx.Claims{
				UserID:   "u",
				DeviceID: tc.claimsDevice,
			}}
			middleware := NewUserAuthMiddleware(svc)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/home", nil)
			req.Header.Set(headerx.HeaderAuthorization, "Bearer token")
			if tc.headerDevice != "" {
				req.Header.Set(headerx.HeaderXDeviceId, tc.headerDevice)
			}
			if tc.cookieDevice != "" {
				req.AddCookie(&http.Cookie{Name: cookiex.DeviceIDName, Value: tc.cookieDevice})
			}
			rec := httptest.NewRecorder()

			called := false
			middleware.Handle(func(http.ResponseWriter, *http.Request) {
				called = true
			})(rec, req)

			if called != tc.wantNext {
				t.Fatalf("next called = %v, want %v (status %d)", called, tc.wantNext, rec.Code)
			}
			if !tc.wantNext && rec.Code != http.StatusUnauthorized {
				t.Fatalf("status = %d, want 401", rec.Code)
			}
		})
	}
}
