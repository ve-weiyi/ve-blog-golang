package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
)

type CategoryRouter struct {
	svcCtx *svc.RouterContext
}

func NewCategoryRouter(ctx *svc.RouterContext) *CategoryRouter {
	return &CategoryRouter{
		svcCtx: ctx,
	}
}

// 初始化 Category 路由信息
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	categoryRouter := publicRouter.Group("blog/category")
	categoryTraceRouter := loginRouter.Group("admin/category")

	var self = s.svcCtx.AppController.CategoryController
	{
		categoryRouter.GET("find", self.FindCategory)    // 根据ID获取Category
		categoryRouter.GET("list", self.GetCategoryList) // 获取Category列表
	}
	{
		categoryTraceRouter.POST("create", self.CreateCategory)             // 新建Category
		categoryTraceRouter.DELETE("delete", self.DeleteCategory)           // 删除Category
		categoryTraceRouter.PUT("update", self.UpdateCategory)              // 更新Category
		categoryTraceRouter.DELETE("deleteByIds", self.DeleteCategoryByIds) // 批量删除Category
	}
}
