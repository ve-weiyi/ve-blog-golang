package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type TagRouter struct {
	svcCtx *svc.RouterContext
}

func NewTagRouter(svcCtx *svc.RouterContext) *TagRouter {
	return &TagRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Tag 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TagRouter) InitTagRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.TagController
	{
		loginRouter.POST("/tag/create_tag", handler.CreateTag)            // 新建Tag
		loginRouter.PUT("/tag/update_tag", handler.UpdateTag)             // 更新Tag
		loginRouter.DELETE("/tag/delete_tag", handler.DeleteTag)          // 删除Tag
		loginRouter.DELETE("/tag/delete_tag_list", handler.DeleteTagList) // 批量删除Tag列表

		publicRouter.POST("/tag/find_tag", handler.FindTag)          // 查询Tag
		publicRouter.POST("/tag/find_tag_list", handler.FindTagList) // 分页查询Tag列表
	}
}
