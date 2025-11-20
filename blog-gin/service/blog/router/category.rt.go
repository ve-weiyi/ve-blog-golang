package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CategoryRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryRouter(svcCtx *svctx.ServiceContext) *CategoryRouter {
	return &CategoryRouter{
		svcCtx: svcCtx,
	}
}

func (s *CategoryRouter) Register(r *gin.RouterGroup) {
	// Category
	// [TerminalToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.TerminalToken)

		handler := controller.NewCategoryController(s.svcCtx)
		// 分页获取文章分类列表
		group.POST("/category/find_category_list", handler.FindCategoryList)
	}
}
