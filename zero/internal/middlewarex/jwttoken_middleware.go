package middlewarex

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
)

type JwtTokenMiddleware struct {
	Token *jtoken.JwtInstance
	Redis *redis.Redis
}

func NewJwtTokenMiddleware(tk *jtoken.JwtInstance, rds *redis.Redis) *JwtTokenMiddleware {
	return &JwtTokenMiddleware{
		Token: tk,
		Redis: rds,
	}
}

func (j *JwtTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("JwtTokenMiddleware Handle")
		var token string
		var uid string

		token = r.Header.Get(constant.HeaderAuthorization)
		uid = r.Header.Get(constant.HeaderUid)

		//token为空或者uid为空
		if token == "" || uid == "" {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, "token or uid is empty"))
			return
		}

		// 解析token
		tok, err := j.Token.ParseToken(token)
		if err != nil {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, err.Error()))
			return
		}

		// token不合法
		if !tok.Valid {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, "token is invalid"))
			return
		}

		// 获取claims
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, "token claims is not jwt.MapClaims"))
			return
		}

		// uid不一致
		if uid != cast.ToString(claims["uid"]) {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, "token cannot use by uid"))
			return
		}

		//token验证成功,但用户在别处登录或退出登录
		if j.IsBlacklist(claims) {
			responsex.Response(r, w, nil, apierr.NewApiError(codex.CodeUserNotPermission, "user already logout or login in other place"))
			return
		}

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

		logx.Infof("JwtMiddleware uid=%s, token=%s", uid, token)
		next.ServeHTTP(w, r)
	}
}

// 已退出登录
func (j *JwtTokenMiddleware) IsBlacklist(claims jwt.MapClaims) bool {
	uid := cast.ToInt64(claims["uid"])
	loginAt := cast.ToInt64(claims[jtoken.JwtIssueAt])

	redisKey := rediskey.GetUserLogoutKey(cast.ToString(uid))

	at, err := j.Redis.Get(redisKey)
	if err != nil {
		return false
	}

	logoutAt := cast.ToInt64(at)

	logx.Infof("loginAt=%d, at.LogoutAt=%d", loginAt, logoutAt)

	if loginAt < logoutAt {
		return true
	}

	return false
}
