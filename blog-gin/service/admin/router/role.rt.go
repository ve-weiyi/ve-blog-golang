package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type RoleRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewRoleRouter(svcCtx *svctx.ServiceContext) *RoleRouter {
	return &RoleRouter{
		svcCtx: svcCtx,
	}
}

func (s *RoleRouter) Register(r *gin.RouterGroup) {
	// Role
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewRoleController(s.svcCtx)
		// 创建角色
		group.POST("/role/add_role", handler.AddRole)
		// 删除角色
		group.POST("/role/deletes_role", handler.DeletesRole)
		// 分页获取角色列表
		group.POST("/role/find_role_list", handler.FindRoleList)
		// 获取角色资源列表
		group.POST("/role/find_role_resources", handler.FindRoleResources)
		// 更新角色
		group.PUT("/role/update_role", handler.UpdateRole)
		// 更新角色接口权限
		group.POST("/role/update_role_apis", handler.UpdateRoleApis)
		// 更新角色菜单权限
		group.POST("/role/update_role_menus", handler.UpdateRoleMenus)
	}
}
