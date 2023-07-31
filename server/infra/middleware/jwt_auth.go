package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		uid := c.Request.Header.Get("uid")
		claims, err := global.JWT.VerifyToken(token, uid)
		//必须要token才能过
		if err != nil {
			global.LOG.Error("token-->", token)
			global.LOG.Error("-->", err)
			// 有错误，直接返回给前端错误，前端直接报错500
			//c.AbortWithStatus(http.StatusInternalServerError)
			// 该方式前端不报错
			c.JSON(http.StatusOK, response.Response{
				Code:    codes.CodeRoleNoPerPermission,
				Message: err.Error(),
				Data:    nil,
			})
			c.Abort()
			return
		}

		global.LOG.JsonIndent(claims)

		c.Set("login_type", claims.LoginType)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Set("domain", claims.Issuer)
		c.Next()

	}
}
