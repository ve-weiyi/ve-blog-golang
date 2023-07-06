package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
)

type RoleRouter struct {
	svcCtx *svc.RouterContext
}

func NewRoleRouter(svcCtx *svc.RouterContext) *RoleRouter {
	return &RoleRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Role 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RoleRouter) InitRoleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	roleTraceRouter := loginRouter.Group("admin")
	roleTraceRouter.Use(middleware.OperationRecord())

	var handler = s.svcCtx.AppController.RoleController
	{
		publicRouter.POST("role/create", handler.CreateRole)   // 新建Role
		publicRouter.PUT("role/update", handler.UpdateRole)    // 更新Role
		publicRouter.DELETE("role/delete", handler.DeleteRole) // 删除Role
		publicRouter.POST("role/query", handler.GetRole)       // 查询Role

		publicRouter.DELETE("role/deleteByIds", handler.DeleteRoleByIds) // 批量删除Role列表
		publicRouter.POST("role/list", handler.FindRoleList)             // 分页查询Role列表
	}
	{
		loginRouter.GET("roles", handler.GetRoleTreeList) // 获取Role列表

		roleTraceRouter.POST("role/update_menus", handler.UpdateRoleMenus)         // 获取Role列表
		roleTraceRouter.POST("role/update_resources", handler.UpdateRoleResources) // 获取Role列表
	}
}
