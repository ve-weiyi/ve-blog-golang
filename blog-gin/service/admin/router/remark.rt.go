package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
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
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewRemarkController(s.svcCtx)
		// 删除留言
		group.DELETE("/remark/deletes_remark", handler.DeletesRemark)
		// 分页获取留言列表
		group.POST("/remark/find_remark_list", handler.FindRemarkList)
		// 更新留言
		group.PUT("/remark/update_remark_review", handler.UpdateRemarkReview)
	}
}
