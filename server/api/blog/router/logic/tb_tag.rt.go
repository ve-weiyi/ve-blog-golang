package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router/svc"
)

type TagRouter struct {
	svcCtx *svc.RouterContext
}

func NewTagRouter(ctx *svc.RouterContext) *TagRouter {
	return &TagRouter{
		svcCtx: ctx,
	}
}

// 初始化 Tag 路由信息
func (s *TagRouter) InitTagRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	tagRouter := publicRouter.Group("blog/tag")
	tagTraceRouter := loginRouter.Group("admin/tag")

	var self = s.svcCtx.AppController.TagController
	{
		tagRouter.GET("find", self.FindTag)    // 根据ID获取Tag
		tagRouter.GET("list", self.GetTagList) // 获取Tag列表
	}
	{
		tagTraceRouter.POST("create", self.CreateTag)             // 新建Tag
		tagTraceRouter.DELETE("delete", self.DeleteTag)           // 删除Tag
		tagTraceRouter.PUT("update", self.UpdateTag)              // 更新Tag
		tagTraceRouter.DELETE("deleteByIds", self.DeleteTagByIds) // 批量删除Tag
	}
}
