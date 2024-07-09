package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type ArticleRouter struct {
	svcCtx *svc.ServiceContext
}

func NewArticleRouter(svcCtx *svc.ServiceContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: svcCtx,
	}
}

// 初始化 Article 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = controller.NewArticleController(s.svcCtx)
	// 后台操作
	{
		loginRouter.POST("//admin/article/save_article", handler.SaveArticle)               // 新建文章
		loginRouter.PUT("//admin/article/top_article", handler.TopArticle)                  // 置顶文章
		loginRouter.PUT("//admin/article/logic_delete_article", handler.LogicDeleteArticle) // 逻辑删除文章(假删除)
		loginRouter.DELETE("/admin/delete_article", handler.DeleteArticle)                  // 删除文章
		loginRouter.POST("//admin/article/find_article", handler.FindArticle)               // 查询文章
		loginRouter.POST("//admin/article/find_article_list", handler.FindArticleList)      // 分页查询文章列表
	}
	// 前台操作接口
	{
		publicRouter.POST("/article/find_article_home_list", handler.FindArticleHomeList)            // 首页文章列表
		publicRouter.POST("/article/article_archives", handler.FindArticleArchives)                  // 文章归档
		publicRouter.POST("/article/article_classify_category", handler.FindArticleClassifyCategory) // 根据条件获取Article列表
		publicRouter.POST("/article/article_classify_tag", handler.FindArticleClassifyTag)           // 根据条件获取Article列表
		publicRouter.POST("/article/find_article_recommend", handler.FindArticleRecommend)           // 查询文章详情
		publicRouter.PUT("/article/like_article", handler.LikeArticle)                               // 点赞文章
	}
}
