package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/interceptorx"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	apirpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/apirpc"
	articlerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/articlerpc"
	authrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/authrpc"
	blogrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/blogrpc"
	categoryrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/categoryrpc"
	chatrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/chatrpc"
	commentrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/commentrpc"
	configrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/configrpc"
	friendlinkrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/friendlinkrpc"
	logrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/logrpc"
	menurpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/menurpc"
	pagerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/pagerpc"
	photorpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/photorpc"
	remarkrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/remarkrpc"
	rolerpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/rolerpc"
	tagrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/tagrpc"
	talkrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/talkrpc"
	uploadrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/uploadrpc"
	userrpcServer "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/server/userrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	nacosIP        = flag.String("nacos-ip", "120.79.136.81", "Input Your Nacos IP")
	nacosPort      = flag.Int64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUserName  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosDataId    = flag.String("nacos-data-id", "rpc", "Input Your Nacos DataId")
	nacosGroup     = flag.String("nacos-group", "veweiyi.cn", "nacos group")
	nacosNameSpace = flag.String("nacos-namespace", "test", "Input Your Nacos NameSpaceId")
)

var configFile = flag.String("f", "", "the config file")

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Llongfile)
	var c config.Config
	if *configFile != "" {
		log.Println("load config from file: " + *configFile)
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
		blog.RegisterCommentRpcServer(grpcServer, commentrpcServer.NewCommentRpcServer(ctx))
		blog.RegisterPhotoRpcServer(grpcServer, photorpcServer.NewPhotoRpcServer(ctx))
		blog.RegisterPageRpcServer(grpcServer, pagerpcServer.NewPageRpcServer(ctx))
		blog.RegisterTalkRpcServer(grpcServer, talkrpcServer.NewTalkRpcServer(ctx))
		blog.RegisterLogRpcServer(grpcServer, logrpcServer.NewLogRpcServer(ctx))
		blog.RegisterChatRpcServer(grpcServer, chatrpcServer.NewChatRpcServer(ctx))
		blog.RegisterUploadRpcServer(grpcServer, uploadrpcServer.NewUploadRpcServer(ctx))
		blog.RegisterBlogRpcServer(grpcServer, blogrpcServer.NewBlogRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(interceptorx.ServerMetaInterceptor)
	s.AddUnaryInterceptors(interceptorx.ServerErrorInterceptor)
	s.AddUnaryInterceptors(interceptorx.ServerLogInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
