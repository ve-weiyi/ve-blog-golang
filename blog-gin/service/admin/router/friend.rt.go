package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
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
	// [JwtToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.MiddlewareJwtToken)
		group.Use(s.svcCtx.MiddlewarePermission)
		group.Use(s.svcCtx.MiddlewareOperationLog)

		handler := controller.NewFriendController(s.svcCtx)
		// 创建友链
		group.POST("/friend/add_friend", handler.AddFriend)
		// 删除友链
		group.DELETE("/friend/deletes_friend", handler.DeletesFriend)
		// 分页获取友链列表
		group.POST("/friend/find_friend_list", handler.FindFriendList)
		// 更新友链
		group.PUT("/friend/update_friend", handler.UpdateFriend)
	}
}
