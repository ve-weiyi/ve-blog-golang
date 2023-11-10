package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
)

type ArticleRouter struct {
	svcCtx *svc.RouterContext
}

func NewArticleRouter(svcCtx *svc.RouterContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Article 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.ArticleController
	{
		publicRouter.POST("article", handler.SaveArticle)         // 新建Article
		publicRouter.DELETE("article/:id", handler.DeleteArticle) // 删除Article
		publicRouter.GET("article/:id", handler.FindArticle)      // 查询Article

		publicRouter.DELETE("article/batch_delete", handler.DeleteArticleByIds) // 批量删除Article列表
		publicRouter.POST("article/list", handler.FindArticleList)              // 分页查询Article列表
	}
	{
		publicRouter.POST("article/delete", handler.UpdateArticleDelete)   // 逻辑删除文章
		publicRouter.POST("article/top", handler.UpdateArticleTop)         // 置顶文章
		publicRouter.POST("article/archives", handler.FindArticleArchives) // 文章归档
		publicRouter.POST("article/series", handler.FindArticleSeries)     // 根据条件获取Article列表

		publicRouter.POST("article/:id/recommend", handler.FindArticleRecommend) // 查询文章推荐列表
		publicRouter.PUT("article/:id/like", handler.LikeArticle)                // 点赞文章
	}
}
