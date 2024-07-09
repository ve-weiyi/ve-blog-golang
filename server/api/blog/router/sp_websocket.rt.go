package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type WebsocketRouter struct {
	svcCtx *svc.ServiceContext
}

func NewWebsocketRouter(svcCtx *svc.ServiceContext) *WebsocketRouter {
	return &WebsocketRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Blog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *WebsocketRouter) InitWebsocketRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewWebsocketController(s.svcCtx)
	{
		publicRouter.GET("/ws", handler.WebSocket) // websocket
	}
}
