package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ve-weiyi/ve-blog-golang/server/api/router"
	"github.com/ve-weiyi/ve-blog-golang/server/api/router/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/docs"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/middleware"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/trace"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	gin.SetMode(gin.DebugMode)
	ctx := svc.NewRouterContext(&global.CONFIG)
	blogRouter := router.NewRouter(ctx)

	// Generate Swagger JSON file
	global.LOG.Info("register swagger handler")
	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 放行后端静态资源目录，为用户头像和文件提供静态地址
	Router.StaticFS(global.CONFIG.Upload.Local.BasePath, http.Dir(global.CONFIG.Upload.Local.BasePath))

	//Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors())       // 直接放行全部跨域请求
	Router.Use(trace.TraceMiddleware()) // 打印请求的traceId
	//Router.Use(middleware.OperationRecord()) // 操作记录
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	//公开接口，不需要token
	publicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	//publicGroup.Use(middleware.GinLogger())
	publicGroup.Use(middleware.LimitIP())
	{
		// 健康监测
		publicGroup.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version":  "1.0.0",
				"trace_id": c.Request.Context().Value("X-Trace-ID").(string),
			})
		})
	}
	// 后台接口，需要token和角色认证，
	adminGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	// 日志收集
	//adminGroup.Use(middleware.GinLogger())
	// 限制IP
	adminGroup.Use(middleware.LimitIP())
	// 鉴权中间件
	adminGroup.Use(middleware.JwtToken())
	{
		blogRouter.WebsiteRouter.InitWebsiteRouter(publicGroup, adminGroup)
		blogRouter.WebsocketRouter.InitWebsocketRouter(publicGroup, adminGroup)
		blogRouter.AuthRouter.InitAuthRouter(publicGroup, adminGroup)
		blogRouter.UserRouter.InitUserRouter(publicGroup, adminGroup)
		blogRouter.ApiRouter.InitApiRouter(publicGroup, adminGroup)
		blogRouter.MenuRouter.InitMenuRouter(publicGroup, adminGroup)
		blogRouter.RoleRouter.InitRoleRouter(publicGroup, adminGroup)
		blogRouter.ArticleRouter.InitArticleRouter(publicGroup, adminGroup)
		blogRouter.CategoryRouter.InitCategoryRouter(publicGroup, adminGroup)
		blogRouter.TagRouter.InitTagRouter(publicGroup, adminGroup)
		blogRouter.FriendLinkRouter.InitFriendLinkRouter(publicGroup, adminGroup)
		blogRouter.CommentRouter.InitCommentRouter(publicGroup, adminGroup)
		blogRouter.PhotoRouter.InitPhotoRouter(publicGroup, adminGroup)
		blogRouter.PhotoAlbumRouter.InitPhotoAlbumRouter(publicGroup, adminGroup)
		blogRouter.TalkRouter.InitTalkRouter(publicGroup, adminGroup)
		blogRouter.PageRouter.InitPageRouter(publicGroup, adminGroup)
		blogRouter.CaptchaRouter.InitCaptchaRouter(publicGroup, adminGroup)
		blogRouter.UploadRouter.InitUploadRouter(publicGroup, adminGroup)
		blogRouter.RemarkRouter.InitRemarkRouter(publicGroup, adminGroup)
		blogRouter.OperationLogRouter.InitOperationLogRouter(publicGroup, adminGroup)
	}

	global.LOG.Info("router register success")
	return Router
}
