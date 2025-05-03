package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type VisitLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogRouter(svcCtx *svctx.ServiceContext) *VisitLogRouter {
	return &VisitLogRouter{
		svcCtx: svcCtx,
	}
}

func (s *VisitLogRouter) Register(r *gin.RouterGroup) {
	// VisitLog
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewVisitLogController(s.svcCtx)
		// 删除操作记录
		group.DELETE("/visit_log/deletes_visit_log", handler.DeletesVisitLog)
		// 分页获取操作记录列表
		group.POST("/visit_log/find_visit_log_list", handler.FindVisitLogList)
	}
}
