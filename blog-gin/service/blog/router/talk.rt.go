package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type TalkRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkRouter(svcCtx *svctx.ServiceContext) *TalkRouter {
	return &TalkRouter{
		svcCtx: svcCtx,
	}
}

func (s *TalkRouter) Register(r *gin.RouterGroup) {
	// Talk
	// [TerminalToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.TerminalToken)

		handler := controller.NewTalkController(s.svcCtx)
		// 分页获取说说列表
		group.POST("/talk/find_talk_list", handler.FindTalkList)
		// 查询说说
		group.POST("/talk/get_talk", handler.GetTalk)
		// 点赞说说
		group.PUT("/talk/like_talk", handler.LikeTalk)
	}
}
