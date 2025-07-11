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
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewMenuController(s.svcCtx)
		// 创建菜单
		group.POST("/menu/add_menu", handler.AddMenu)
		// 清空菜单列表
		group.POST("/menu/clean_menu_list", handler.CleanMenuList)
		// 删除菜单
		group.DELETE("/menu/deletes_menu", handler.DeletesMenu)
		// 分页获取菜单列表
		group.POST("/menu/find_menu_list", handler.FindMenuList)
		// 同步菜单列表
		group.POST("/menu/sync_menu_list", handler.SyncMenuList)
		// 更新菜单
		group.PUT("/menu/update_menu", handler.UpdateMenu)
	}
}
