package core

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

// 初始化总路由
func RegisterRouters(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {

	// r.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 r.RunTLS("端口","你的cre/pem文件","你的key文件")
	// r.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// r.Use(middleware.GinLogger())  // 访问记录
	r.Use(middleware.Cors())             // 直接放行全部跨域请求
	r.Use(middleware.TraceMiddleware())  // 打印请求的traceId
	r.Use(middleware.LimitIP(serverCtx)) // 限制IP

	//公开接口，不需要token
	publicGroup := r.Group("/")
	publicGroup.Use(middleware.SignToken()) // 签名 校验
	// 后台接口，需要token和角色认证，
	adminGroup := r.Group("/")
	adminGroup.Use(middleware.JwtToken(serverCtx))          // jwt token 校验
	adminGroup.Use(middleware.PermissionHandler(serverCtx)) // 接口权限校验
	adminGroup.Use(middleware.OperationRecord(serverCtx))   // 访问记录 > 私有接口 > 操作记录
	{
		router.NewWebsiteRouter(serverCtx).InitWebsiteRouter(publicGroup, adminGroup)
		router.NewWebsocketRouter(serverCtx).InitWebsocketRouter(publicGroup, adminGroup)
		router.NewAuthRouter(serverCtx).InitAuthRouter(publicGroup, adminGroup)
		router.NewUserRouter(serverCtx).InitUserRouter(publicGroup, adminGroup)
		router.NewApiRouter(serverCtx).InitApiRouter(publicGroup, adminGroup)
		router.NewMenuRouter(serverCtx).InitMenuRouter(publicGroup, adminGroup)
		router.NewRoleRouter(serverCtx).InitRoleRouter(publicGroup, adminGroup)
		router.NewArticleRouter(serverCtx).InitArticleRouter(publicGroup, adminGroup)
		router.NewCategoryRouter(serverCtx).InitCategoryRouter(publicGroup, adminGroup)
		router.NewTagRouter(serverCtx).InitTagRouter(publicGroup, adminGroup)
		router.NewFriendLinkRouter(serverCtx).InitFriendLinkRouter(publicGroup, adminGroup)
		router.NewCommentRouter(serverCtx).InitCommentRouter(publicGroup, adminGroup)
		router.NewPhotoRouter(serverCtx).InitPhotoRouter(publicGroup, adminGroup)
		router.NewPhotoAlbumRouter(serverCtx).InitPhotoAlbumRouter(publicGroup, adminGroup)
		router.NewTalkRouter(serverCtx).InitTalkRouter(publicGroup, adminGroup)
		router.NewPageRouter(serverCtx).InitPageRouter(publicGroup, adminGroup)
		router.NewCaptchaRouter(serverCtx).InitCaptchaRouter(publicGroup, adminGroup)
		router.NewUploadRouter(serverCtx).InitUploadRouter(publicGroup, adminGroup)
		router.NewRemarkRouter(serverCtx).InitRemarkRouter(publicGroup, adminGroup)
		router.NewOperationLogRouter(serverCtx).InitOperationLogRouter(publicGroup, adminGroup)
		router.NewAIRouter(serverCtx).InitAIRouter(publicGroup, adminGroup)
	}
}
