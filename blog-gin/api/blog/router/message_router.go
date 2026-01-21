package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/blog/handler"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type MessageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageRouter(svcCtx *svctx.ServiceContext) *MessageRouter {
	return &MessageRouter{
		svcCtx: svcCtx,
	}
}

func (s *MessageRouter) Register(r *gin.RouterGroup) {
	// Message
	// []
	{
		group := r.Group("/blog-api/v1")

		h := handler.NewMessageController(s.svcCtx)
		// 分页获取留言列表
		group.POST("/message/find_message_list", h.FindMessageList)
	}
	// Message
	// []
	{
		group := r.Group("/blog-api/v1")

		h := handler.NewMessageController(s.svcCtx)
		// 创建留言
		group.POST("/message/add_message", h.AddMessage)
	}
}
