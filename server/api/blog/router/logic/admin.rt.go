package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type AdminRouter struct {
	svcCtx *svc.RouterContext
}

func NewAdminRouter(ctx *svc.RouterContext) *AdminRouter {
	return &AdminRouter{
		svcCtx: ctx,
	}
}

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *AdminRouter) InitAdminRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.AdminController

	{
		loginRouter.POST("admin/roles", handler.GetRoleTreeList) // 获取Role列表
		loginRouter.POST("admin/menus", handler.GetMenuTreeList) // 获取Menu列表
		loginRouter.POST("admin/apis", handler.GetApiTreeList)   // 获取Api列表

		loginRouter.POST("admin/role/update_menus", handler.UpdateRoleMenus)         // 获取Role列表
		loginRouter.POST("admin/role/update_resources", handler.UpdateRoleResources) // 获取Role列表
	}
	{
		loginRouter.GET("admin/home", handler.GetHomeInfo)           // 获取首页信息
		loginRouter.POST("admin/users", handler.GetUserList)         // 获取用户列表
		loginRouter.POST("admin/user/areas", handler.GetUserAreas)   // 获取用户地区
		loginRouter.POST("admin/comments", handler.GetAdminComments) // 获取评论列表
		loginRouter.POST("admin/about", handler.UpdateAbout)         // 获取首页信息
	}
}
