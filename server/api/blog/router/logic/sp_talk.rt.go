package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type TalkRouter struct {
	svcCtx *svc.RouterContext
}

func NewTalkRouter(svcCtx *svc.RouterContext) *TalkRouter {
	return &TalkRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Talk 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TalkRouter) InitTalkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.TalkController
	{
		publicRouter.POST("/talk/create_talk", handler.CreateTalk)   // 新建Talk
		publicRouter.PUT("/talk/update_talk", handler.UpdateTalk)    // 更新Talk
		publicRouter.DELETE("/talk/delete_talk", handler.DeleteTalk) // 删除Talk
		publicRouter.POST("/talk/find_talk", handler.FindTalk)       // 查询Talk

		publicRouter.DELETE("/talk/delete_talk_list", handler.DeleteTalkList) // 批量删除Talk列表
		publicRouter.POST("/talk/find_talk_list", handler.FindTalkList)       // 分页查询Talk列表

		publicRouter.POST("/talk/find_talk_details", handler.FindTalkDetail)           // 获取Talk详情
		publicRouter.POST("/talk/like_talk", handler.LikeTalk)                         // 点赞Talk
		publicRouter.POST("/talk/find_talk_details_list", handler.FindTalkDetailsList) // 获取Talk详情列表
	}
}
