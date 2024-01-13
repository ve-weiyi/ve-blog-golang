package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/apierror"
)

func PermissionHandler() gin.HandlerFunc {
	permissionHolder := global.Permission

	return func(c *gin.Context) {

		//获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		uid := c.GetString("uid")

		// 权限校验
		err := permissionHolder.CheckUserAccessApi(uid, obj, act)
		if err != nil {
			global.LOG.Error(err)
			c.JSON(http.StatusOK, response.Response{
				Code:    apierror.ErrorForbidden.Code(),
				Message: "用户无权限访问",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 检测接口是否可用
		permission, err := permissionHolder.FindApiPermission(obj, act)
		if err != nil {
			global.LOG.Error(err)
		}
		if permission.Status != 1 {
			c.JSON(http.StatusOK, response.Response{
				Code:    apierror.ErrorNotFound.Code(),
				Message: "该接口未开放",
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
