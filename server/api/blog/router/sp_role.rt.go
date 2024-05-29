package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type RoleRouter struct {
	svcCtx *svc.ServiceContext
}

func NewRoleRouter(svcCtx *svc.ServiceContext) *RoleRouter {
	return &RoleRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Role 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RoleRouter) InitRoleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewRoleController(s.svcCtx)
	{
		loginRouter.POST("/role/create_role", handler.CreateRole)   // 新建Role
		loginRouter.PUT("/role/update_role", handler.UpdateRole)    // 更新Role
		loginRouter.DELETE("/role/delete_role", handler.DeleteRole) // 删除Role
		loginRouter.POST("/role/find_role", handler.FindRole)       // 查询Role

		loginRouter.DELETE("/role/find_delete_role_list", handler.DeleteRoleList) // 批量删除Role列表
		loginRouter.POST("/role/find_role_list", handler.FindRoleList)            // 分页查询Role列表

		loginRouter.POST("/role/details_list", handler.FindRoleDetailsList)
		loginRouter.POST("/role/update_role_menus", handler.UpdateRoleMenus)
		loginRouter.POST("/role/update_role_resources", handler.UpdateRoleResources)
	}
}
