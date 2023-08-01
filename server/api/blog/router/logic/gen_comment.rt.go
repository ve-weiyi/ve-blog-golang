package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type CommentRouter struct {
	svcCtx *svc.RouterContext
}

func NewCommentRouter(svcCtx *svc.RouterContext) *CommentRouter {
	return &CommentRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Comment 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CommentRouter) InitCommentGenRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.CommentController
	{
		publicRouter.POST("comment", handler.CreateComment)       // 新建Comment
		publicRouter.PUT("comment", handler.UpdateComment)        // 更新Comment
		publicRouter.DELETE("comment/:id", handler.DeleteComment) // 删除Comment
		publicRouter.GET("comment/:id", handler.FindComment)      // 查询Comment

		publicRouter.DELETE("comment/batch_delete", handler.DeleteCommentByIds) // 批量删除Comment列表
		publicRouter.POST("comment/list", handler.FindCommentList)              // 分页查询Comment列表
	}
}
