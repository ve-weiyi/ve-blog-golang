package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/swagger"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/docs"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/handler"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
)

var (
	nacosIP        = flag.String("nacos-ip", "veweiyi.cn", "Input Your Nacos IP")
	nacosPort      = flag.Int64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUserName  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosNameSpace = flag.String("nacos-namespace", "test", "Input Your Nacos NameSpaceId")
	nacosGroup     = flag.String("nacos-group", "veweiyi.cn", "nacos group")
	nacosDataId    = flag.String("nacos-data-id", "api", "Input Your Nacos DataId")
)

var configFile = flag.String("f", "", "the config file")

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
			RuntimeDir:  "runtime/blog-api/nacos",
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

	swagger.RegisterHttpSwagHandler(server, "/api/v1/swagger/", []byte(docs.Docs))

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

	server.Start()
}
