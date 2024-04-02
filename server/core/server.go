package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/docs"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type server interface {
	ListenAndServe() (err error)
	RegisterOnShutdown(f func())
	// 优雅的停机重启
	Shutdown(ctx context.Context) error
}

func RunWindowsServer(c *config.Config) {
	// 初始化zap日志库
	SetLog(c.Zap)

	// 设置ReleaseMode则不会打印路由注册日志
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	r := engine.Group(c.System.RouterPrefix)

	// 放行后端静态资源目录，为用户头像和文件提供静态地址
	r.StaticFS(c.System.RuntimePath, http.Dir(c.System.RuntimePath))

	// Generate Swagger JSON file
	docs.SwaggerInfo.Version = c.System.Version
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctx := svc.NewServiceContext(c)
	RegisterRouters(r, ctx)

	glog.Info("register router success")

	address := fmt.Sprintf(":%d", c.System.Port)
	var s server
	s = &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	glog.Infof("run server on http://localhost:%v success", c.System.Port)

	fmt.Printf(`
	欢迎使用 ve-blog-golang
	当前版本: v1.0.0
	微信号：wy791422171 QQ：791422171
	默认接口文档地址:http://127.0.0.1%s/api/v1/swagger/index.html
	默认前端运行地址:http://127.0.0.1:9090
`, address)
	glog.Error(s.ListenAndServe().Error())
}
