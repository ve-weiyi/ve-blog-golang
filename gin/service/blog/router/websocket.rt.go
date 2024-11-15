package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type WebsocketRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketRouter(svcCtx *svctx.ServiceContext) *WebsocketRouter {
	return &WebsocketRouter{
		svcCtx: svcCtx,
	}
}

func (s *WebsocketRouter) Register(r *gin.RouterGroup) {
	// Websocket
	// []
	{
		group := r.Group("/api/v1")

		handler := controller.NewWebsocketController(s.svcCtx)
		// WebSocket消息
		group.GET("/websocket", handler.WebSocket)
	}
}
