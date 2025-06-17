package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type LoginLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewLoginLogRouter(svcCtx *svctx.ServiceContext) *LoginLogRouter {
	return &LoginLogRouter{
		svcCtx: svcCtx,
	}
}

func (s *LoginLogRouter) Register(r *gin.RouterGroup) {
	// LoginLog
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewLoginLogController(s.svcCtx)
		// 删除登录日志
		group.DELETE("/login_log/deletes_login_log", handler.DeletesLoginLog)
		// 查询登录日志
		group.POST("/user/find_login_log_list", handler.FindLoginLogList)
	}
}
