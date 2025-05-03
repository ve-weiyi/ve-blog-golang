package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type RemarkRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewRemarkRouter(svcCtx *svctx.ServiceContext) *RemarkRouter {
	return &RemarkRouter{
		svcCtx: svcCtx,
	}
}

func (s *RemarkRouter) Register(r *gin.RouterGroup) {
	// Remark
	// [TimeToken VisitLog]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)
		group.Use(s.svcCtx.MiddlewareVisitLog)

		handler := controller.NewRemarkController(s.svcCtx)
		// 分页获取留言列表
		group.POST("/remark/find_remark_list", handler.FindRemarkList)
	}
	// Remark
	// [TimeToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)

		handler := controller.NewRemarkController(s.svcCtx)
		// 创建留言
		group.POST("/remark/add_remark", handler.AddRemark)
	}
}
