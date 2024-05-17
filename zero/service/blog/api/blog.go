package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/nacos"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/handler"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	nacosIP        = flag.String("nacos-ip", "120.79.136.81", "Input Your Nacos IP")
	nacosPort      = flag.Int64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUserName  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosDataId    = flag.String("nacos-data-id", "api", "Input Your Nacos DataId")
	nacosGroup     = flag.String("nacos-group", "veweiyi.cn", "nacos group")
	nacosNameSpace = flag.String("nacos-namespace", "test", "Input Your Nacos NameSpaceID")
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
			NameSpaceID: *nacosNameSpace,
			Group:       *nacosGroup,
			DataID:      *nacosDataId,
			RuntimeDir:  "runtime/nacos",
			LogLevel:    "debug",
			Timeout:     5000,
		}

		nacos.New(&nc).Init(func(content string) error {
			log.Println("nacos get config:\n" + content)
			return conf.LoadFromYamlBytes([]byte(content), &c)
		})
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.Use(middlewarex.CtxMetadataHandler)
	server.Use(middlewarex.JwtAuthHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
