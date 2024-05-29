package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

// JwtToken jwt中间件
func JwtToken(svcCtx *svc.ServiceContext) gin.HandlerFunc {

	tk := svcCtx.Token

	return func(c *gin.Context) {
		token := c.Request.Header.Get(constant.HeaderXAuthToken)
		uid := c.Request.Header.Get(constant.HeaderXUserId)

		//token为空或者uid为空
		if token == "" || uid == "" {
			// 有错误，直接返回给前端错误，前端直接报错500
			//c.AbortWithStatus(http.StatusInternalServerError)
			// 该方式前端不报错
			c.JSON(http.StatusOK, apierr.ErrorUserUnLogin)
			c.Abort()
			return
		}

		// 解析token
		claims, err := tk.ParserToken(token)

		// token验证失败
		if err != nil {
			c.JSON(http.StatusOK, apierr.ErrorUserUnLogin.WrapError(err))
			c.Abort()
			return
		}

		// uid不一致
		if uid != strconv.Itoa(claims.Ext.Uid) {
			c.JSON(http.StatusOK, apierr.ErrorUserUnLogin.WrapMessage("uid is not equal"))
			c.Abort()
			return
		}

		//token验证成功,但用户在别处登录或退出登录
		//if jwtService.IsBlacklist(token) {
		//
		//}

		glog.Infof("user login-->%v", claims.Ext.Username)
		//glog.JsonIndent(claims)

		c.Set("token", token)
		c.Set("uid", uid)
		c.Set("domain", claims.Issuer)
		c.Set("username", claims.Ext.Username)
		c.Set("login_type", claims.Ext.LoginType)
		c.Next()

	}
}
