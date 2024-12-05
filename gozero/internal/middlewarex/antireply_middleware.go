package middlewarex

import (
	"net/http"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/headerconst"
)

type AntiReplyMiddleware struct {
}

func NewAntiReplyMiddleware() *AntiReplyMiddleware {
	return &AntiReplyMiddleware{}
}

// 防重放
func (m *AntiReplyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("AntiReplyMiddleware Handle")

		if r.Method != http.MethodGet {
			ts := r.Header.Get(headerconst.HeaderTimestamp)
			if ts == "" {
				responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "timestamp is empty"))
				return
			}

			now := time.Now().Unix()
			if now-cast.ToInt64(ts) > 3600 {
				responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "timestamp is invalid"))
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}
