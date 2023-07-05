package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"
)

type ArticleRouter struct {
	svcCtx *svc.RouterContext
}

func NewArticleRouter(ctx *svc.RouterContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: ctx,
	}
}

// 初始化 Article 路由信息
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	articleRouter := publicRouter.Group("blog/article")
	articleTraceRouter := loginRouter.Group("admin/article")

	var self = s.svcCtx.AppController.ArticleController
	{
		articleRouter.POST("find", self.FindArticle)    // 根据ID获取Article
		articleRouter.POST("list", self.GetArticleList) // 获取Article列表

		articleRouter.POST("condition", self.GetArticleListByCondition) // 根据条件获取Article列表
		articleRouter.GET("archives", self.GetArticleArchives)          // 文章归档
	}
	{
		articleTraceRouter.POST("create", self.CreateArticle)             // 新建Article
		articleTraceRouter.DELETE("delete", self.DeleteArticle)           // 删除Article
		articleTraceRouter.PUT("update", self.UpdateArticle)              // 更新Article
		articleTraceRouter.DELETE("deleteByIds", self.DeleteArticleByIds) // 批量删除Article
	}
}
