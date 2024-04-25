package main

import (
	"flag"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/config"
	apirpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/apirpc"
	articlerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/articlerpc"
	authrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/authrpc"
	categoryrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/categoryrpc"
	configrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/configrpc"
	friendlinkrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/friendlinkrpc"
	menurpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/menurpc"
	remarkrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/remarkrpc"
	rolerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/rolerpc"
	tagrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/tagrpc"
	userrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/server/userrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
		blog.RegisterAuthRpcServer(grpcServer, authrpcServer.NewAuthRpcServer(ctx))
		blog.RegisterApiRpcServer(grpcServer, apirpcServer.NewApiRpcServer(ctx))
		blog.RegisterMenuRpcServer(grpcServer, menurpcServer.NewMenuRpcServer(ctx))
		blog.RegisterRoleRpcServer(grpcServer, rolerpcServer.NewRoleRpcServer(ctx))
		blog.RegisterUserRpcServer(grpcServer, userrpcServer.NewUserRpcServer(ctx))
		blog.RegisterConfigRpcServer(grpcServer, configrpcServer.NewConfigRpcServer(ctx))
		blog.RegisterArticleRpcServer(grpcServer, articlerpcServer.NewArticleRpcServer(ctx))
		blog.RegisterCategoryRpcServer(grpcServer, categoryrpcServer.NewCategoryRpcServer(ctx))
		blog.RegisterTagRpcServer(grpcServer, tagrpcServer.NewTagRpcServer(ctx))
		blog.RegisterFriendLinkRpcServer(grpcServer, friendlinkrpcServer.NewFriendLinkRpcServer(ctx))
		blog.RegisterRemarkRpcServer(grpcServer, remarkrpcServer.NewRemarkRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
