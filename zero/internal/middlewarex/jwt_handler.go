package middlewarex

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/constantx"
)

// jwt handler
func JwtAuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var uid string
		var token string
		// 从真实http 头获取header
		if len(r.Header) > 0 {
			for k, v := range r.Header {
				if len(k) == 0 {
					continue
				}
				var value string
				if len(v) > 0 {
					value = v[0]
				}
				keyLowercase := strings.ToLower(k)

				if keyLowercase == constantx.HeaderXAuthToken {
					token = value
				}
				if keyLowercase == constantx.HeaderXUserID {
					uid = value
				}
			}
		}
		logx.Infof("JwtAuthHandler uid=%s, token=%s", uid, token)
		next.ServeHTTP(w, r)
	}
}
