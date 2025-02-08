package middlewarex

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tokenx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
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

		token = r.Header.Get(restx.HeaderToken)
		uid = r.Header.Get(restx.HeaderUid)

		// 请求头缺少参数
		if token == "" || uid == "" {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserUnLogin, "用户未登录,缺少用户签名"))
			return
		}

		err := j.verifier.VerifyToken(r.Context(), token, uid)
		if err != nil {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserLoginExpired, err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
