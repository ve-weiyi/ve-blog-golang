package main

import (
	"flag"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/common/middlewarex/metadata"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/config"
	apirpcServer "github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/server/apirpc"
	blogrpcServer "github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/server/blogrpc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		blog.RegisterBlogRpcServer(grpcServer, blogrpcServer.NewBlogRpcServer(ctx))
		blog.RegisterApiRpcServer(grpcServer, apirpcServer.NewApiRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(metadata.AppendToOutgoingContextInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
