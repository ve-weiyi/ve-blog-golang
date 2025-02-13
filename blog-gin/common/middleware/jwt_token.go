package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jwtx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
)

// JwtToken jwt中间件
func JwtToken(tk *jwtx.JwtInstance) gin.HandlerFunc {

	parser := tk

	return func(c *gin.Context) {
		var token string
		var uid string

		token = c.Request.Header.Get(restx.HeaderToken)
		uid = c.Request.Header.Get(restx.HeaderUid)

		// token为空或者uid为空
		if token == "" || uid == "" {
			// 有错误，直接返回给前端错误，前端直接报错500
			// c.AbortWithStatus(http.StatusInternalServerError)
			// 该方式前端不报错
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserUnLogin, "用户未登录"))
			c.Abort()
			return
		}

		// 解析token
		tok, err := parser.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserLoginExpired, err.Error()))
			c.Abort()
			return
		}

		// token不合法
		if !tok.Valid {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserLoginExpired, "token is invalid"))
			c.Abort()
			return
		}

		// 获取claims
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserLoginExpired, "token claims is not jwt.MapClaims"))
			c.Abort()
			return
		}

		// uid不一致
		if uid != cast.ToString(claims["uid"]) {
			c.JSON(http.StatusOK, apierr.NewApiError(apierr.CodeUserLoginExpired, "token uid is not equal"))
			c.Abort()
			return
		}

		// token验证成功,但用户在别处登录或退出登录
		// if jwtService.IsBlacklist(token) {
		//
		// }

		glog.Infof("user login-->%v", claims)

		// 写入上下文
		ctx := c.Request.Context()
		for k, v := range claims {
			switch k {
			case jwtx.JwtAudience, jwtx.JwtExpire, jwtx.JwtId, jwtx.JwtIssueAt, jwtx.JwtIssuer, jwtx.JwtNotBefore, jwtx.JwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}

		c.Next()

	}
}
