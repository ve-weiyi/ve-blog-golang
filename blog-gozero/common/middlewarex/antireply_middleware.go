package middlewarex

import (
	"net/http"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
)

type AntiReplyMiddleware struct {
}

func NewAntiReplyMiddleware() *AntiReplyMiddleware {
	return &AntiReplyMiddleware{}
}

// 防重放
func (m *AntiReplyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debugf("AntiReplyMiddleware Handle")

		if r.Method != http.MethodGet {
			ts := r.Header.Get(restx.HeaderTimestamp)
			if ts == "" {
				responsex.Response(r, w, nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "request header field 'timestamp' is missing"))
				return
			}

			now := time.Now().Unix()
			if now-cast.ToInt64(ts) > 3600 {
				responsex.Response(r, w, nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "timestamp is expired"))
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}
