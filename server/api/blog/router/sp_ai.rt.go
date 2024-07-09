package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type AIRouter struct {
	svcCtx *svc.ServiceContext
}

func NewAIRouter(svcCtx *svc.ServiceContext) *AIRouter {
	return &AIRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 AI 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *AIRouter) InitAIRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewAIController(s.svcCtx)
	{
		publicRouter.POST("/ai/chat", handler.ChatAI)
		publicRouter.POST("/ai/cos", handler.ChatCos)
		publicRouter.GET("/ai/chat/stream", handler.ChatStream)
		publicRouter.POST("/ai/assistant/history", handler.ChatAssistantHistory)
	}
}
