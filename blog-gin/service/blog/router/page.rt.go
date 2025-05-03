package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type PageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewPageRouter(svcCtx *svctx.ServiceContext) *PageRouter {
	return &PageRouter{
		svcCtx: svcCtx,
	}
}

func (s *PageRouter) Register(r *gin.RouterGroup) {
	// Page
	// [TimeToken]
	{
		group := r.Group("/blog-api/v1")
		group.Use(s.svcCtx.MiddlewareTimeToken)

		handler := controller.NewPageController(s.svcCtx)
		// 分页获取页面列表
		group.POST("/page/find_page_list", handler.FindPageList)
	}
}
