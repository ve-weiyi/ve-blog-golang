package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type NoticeRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewNoticeRouter(svcCtx *svctx.ServiceContext) *NoticeRouter {
	return &NoticeRouter{
		svcCtx: svcCtx,
	}
}

func (s *NoticeRouter) Register(r *gin.RouterGroup) {
	// Notice
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		h := handler.NewNoticeController(s.svcCtx)
		// 创建通知
		group.POST("/notice/add_notice", h.AddNotice)
		// 删除通知
		group.DELETE("/notice/deletes_notice", h.DeletesNotice)
		// 分页获取通知列表
		group.POST("/notice/find_notice_list", h.FindNoticeList)
		// 查询用户通知列表
		group.POST("/notice/find_user_notice_list", h.FindUserNoticeList)
		// 查询通知详情
		group.GET("/notice/get_notice", h.GetNotice)
		// 更新通知
		group.PUT("/notice/update_notice", h.UpdateNotice)
		// 更新通知状态
		group.PUT("/notice/update_notice_status", h.UpdateNoticeStatus)
	}
}
