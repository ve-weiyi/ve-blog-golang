package middlewarex

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
)

type TimeTokenMiddleware struct {
}

func NewTimeTokenMiddleware() *TimeTokenMiddleware {
	return &TimeTokenMiddleware{}
}

// 游客token
// token = md5(tm,ts)
func (m *TimeTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("TimeTokenMiddleware Handle")
		tk := r.Header.Get(restx.HeaderXAuthToken)
		ts := r.Header.Get(restx.HeaderTimestamp)
		tm := r.Header.Get(restx.HeaderTerminal)

		// 请求头缺少参数
		if tk == "" || tm == "" || ts == "" {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求,缺少游客签名"))
			return
		}
		// 判断 token = md5(tm,ts)
		if tk != crypto.Md5v(tm, ts) {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求,游客签名错误"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
