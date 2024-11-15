package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/gin/config"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
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

	ctx := svctx.NewServiceContext(c)

	// 设置ReleaseMode则不会打印路由注册日志
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	RegisterRouters(engine, ctx)

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
	当前版本: %s
	微信号：wy791422171 QQ：791422171
	默认接口文档地址:http://localhost%s/api/v1/swagger/index.html
`, c.System.Version, address)
	glog.Error(s.ListenAndServe().Error())
}
