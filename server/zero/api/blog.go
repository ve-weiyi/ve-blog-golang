package main

import (
	"flag"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/middlewarex/metadata"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/handler"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.Use(metadata.CtxMetadataHandel)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
