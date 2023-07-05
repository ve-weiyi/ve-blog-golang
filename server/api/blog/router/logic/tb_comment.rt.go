package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
)

type CommentRouter struct {
	svcCtx *svc.RouterContext
}

func NewCommentRouter(ctx *svc.RouterContext) *CommentRouter {
	return &CommentRouter{
		svcCtx: ctx,
	}
}

// 初始化 Comment 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CommentRouter) InitCommentRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.CommentController
	{
		loginRouter.POST("comment/create", handler.CreateComment)    // 新建Comment
		publicRouter.PUT("comment/update", handler.UpdateComment)    // 更新Comment
		publicRouter.DELETE("comment/delete", handler.DeleteComment) // 删除Comment
		publicRouter.POST("comment/query", handler.GetComment)       // 查询Comment

		publicRouter.DELETE("comment/deleteByIds", handler.DeleteCommentByIds) // 批量删除Comment列表
		publicRouter.POST("comment/list", handler.FindCommentList)             // 分页查询Comment列表

	}
	{
		publicRouter.POST("comment/:id/reply_list", handler.ReplyComment) // 查询评论回复列表
		loginRouter.POST("comment/:id/like", handler.LikeComment)         // 评论点赞
	}
}
