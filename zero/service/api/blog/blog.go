package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/swagger"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/common/task"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/handler"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
)

var (
	nacosIP        = flag.String("nacos-ip", "120.79.136.81", "Input Your Nacos IP")
	nacosPort      = flag.Int64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUserName  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosDataId    = flag.String("nacos-data-id", "api", "Input Your Nacos DataId")
	nacosGroup     = flag.String("nacos-group", "veweiyi.cn", "nacos group")
	nacosNameSpace = flag.String("nacos-namespace", "test", "Input Your Nacos NameSpaceId")
)

var configFile = flag.String("f", "", "the config file")

//go:embed docs/blog.json
var docs []byte

func main() {
	flag.Parse()

	var c config.Config
	if *configFile != "" {
		conf.MustLoad(*configFile, &c)
	} else {
		nc := nacos.NacosConfig{
			IP:          *nacosIP,
			Port:        uint64(*nacosPort),
			UserName:    *nacosUserName,
			Password:    *nacosPassword,
			NameSpaceId: *nacosNameSpace,
			Group:       *nacosGroup,
			DataId:      *nacosDataId,
			RuntimeDir:  "runtime/log/nacos",
			LogLevel:    "debug",
			Timeout:     5000,
		}

		nr := nacos.New(&nc)

		content, err := nr.GetConfig()
		if err != nil {
			log.Fatal("nacos get config fail", err)
		}

		err = conf.LoadFromYamlBytes([]byte(content), &c)
		if err != nil {
			log.Fatal(err)
		}
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	swagger.RegisterHttpSwagHandler(server, "/api/v1/swagger/", docs)

	server.Use(middlewarex.NewCtxMetaMiddleware().Handle)
	server.Use(middlewarex.NewAntiReplyMiddleware().Handle)

	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes()
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	return http.StatusInternalServerError, err
	//})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf(`
	默认接口文档地址:http://%s:%d/api/v1/swagger/index.html
`, c.Host, c.Port)

	task.Run(ctx)
	server.Start()
}
