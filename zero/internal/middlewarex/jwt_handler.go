package middlewarex

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
)

type (
	JwtTokenClaims struct {
		// 用户自定义字段
		jwt.RegisteredClaims
		jwt.MapClaims
	}

	JwtMiddleware struct {
		// 依赖注入
		Token *jtoken.JWTInstance
	}
)

func NewJwtMiddleware(tk *jtoken.JWTInstance) *JwtMiddleware {
	return &JwtMiddleware{
		Token: tk,
	}
}

// jwt handler
func (j *JwtMiddleware) JwtAuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string
		var uid string

		token = r.Header.Get(constant.HeaderToken)
		uid = r.Header.Get(constant.HeaderUid)

		//token为空或者uid为空
		if token == "" || uid == "" {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage("token or uid is empty"))
			return
		}

		// 解析token
		tok, err := j.Token.ParseToken(token)
		if err != nil {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage(err.Error()))
			return
		}

		// token不合法
		if !tok.Valid {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage("token is invalid"))
			return
		}

		// 获取claims
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage("token claims is not jwt.MapClaims"))
			return
		}

		// uid不一致
		if uid != cast.ToString(claims["uid"]) {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage("token uid is not equal"))
			return
		}

		//token验证成功,但用户在别处登录或退出登录
		//if jwtService.IsBlacklist(token) {
		//
		//}

		// 写入上下文
		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jtoken.JwtAudience, jtoken.JwtExpire, jtoken.JwtId, jtoken.JwtIssueAt, jtoken.JwtIssuer, jtoken.JwtNotBefore, jtoken.JwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}

		logx.Infof("JwtAuthHandler uid=%s, token=%s", uid, token)
		next.ServeHTTP(w, r)
	}
}
