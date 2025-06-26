package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
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
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewTalkController(s.svcCtx)
		// 创建说说
		group.POST("/talk/add_talk", handler.AddTalk)
		// 删除说说
		group.DELETE("/talk/delete_talk", handler.DeleteTalk)
		// 分页获取说说列表
		group.POST("/talk/find_talk_list", handler.FindTalkList)
		// 查询说说
		group.POST("/talk/get_talk", handler.GetTalk)
		// 更新说说
		group.PUT("/talk/update_talk", handler.UpdateTalk)
	}
}
