package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type AIRouter struct {
	svcCtx *svc.RouterContext
}

func NewAIRouter(ctx *svc.RouterContext) *AIRouter {
	return &AIRouter{
		svcCtx: ctx,
	}
}

// 初始化 AI 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *AIRouter) InitAIRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AIController
	{
		loginRouter.POST("ai/assistant", handler.ChatAssistant)
		loginRouter.POST("ai/assistant/history", handler.ChatAssistantHistory)
		publicRouter.POST("ai/chat", handler.ChatAI)
		publicRouter.POST("ai/cos", handler.ChatCos)
	}
}
