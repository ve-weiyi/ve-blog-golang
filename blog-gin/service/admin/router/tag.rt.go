package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type TagRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewTagRouter(svcCtx *svctx.ServiceContext) *TagRouter {
	return &TagRouter{
		svcCtx: svcCtx,
	}
}

func (s *TagRouter) Register(r *gin.RouterGroup) {
	// Tag
	// [AdminToken Permission OperationLog]
	{
		group := r.Group("/admin-api/v1")
		group.Use(s.svcCtx.AdminToken)
		group.Use(s.svcCtx.Permission)
		group.Use(s.svcCtx.OperationLog)

		handler := controller.NewTagController(s.svcCtx)
		// 创建标签
		group.POST("/tag/add_tag", handler.AddTag)
		// 删除标签
		group.DELETE("/tag/deletes_tag", handler.DeletesTag)
		// 分页获取标签列表
		group.POST("/tag/find_tag_list", handler.FindTagList)
		// 更新标签
		group.PUT("/tag/update_tag", handler.UpdateTag)
	}
}
