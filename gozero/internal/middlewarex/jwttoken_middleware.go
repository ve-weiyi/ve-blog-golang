package middlewarex

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr/codex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jwtx"
)

type TokenVerifier interface {
	VerifyToken(ctx context.Context, token string, uid string) (jwt.MapClaims, error)
}

type JwtTokenMiddleware struct {
	verifier TokenVerifier
}

func NewJwtTokenMiddleware(verifier TokenVerifier) *JwtTokenMiddleware {
	return &JwtTokenMiddleware{
		verifier: verifier,
	}
}

func (j *JwtTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("JwtTokenMiddleware Handle")
		var token string
		var uid string

		token = r.Header.Get(constant.HeaderAuthorization)
		uid = r.Header.Get(constant.HeaderUid)

		claims, err := j.verifier.VerifyToken(r.Context(), token, uid)
		if err != nil {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, err.Error()))
			return
		}

		// 写入上下文
		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jwtx.JwtAudience, jwtx.JwtExpire, jwtx.JwtId, jwtx.JwtIssueAt, jwtx.JwtIssuer, jwtx.JwtNotBefore, jwtx.JwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
