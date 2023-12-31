package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
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
func (s *TagRouter) InitTagBasicRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.TagController
	{
		publicRouter.POST("tag", handler.CreateTag)       // 新建Tag
		publicRouter.PUT("tag", handler.UpdateTag)        // 更新Tag
		publicRouter.DELETE("tag/:id", handler.DeleteTag) // 删除Tag
		publicRouter.GET("tag/:id", handler.FindTag)      // 查询Tag

		publicRouter.DELETE("tag/batch_delete", handler.DeleteTagByIds) // 批量删除Tag列表
		publicRouter.POST("tag/list", handler.FindTagList)              // 分页查询Tag列表
	}
}
