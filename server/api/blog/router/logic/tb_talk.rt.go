package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type TalkRouter struct {
	svcCtx *svc.RouterContext
}

func NewTalkRouter(ctx *svc.RouterContext) *TalkRouter {
	return &TalkRouter{
		svcCtx: ctx,
	}
}

// 初始化 Talk 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TalkRouter) InitTalkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.TalkController
	{
		publicRouter.POST("talk/create", handler.CreateTalk)   // 新建Talk
		publicRouter.PUT("talk/update", handler.UpdateTalk)    // 更新Talk
		publicRouter.DELETE("talk/delete", handler.DeleteTalk) // 删除Talk
		publicRouter.POST("talk/query", handler.GetTalk)       // 查询Talk

		publicRouter.DELETE("talk/deleteByIds", handler.DeleteTalkByIds) // 批量删除Talk列表
		publicRouter.POST("talk/list", handler.FindTalkList)             // 分页查询Talk列表
	}
}
