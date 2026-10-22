package cookiex

import (
	"net/http"
	"strings"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/auth/identityx"
)

const (
	// DeviceIDName 是游客设备标识的第一方 Cookie 名。
	DeviceIDName = "blog_device_id"

	deviceIDMaxAge = 400 * 24 * time.Hour
)

// DeviceID 读取 Cookie 中的游客设备标识；不存在或形态不合法时返回空串。
func DeviceID(r *http.Request) string {
	cookie, err := r.Cookie(DeviceIDName)
	if err != nil {
		return ""
	}
	if identityx.KindOf(cookie.Value) != identityx.KindGuest {
		return ""
	}
	return cookie.Value
}

// SetDeviceID 写入游客设备标识。Secure 按请求是否来自 HTTPS 决定，兼容本地 HTTP 调试。
func SetDeviceID(w http.ResponseWriter, r *http.Request, id string) {
	if identityx.KindOf(id) != identityx.KindGuest {
		return
	}

	secure := r.TLS != nil || strings.EqualFold(r.Header.Get("X-Forwarded-Proto"), "https")
	http.SetCookie(w, &http.Cookie{
		Name:     DeviceIDName,
		Value:    id,
		Path:     "/",
		MaxAge:   int(deviceIDMaxAge.Seconds()),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
