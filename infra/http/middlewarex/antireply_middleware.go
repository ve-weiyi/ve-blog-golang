package middlewarex

import (
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
	"github.com/ve-weiyi/blog-cloud/infra/http/responsex"
)

// AntiReplyMiddleware 校验请求时间戳，限制重放窗口。
type AntiReplyMiddleware struct {
}

// NewAntiReplyMiddleware 创建防重放中间件。
func NewAntiReplyMiddleware() *AntiReplyMiddleware {
	return &AntiReplyMiddleware{}
}

// 防重放
func (m *AntiReplyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debugf("AntiReplyMiddleware Handle")

		ts := r.Header.Get(headerx.HeaderTimestamp)
		if ts == "" {
			// 读取类请求允许不带时间戳；写操作必须带
			switch r.Method {
			case http.MethodGet, http.MethodHead, http.MethodOptions:
				next.ServeHTTP(w, r)
			default:
				responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeParamMissing, "timestamp is required"))
			}
			return
		}

		value, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeRequestSignInvalid, "invalid timestamp"))
			return
		}

		// 只拦过去会让未来时间戳长期有效，因此两端都要限制
		delta := time.Now().Unix() - value
		if delta > 3600 || delta < -3600 {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizcode.CodeRequestSignInvalid, "timestamp is expired"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
