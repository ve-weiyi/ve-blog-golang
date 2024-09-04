package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr/codex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

func PermissionHandler(svcCtx *svctx.ServiceContext) gin.HandlerFunc {
	permissionHolder := svcCtx.RbacHolder

	return func(c *gin.Context) {

		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		uid := c.GetString("uid")

		// 权限校验
		err := permissionHolder.CheckUserAccessApi(uid, obj, act)
		if err != nil {
			glog.Error(err)
			c.JSON(http.StatusOK, apierr.NewApiError(codex.CodeUserNotPermission, "无操作权限"))
			c.Abort()
			return
		}

		// 检测接口是否可用
		permission, err := permissionHolder.FindApiPermission(obj, act)
		if err != nil {
			glog.Error(err)
		}
		if permission != nil && permission.Status != 1 {
			c.JSON(http.StatusOK, apierr.NewApiError(codex.CodeUserNotPermission, "接口已禁用"))
			c.Abort()
			return
		}
		// 挂起当前中间件，执行下一个中间件
		c.Next()
	}
}
