package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type MenuRouter struct {
	svcCtx *svc.RouterContext
}

func NewMenuRouter(svcCtx *svc.RouterContext) *MenuRouter {
	return &MenuRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Menu 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *MenuRouter) InitMenuRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.MenuController
	{
		publicRouter.POST("menu/create", handler.CreateMenu)   // 新建Menu
		publicRouter.PUT("menu/update", handler.UpdateMenu)    // 更新Menu
		publicRouter.DELETE("menu/delete", handler.DeleteMenu) // 删除Menu
		publicRouter.POST("menu/find", handler.FindMenu)       // 查询Menu

		publicRouter.DELETE("menu/deleteByIds", handler.DeleteMenuByIds) // 批量删除Menu列表
		publicRouter.POST("menu/list", handler.FindMenuList)             // 分页查询Menu列表
	}
}
