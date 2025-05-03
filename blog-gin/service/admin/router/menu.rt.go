package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type MenuRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewMenuRouter(svcCtx *svctx.ServiceContext) *MenuRouter {
	return &MenuRouter{
		svcCtx: svcCtx,
	}
}

func (s *MenuRouter) Register(r *gin.RouterGroup) {
	// Menu
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewMenuController(s.svcCtx)
		// 创建菜单
		group.POST("/menu/add_menu", handler.AddMenu)
		// 批量删除菜单
		group.DELETE("/menu/batch_delete_menu", handler.BatchDeleteMenu)
		// 清空菜单列表
		group.POST("/menu/clean_menu_list", handler.CleanMenuList)
		// 删除菜单
		group.DELETE("/menu/delete_menu", handler.DeleteMenu)
		// 分页获取菜单列表
		group.POST("/menu/find_menu_list", handler.FindMenuList)
		// 同步菜单列表
		group.POST("/menu/sync_menu_list", handler.SyncMenuList)
		// 更新菜单
		group.PUT("/menu/update_menu", handler.UpdateMenu)
	}
}
