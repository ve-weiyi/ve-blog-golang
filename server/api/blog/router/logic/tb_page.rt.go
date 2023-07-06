package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type PageRouter struct {
	svcCtx *svc.RouterContext
}

func NewPageRouter(ctx *svc.RouterContext) *PageRouter {
	return &PageRouter{
		svcCtx: ctx,
	}
}

// 初始化 Page 路由信息
func (s *PageRouter) InitPageRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	pageRouter := publicRouter.Group("blog/page")
	pageTraceRouter := loginRouter.Group("admin/page")

	var self = s.svcCtx.AppController.PageController
	{
		pageRouter.GET("find", self.FindPage)    // 根据ID获取Page
		pageRouter.GET("list", self.GetPageList) // 获取Page列表
	}
	{
		pageTraceRouter.POST("create", self.CreatePage)             // 新建Page
		pageTraceRouter.DELETE("delete", self.DeletePage)           // 删除Page
		pageTraceRouter.PUT("update", self.UpdatePage)              // 更新Page
		pageTraceRouter.DELETE("deleteByIds", self.DeletePageByIds) // 批量删除Page
	}
}
