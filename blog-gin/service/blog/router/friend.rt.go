package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FriendRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendRouter(svcCtx *svctx.ServiceContext) *FriendRouter {
	return &FriendRouter{
		svcCtx: svcCtx,
	}
}

func (s *FriendRouter) Register(r *gin.RouterGroup) {
	// Friend
	// [TimeToken VisitLog]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)
		group.Use(s.svcCtx.MiddlewareVisitLog)

		handler := controller.NewFriendController(s.svcCtx)
		// 分页获取友链列表
		group.POST("/friend_link/find_friend_list", handler.FindFriendList)
	}
}
