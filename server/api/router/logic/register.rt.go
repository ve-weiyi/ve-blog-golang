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
		publicRouter.POST("api/sync", handler.SyncApiList) // 同步Api列表

		loginRouter.POST("api/details_list", handler.FindApiDetailsList) // 获取Api列表
	}
}

// 初始化 Category 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *CategoryRouter) InitCategoryRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitCategoryBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.CategoryController
	{
		publicRouter.POST("category/details_list", handler.FindCategoryDetailsList) // 查询Category详情列表
	}
}

// 初始化 Role 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *RoleRouter) InitRoleRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	var handler = s.svcCtx.RoleController
	{
		loginRouter.POST("role", handler.CreateRole)       // 新建Role
		loginRouter.PUT("role", handler.UpdateRole)        // 更新Role
		loginRouter.DELETE("role/:id", handler.DeleteRole) // 删除Role
		loginRouter.GET("role/:id", handler.FindRole)      // 查询Role

		loginRouter.DELETE("role/batch_delete", handler.DeleteRoleByIds) // 批量删除Role列表
		loginRouter.POST("role/list", handler.FindRoleList)              // 分页查询Role列表

		loginRouter.POST("role/details_list", handler.FindRoleDetailsList)
	}
}

// 初始化 Menu 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *MenuRouter) InitMenuRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitMenuBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.MenuController
	{
		loginRouter.POST("menu/details_list", handler.FindMenuDetailsList)
	}
}

// 初始化 Tag 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *TagRouter) InitTagRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {
	s.InitTagBasicRouter(publicRouter, loginRouter)
	var handler = s.svcCtx.TagController
	{
		loginRouter.POST("tag/details_list", handler.FindTagDetailsList)
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
		publicRouter.POST("photo_album/details_list", handler.FindPhotoAlbumDetailsList) // 获取PhotoAlbum详情列表
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
		publicRouter.POST("talk/:id/like", handler.LikeTalk)                // 点赞Talk
		publicRouter.POST("talk/details_list", handler.FindTalkDetailsList) // 获取Talk详情列表
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
