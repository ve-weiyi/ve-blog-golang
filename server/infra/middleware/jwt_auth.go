package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
)

// JwtToken jwt中间件
func JwtToken(isNeed bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := global.JWT.ParseTokenByGin(c)

		//必须要token才能过
		if err != nil {
			if !isNeed {
				c.Next()
				return
			}
			global.LOG.Error("token-->", c.Request.Header.Get("Authorization"))
			global.LOG.Error("-->", err)
			// 有错误，直接返回给前端错误，前端直接报错500
			//c.AbortWithStatus(http.StatusInternalServerError)
			// 该方式前端不报错
			c.JSON(http.StatusOK, response.Response{
				Code: codes.CodeRoleNoPerPermission,
				Msg:  err.Error(),
				Data: nil,
			})
			c.Abort()
			return
		}

		global.LOG.JsonIndent(claims)
		c.Set("login_type", constant.LoginJwtToken)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Set("domain", claims.Issuer)
		c.Next()

	}
}
