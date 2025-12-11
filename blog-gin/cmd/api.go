package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/config"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/initialize"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/logz"
)

var configOpt = &initialize.LoadConfigOption{
	ConfigMode:      "file",
	ConfigType:      "yaml",
	LocalPath:       "config.yaml",
	NacosHost:       "veweiyi.cn",
	NacosPort:       8848,
	NacosNamespace:  "dev",
	NacosDataID:     "ve-blog-golang",
	NacosGroup:      "blog",
	NacosRuntimeDir: "runtime/nacos",
	NacosUsername:   "nacos",
	NacosPassword:   "nacos",
}

func NewApiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "启动接口服务",
		Long:  `启动接口服务`,
		RunE:  runApi,
	}

	// 配置文件相关
	cmd.Flags().StringVarP(&configOpt.ConfigMode, "config", "c", configOpt.ConfigMode, "the way of read config file (file|nacos)")
	cmd.Flags().StringVarP(&configOpt.LocalPath, "filepath", "f", configOpt.LocalPath, "config file path")

	// Nacos 相关
	cmd.Flags().StringVar(&configOpt.NacosHost, "n-host", configOpt.NacosHost, "the host for nacos")
	cmd.Flags().Uint64Var(&configOpt.NacosPort, "n-port", configOpt.NacosPort, "the port for nacos")
	cmd.Flags().StringVar(&configOpt.NacosNamespace, "n-namespace", configOpt.NacosNamespace, "the namespace for nacos")
	cmd.Flags().StringVar(&configOpt.NacosDataID, "n-data-id", configOpt.NacosDataID, "the DataId for nacos")
	cmd.Flags().StringVar(&configOpt.NacosGroup, "n-group", configOpt.NacosGroup, "the group for nacos")
	cmd.Flags().StringVar(&configOpt.NacosUsername, "n-user", configOpt.NacosUsername, "the user for nacos")
	cmd.Flags().StringVar(&configOpt.NacosPassword, "n-password", configOpt.NacosPassword, "the password for nacos")

	return cmd
}

func runApi(cmd *cobra.Command, args []string) error {
	c, err := initialize.LoadConfig(configOpt)
	if err != nil {
		log.Printf("failed to initialize config: %v\n", err)
		os.Exit(1)
	}

	RunServer(c)
	return nil
}

func RunServer(c *config.Config) {
	// 初始化zap日志库
	initialize.SetLog(c.Zap)

	ctx := svctx.NewServiceContext(c)

	// 设置ReleaseMode则不会打印路由注册日志
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	api.RegisterRouters(engine, ctx)

	logz.Infof("register router success")

	address := fmt.Sprintf(":%d", c.System.Port)
	s := &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	logz.Infof("run server on http://localhost:%v success", c.System.Port)

	fmt.Printf(`
	欢迎使用 ve-blog-golang
	当前版本: %s
	微信号：wy791422171 QQ：791422171
	默认接口文档地址:http://localhost%s/api/v1/swagger/index.html
`, c.System.Version, address)
	fmt.Println(s.ListenAndServe().Error())
}
