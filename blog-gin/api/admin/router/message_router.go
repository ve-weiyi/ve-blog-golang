package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/handler"
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
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		h := handler.NewMessageController(s.svcCtx)
		// 删除留言
		group.DELETE("/message/deletes_message", h.DeletesMessage)
		// 分页获取留言列表
		group.POST("/message/find_message_list", h.FindMessageList)
		// 更新留言状态
		group.PUT("/message/update_message_status", h.UpdateMessageStatus)
	}
}
