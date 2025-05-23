package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CommentRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentRouter(svcCtx *svctx.ServiceContext) *CommentRouter {
	return &CommentRouter{
		svcCtx: svcCtx,
	}
}

func (s *CommentRouter) Register(r *gin.RouterGroup) {
	// Comment
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewCommentController(s.svcCtx)
		// 批量删除评论
		group.DELETE("/comment/batch_delete_comment", handler.BatchDeleteComment)
		// 删除评论
		group.DELETE("/comment/delete_comment", handler.DeleteComment)
		// 查询评论列表(后台)
		group.POST("/comment/find_comment_back_list", handler.FindCommentBackList)
		// 更新评论审核状态
		group.PUT("/comment/update_comment_review", handler.UpdateCommentReview)
	}
}
