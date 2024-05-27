package initialize

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/router"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/docs"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Generate Swagger JSON file
	glog.Info("register swagger handler")
	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 放行后端静态资源目录，为用户头像和文件提供静态地址
	Router.StaticFS(global.CONFIG.Upload.Local.BasePath, http.Dir(global.CONFIG.Upload.Local.BasePath))

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors())            // 直接放行全部跨域请求
	Router.Use(middleware.TraceMiddleware()) // 打印请求的traceId
	//Router.Use(middleware.GinLogger())  // 访问记录
	Router.Use(middleware.LimitIP()) // 限制IP

	//公开接口，不需要token
	publicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	publicGroup.Use(middleware.SignToken()) // 签名 校验
	{
		now := time.Now()
		// 健康监测
		publicGroup.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"runtime":  now.String(),
				"version":  "1.0.0",
				"trace_id": c.Request.Context().Value("X-Trace-ID").(string),
			})
		})
	}
	// 后台接口，需要token和角色认证，
	adminGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	adminGroup.Use(middleware.JwtToken())          // jwt token 校验
	adminGroup.Use(middleware.PermissionHandler()) // 接口权限校验
	adminGroup.Use(middleware.OperationRecord())   // 访问记录 > 私有接口 > 操作记录
	{
		ctx := svc.NewServiceContext(&global.CONFIG)
		router.NewWebsiteRouter(ctx).InitWebsiteRouter(publicGroup, adminGroup)
		router.NewWebsocketRouter(ctx).InitWebsocketRouter(publicGroup, adminGroup)
		router.NewAuthRouter(ctx).InitAuthRouter(publicGroup, adminGroup)
		router.NewUserRouter(ctx).InitUserRouter(publicGroup, adminGroup)
		router.NewApiRouter(ctx).InitApiRouter(publicGroup, adminGroup)
		router.NewMenuRouter(ctx).InitMenuRouter(publicGroup, adminGroup)
		router.NewRoleRouter(ctx).InitRoleRouter(publicGroup, adminGroup)
		router.NewArticleRouter(ctx).InitArticleRouter(publicGroup, adminGroup)
		router.NewCategoryRouter(ctx).InitCategoryRouter(publicGroup, adminGroup)
		router.NewTagRouter(ctx).InitTagRouter(publicGroup, adminGroup)
		router.NewFriendLinkRouter(ctx).InitFriendLinkRouter(publicGroup, adminGroup)
		router.NewCommentRouter(ctx).InitCommentRouter(publicGroup, adminGroup)
		router.NewPhotoRouter(ctx).InitPhotoRouter(publicGroup, adminGroup)
		router.NewPhotoAlbumRouter(ctx).InitPhotoAlbumRouter(publicGroup, adminGroup)
		router.NewTalkRouter(ctx).InitTalkRouter(publicGroup, adminGroup)
		router.NewPageRouter(ctx).InitPageRouter(publicGroup, adminGroup)
		router.NewCaptchaRouter(ctx).InitCaptchaRouter(publicGroup, adminGroup)
		router.NewUploadRouter(ctx).InitUploadRouter(publicGroup, adminGroup)
		router.NewRemarkRouter(ctx).InitRemarkRouter(publicGroup, adminGroup)
		router.NewOperationLogRouter(ctx).InitOperationLogRouter(publicGroup, adminGroup)
		router.NewAIRouter(ctx).InitAIRouter(publicGroup, adminGroup)
	}

	glog.Info("router register success")
	return Router
}
