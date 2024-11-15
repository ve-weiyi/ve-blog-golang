package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel/trace"

	"github.com/ve-weiyi/ve-blog-golang/gin/docs"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/middleware"
	"github.com/ve-weiyi/ve-blog-golang/gin/service/admin"
	"github.com/ve-weiyi/ve-blog-golang/gin/service/blog"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

// 初始化总路由
func RegisterRouters(engine *gin.Engine, svCtx *svctx.ServiceContext) {
	r := engine.Group("")

	RegisterStaticHandlers(r, svCtx)

	// r.Use(middleware.GinLogger())  // 访问记录，gin.Default()自带了
	// r.Use(middleware.GinRecovery(true)) // 捕获panic，gin.Default()自带了
	// r.Use(middleware.LoadTls())  // 使用https，一般在nginx处理。前往 core/server.go 将启动模式 更变为 r.RunTLS("端口","你的cre/pem文件","你的key文件")
	// r.Use(middleware.Limit(svCtx.RedisEngin)) // 请求限频，一般在nginx层处理。

	r.Use(middleware.Cors())  // 直接放行全部跨域请求
	r.Use(middleware.Trace()) // 打印请求的traceId

	RegisterPingHandlers(r, svCtx)
	admin.RegisterHandlers(r, svCtx)
	blog.RegisterHandlers(r, svCtx)
}

// 健康检查
func RegisterPingHandlers(r *gin.RouterGroup, svCtx *svctx.ServiceContext) {
	now := time.Now()
	// 健康监测
	r.GET("/api/v1/version", func(c *gin.Context) {
		traceID := trace.SpanContextFromContext(c.Request.Context())

		c.JSON(http.StatusOK, gin.H{
			"version":  svCtx.Config.System.Version,
			"runtime":  now.String(),
			"trace_id": traceID,
		})
	})
}

// 静态资源处理
func RegisterStaticHandlers(r *gin.RouterGroup, svCtx *svctx.ServiceContext) {
	staticRouter := r.Group("/api/v1")
	// 放行后端静态资源目录，为用户头像和文件提供静态地址
	staticRouter.StaticFS(svCtx.Config.System.RuntimePath, http.Dir(svCtx.Config.System.RuntimePath))

	// Generate Swagger JSON file
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", svCtx.Config.System.Port)
	docs.SwaggerInfo.Version = svCtx.Config.System.Version
	staticRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
