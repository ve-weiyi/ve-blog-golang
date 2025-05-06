package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
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
	// [TimeToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)

		handler := controller.NewCommentController(s.svcCtx)
		// 查询评论列表
		group.POST("/comment/find_comment_list", handler.FindCommentList)
		// 查询最新评论回复列表
		group.POST("/comment/find_comment_recent_list", handler.FindCommentRecentList)
		// 查询评论回复列表
		group.POST("/comment/find_comment_reply_list", handler.FindCommentReplyList)
	}
	// Comment
	// [TimeToken SignToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)
		group.Use(s.svcCtx.MiddlewareSignToken)

		handler := controller.NewCommentController(s.svcCtx)
		// 创建评论
		group.POST("/comment/add_comment", handler.AddComment)
		// 点赞评论
		group.POST("/comment/like_comment", handler.LikeComment)
		// 更新评论
		group.POST("/comment/update_comment", handler.UpdateComment)
	}
}
