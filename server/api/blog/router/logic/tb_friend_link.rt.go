package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type FriendLinkRouter struct {
	svcCtx *svc.RouterContext
}

func NewFriendLinkRouter(svcCtx *svc.RouterContext) *FriendLinkRouter {
	return &FriendLinkRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 FriendLink 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *FriendLinkRouter) InitFriendLinkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.FriendLinkController
	{
		publicRouter.POST("friendLink/create", handler.CreateFriendLink)   // 新建FriendLink
		publicRouter.PUT("friendLink/update", handler.UpdateFriendLink)    // 更新FriendLink
		publicRouter.DELETE("friendLink/delete", handler.DeleteFriendLink) // 删除FriendLink
		publicRouter.POST("friendLink/find", handler.FindFriendLink)       // 查询FriendLink

		publicRouter.DELETE("friendLink/deleteByIds", handler.DeleteFriendLinkByIds) // 批量删除FriendLink列表
		publicRouter.POST("friendLink/list", handler.FindFriendLinkList)             // 分页查询FriendLink列表
	}
}
