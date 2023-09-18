package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
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
func (s *RoleRouter) InitRoleBasicRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.RoleController
	{
		publicRouter.POST("role", handler.CreateRole)       // 新建Role
		publicRouter.PUT("role", handler.UpdateRole)        // 更新Role
		publicRouter.DELETE("role/:id", handler.DeleteRole) // 删除Role
		publicRouter.GET("role/:id", handler.FindRole)      // 查询Role

		publicRouter.DELETE("role/batch_delete", handler.DeleteRoleByIds) // 批量删除Role列表
		publicRouter.POST("role/list", handler.FindRoleList)              // 分页查询Role列表
	}
}
