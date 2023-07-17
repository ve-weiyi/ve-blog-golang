package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
)

type UserRouter struct {
	svcCtx *svc.RouterContext
}

func NewUserRouter(ctx *svc.RouterContext) *UserRouter {
	return &UserRouter{
		svcCtx: ctx,
	}
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	// 需要用户角色验证
	roleRouter := publicRouter.Group("")
	roleRouter.Use(middleware.CasbinHandler())

	// 留下操作记录
	traceRouter := loginRouter.Group("")
	traceRouter.Use(middleware.OperationRecord())

	var self = s.svcCtx.AppController.UserController
	{
		loginRouter.GET("user/info", self.GetUserinfo)              // 用户信息
		loginRouter.GET("user/login_history", self.GetLoginHistory) // 用户信息
		loginRouter.GET("user/menus", self.GetUserMenus)            // 用户菜单
		loginRouter.GET("user/resources", self.GetUserApis)         // 用户资源
		loginRouter.POST("user/avatar", self.UpdateUserAvatar)      // 更新用户头像
	}
}
