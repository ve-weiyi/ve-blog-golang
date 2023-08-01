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
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	ctx := svc.NewRouterContext(&global.CONFIG)
	blogRouter := router.NewRouter(ctx)
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.CONFIG.Upload.Local.Path, http.Dir(global.CONFIG.Upload.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.LOG.Info("use middleware cors")
	// Generate Swagger JSON file
	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//公开接口，不需要token
	publicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	//publicGroup.Use(middleware.GinLogger())
	publicGroup.Use(middleware.LimitIP())
	publicGroup.Use(middleware.OperationRecord())
	{
		// 健康监测
		publicGroup.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, "1.0.0")
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
	// 操作记录
	adminGroup.Use(middleware.OperationRecord())
	{
		blogRouter.BlogRouter.InitBlogRouter(publicGroup, adminGroup)
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
		blogRouter.CaptchaRouter.InitCaptchaRouter(publicGroup, adminGroup)
		blogRouter.UploadRouter.InitUploadRouter(publicGroup, adminGroup)
		blogRouter.RemarkRouter.InitRemarkRouter(publicGroup, adminGroup)
		blogRouter.OperationLogRouter.InitOperationLogRouter(publicGroup, adminGroup)
	}

	global.LOG.Info("router register success")
	return Router
}
