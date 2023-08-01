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
func (s *ArticleRouter) InitArticleGenRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.ArticleController
	{
		publicRouter.POST("article", handler.CreateArticle)       // 新建Article
		publicRouter.PUT("article", handler.UpdateArticle)        // 更新Article
		publicRouter.DELETE("article/:id", handler.DeleteArticle) // 删除Article
		publicRouter.GET("article/:id", handler.FindArticle)      // 查询Article

		publicRouter.DELETE("article/batch_delete", handler.DeleteArticleByIds) // 批量删除Article列表
		publicRouter.POST("article/list", handler.FindArticleList)              // 分页查询Article列表
	}
}
