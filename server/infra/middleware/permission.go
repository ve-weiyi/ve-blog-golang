package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
)

func PermissionHandler() gin.HandlerFunc {
	permission := global.Permission

	return func(c *gin.Context) {

		//获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		uid := c.GetInt("uid")
		// 判断用户是否有权限
		ok, err := permission.VerifyUserPermissions(uid, obj, act)
		if err != nil {
			return
		}
		// 没有权限
		if !ok {
			c.JSON(http.StatusForbidden,
				response.Response{
					Code:    codes.CodeRoleNoPerPermission,
					Message: "角色权限不足",
					Data:    nil,
				})
			c.Abort()
			return
		}

		c.Next()
	}
}
