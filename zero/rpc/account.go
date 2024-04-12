package main

import (
	"flag"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/config"
	accountrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/server/accountrpc"
	apirpcServer "github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/server/apirpc"
	menurpcServer "github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/server/menurpc"
	rolerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/server/rolerpc"
	userServer "github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/server/userrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/account.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		account.RegisterAccountRpcServer(grpcServer, accountrpcServer.NewAccountRpcServer(ctx))
		account.RegisterApiRpcServer(grpcServer, apirpcServer.NewApiRpcServer(ctx))
		account.RegisterMenuRpcServer(grpcServer, menurpcServer.NewMenuRpcServer(ctx))
		account.RegisterRoleRpcServer(grpcServer, rolerpcServer.NewRoleRpcServer(ctx))
		account.RegisterUserRpcServer(grpcServer, userServer.NewUserRpcServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
