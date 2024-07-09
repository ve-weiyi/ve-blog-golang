package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type MenuRouter struct {
	svcCtx *svc.ServiceContext
}

func NewMenuRouter(svcCtx *svc.ServiceContext) *MenuRouter {
	return &MenuRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Menu 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *MenuRouter) InitMenuRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewMenuController(s.svcCtx)
	{
		loginRouter.POST("/menu/create_menu", handler.CreateMenu)            // 新建Menu
		loginRouter.PUT("/menu/update_menu", handler.UpdateMenu)             // 更新Menu
		loginRouter.DELETE("/menu/delete_menu", handler.DeleteMenu)          // 删除Menu
		loginRouter.DELETE("/menu/delete_menu_list", handler.DeleteMenuList) // 批量删除Menu列表

		loginRouter.POST("/menu/find_menu", handler.FindMenu)          // 查询Menu
		loginRouter.POST("/menu/find_menu_list", handler.FindMenuList) // 分页查询Menu列表

		loginRouter.POST("/menu/find_menu_details_list", handler.FindMenuDetailsList)

		loginRouter.POST("/menu/sync_menu_list", handler.SyncMenuList)   // 同步Menu列表
		loginRouter.POST("/menu/clean_menu_list", handler.CleanMenuList) // 清空Menu列表
	}
}
