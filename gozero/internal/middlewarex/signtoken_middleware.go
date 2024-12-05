package middlewarex

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tokenx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/headerconst"
)

type SignTokenMiddleware struct {
	verifier tokenx.TokenHolder
}

func NewSignTokenMiddleware(verifier tokenx.TokenHolder) *SignTokenMiddleware {
	return &SignTokenMiddleware{
		verifier: verifier,
	}
}

// 用户token
// 缓存验证
func (j *SignTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("SignTokenMiddleware Handle")
		var token string
		var uid string

		token = r.Header.Get(headerconst.HeaderToken)
		uid = r.Header.Get(headerconst.HeaderUid)

		// 请求头缺少参数
		if token == "" || uid == "" {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求,缺少用户签名"))
			return
		}

		err := j.verifier.VerifyToken(r.Context(), token, uid)
		if err != nil {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
