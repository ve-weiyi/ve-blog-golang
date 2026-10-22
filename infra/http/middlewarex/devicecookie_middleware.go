package middlewarex

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/blog-cloud/infra/http/cookiex"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

// DeviceCookieMiddleware 让第一方 Cookie 成为设备标识的权威值。
// Cookie 存在时覆盖请求头；Cookie 缺失时保留请求头，由 /guest 在签发时读取并下发 Cookie。
type DeviceCookieMiddleware struct {
}

func NewDeviceCookieMiddleware() *DeviceCookieMiddleware {
	return &DeviceCookieMiddleware{}
}

func (m *DeviceCookieMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := cookiex.DeviceID(r)
		if id == "" {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), headerx.HeaderXDeviceId, id)
		if md, ok := metadata.FromOutgoingContext(ctx); ok {
			md.Set(headerx.HeaderXDeviceId, id)
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
