package logic

import "github.com/gin-gonic/gin"

// 初始化 Api 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ApiRouter) InitApiRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitApiGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.ApiController
	{
		loginRouter.POST("api/list/details", handler.FindApiListDetails) // 获取Api列表
	}
}

// 初始化 Category 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitCategoryGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.CategoryController
	{
		publicRouter.POST("category/list/details", handler.FindCategoryListDetails) // 查询Category详情列表
	}
}

// 初始化 Comment 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CommentRouter) InitCommentRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitCommentGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.CommentController

	{
		publicRouter.POST("comment/:id/like", handler.LikeComment)                // 点赞评论
		publicRouter.POST("comment/:id/reply_list", handler.FindCommentReplyList) // 分页查询Comment列表
		publicRouter.POST("comment/list/details", handler.FindCommentListDetails) // 分页查询Comment列表
		publicRouter.POST("comment/list/back", handler.FindCommentBackList)       // 分页查询Comment列表
	}
}

// 初始化 Role 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RoleRouter) InitRoleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitRoleGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.RoleController
	{
		loginRouter.POST("role/list/details", handler.FindRoleListDetails)
	}
}

// 初始化 Menu 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *MenuRouter) InitMenuRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitMenuGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.MenuController
	{
		loginRouter.POST("menu/list/details", handler.FindMenuListDetails)
	}
}

// 初始化 Article 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *ArticleRouter) InitArticleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitArticleGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.ArticleController
	{
		publicRouter.GET("article/:id/details", handler.GetArticleDetails) // 获取Article详情
		publicRouter.PUT("article/:id/like", handler.LikeArticle)          // 点赞文章

		publicRouter.GET("article/archives", handler.GetArticleArchives)                // 文章归档
		publicRouter.POST("article/list/condition", handler.FindArticleListByCondition) // 根据条件获取Article列表
	}
}

// 初始化 Tag 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TagRouter) InitTagRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitTagGenRouter(publicRouter, loginRouter)
}

// 初始化 Photo 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoRouter) InitPhotoRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPhotoGenRouter(publicRouter, loginRouter)
}

// 初始化 PhotoAlbum 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PhotoAlbumRouter) InitPhotoAlbumRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPhotoAlbumGenRouter(publicRouter, loginRouter)
}

// 初始化 Page 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *PageRouter) InitPageRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitPageGenRouter(publicRouter, loginRouter)
}

// 初始化 Talk 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TalkRouter) InitTalkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitTalkGenRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.AppController.TalkController
	{
		publicRouter.GET("talk/:id/details", handler.FindTalkDetail)        // 获取Talk详情
		publicRouter.POST("talk/list/details", handler.FindTalkListDetails) // 获取Talk详情列表
	}
}

// 初始化 FriendLink 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *FriendLinkRouter) InitFriendLinkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitFriendLinkGenRouter(publicRouter, loginRouter)
}

// 初始化 OperationLog 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *OperationLogRouter) InitOperationLogRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitOperationLogGenRouter(publicRouter, loginRouter)
}

// 初始化 Remark 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RemarkRouter) InitRemarkRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitRemarkGenRouter(publicRouter, loginRouter)
}
