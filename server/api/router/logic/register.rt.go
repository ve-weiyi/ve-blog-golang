package logic

import (
	"github.com/gin-gonic/gin"
)

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ApiRouter) InitApiRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitApiBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.ApiController
	{
		loginRouter.POST("api/list/details", handler.FindApiDetailsList) // 获取Api列表
	}
}

// 初始化 Category 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitCategoryBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.CategoryController
	{
		publicRouter.POST("category/list/details", handler.FindCategoryDetailsList) // 查询Category详情列表
	}
}

// 初始化 Comment 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CommentRouter) InitCommentRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitCommentBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.CommentController

	{
		publicRouter.POST("comment/:id/like", handler.LikeComment)                // 点赞评论
		publicRouter.POST("comment/:id/reply_list", handler.FindCommentReplyList) // 分页查询Comment列表
		publicRouter.POST("comment/list/details", handler.FindCommentDetailsList) // 分页查询Comment列表
		publicRouter.POST("comment/list/back", handler.FindCommentListBack)       // 分页查询Comment列表
	}
}

// 初始化 Role 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RoleRouter) InitRoleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitRoleBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.RoleController
	{
		loginRouter.POST("role/list/details", handler.FindRoleDetailsList)
	}
}

// 初始化 Menu 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *MenuRouter) InitMenuRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitMenuBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.MenuController
	{
		loginRouter.POST("menu/list/details", handler.FindMenuDetailsList)
	}
}

// 初始化 Article 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitArticleBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.ArticleController
	{
		publicRouter.GET("article/:id/details", handler.FindArticleDetails) // 获取Article详情
		publicRouter.PUT("article/:id/like", handler.LikeArticle)           // 点赞文章

		publicRouter.POST("article/archives", handler.FindArticleArchives)              // 文章归档
		publicRouter.POST("article/list/details", handler.FindArticleDetailsList)       // 根据条件获取Article列表
		publicRouter.POST("article/list/condition", handler.FindArticleListByCondition) // 根据条件获取Article列表
	}
}

// 初始化 Tag 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TagRouter) InitTagRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitTagBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.TagController
	{
		loginRouter.POST("tag/list/details", handler.FindTagDetailsList)
	}
}

// 初始化 Photo 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoRouter) InitPhotoRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPhotoBasicRouter(publicRouter, loginRouter)
}

// 初始化 PhotoAlbum 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoAlbumRouter) InitPhotoAlbumRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPhotoAlbumBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.PhotoAlbumController
	{
		publicRouter.GET("photo_album/:id/details", handler.FindPhotoAlbumDetails)       // 获取PhotoAlbum详情
		publicRouter.POST("photo_album/list/details", handler.FindPhotoAlbumDetailsList) // 获取PhotoAlbum详情列表
	}
}

// 初始化 Page 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PageRouter) InitPageRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPageBasicRouter(publicRouter, loginRouter)
}

// 初始化 Talk 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TalkRouter) InitTalkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitTalkBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.TalkController
	{
		publicRouter.GET("talk/:id/details", handler.FindTalkDetail)        // 获取Talk详情
		publicRouter.POST("talk/list/details", handler.FindTalkDetailsList) // 获取Talk详情列表
	}
}

// 初始化 FriendLink 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *FriendLinkRouter) InitFriendLinkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitFriendLinkBasicRouter(publicRouter, loginRouter)
}

// 初始化 OperationLog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *OperationLogRouter) InitOperationLogRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitOperationLogBasicRouter(publicRouter, loginRouter)
}

// 初始化 Remark 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RemarkRouter) InitRemarkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitRemarkBasicRouter(publicRouter, loginRouter)
}
