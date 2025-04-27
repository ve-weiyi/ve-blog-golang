package middlewarex

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/responsex"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/tokenx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
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
		if token == "" {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "request header field 'token' is missing"))
			return
		}

		if uid == "" {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "request header field 'uid' is missing"))
			return
		}

		err := j.verifier.VerifyToken(r.Context(), token, uid)
		if err != nil {
			responsex.Response(r, w, nil, bizerr.NewBizError(bizerr.CodeUserLoginExpired, err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
