package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/interceptorx"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	accountrpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/accountrpc"
	articlerpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/articlerpc"
	configrpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/configrpc"
	messagerpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/messagerpc"
	permissionrpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/permissionrpc"
	photorpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/photorpc"
	resourcerpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/resourcerpc"
	syslogrpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/syslogrpc"
	talkrpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/talkrpc"
	websiterpcServer "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/server/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"
)

var (
	nacosIP        = flag.String("nacos-ip", "veweiyi.cn", "Input Your Nacos IP")
	nacosPort      = flag.Int64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUserName  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosNameSpace = flag.String("nacos-namespace", "test", "Input Your Nacos NameSpaceId")
	nacosGroup     = flag.String("nacos-group", "veweiyi.cn", "nacos group")
	nacosDataId    = flag.String("nacos-data-id", "rpc", "Input Your Nacos DataId")
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
		accountrpc.RegisterAccountRpcServer(grpcServer, accountrpcServer.NewAccountRpcServer(ctx))
		articlerpc.RegisterArticleRpcServer(grpcServer, articlerpcServer.NewArticleRpcServer(ctx))
		permissionrpc.RegisterPermissionRpcServer(grpcServer, permissionrpcServer.NewPermissionRpcServer(ctx))
		photorpc.RegisterPhotoRpcServer(grpcServer, photorpcServer.NewPhotoRpcServer(ctx))
		syslogrpc.RegisterSyslogRpcServer(grpcServer, syslogrpcServer.NewSyslogRpcServer(ctx))
		messagerpc.RegisterMessageRpcServer(grpcServer, messagerpcServer.NewMessageRpcServer(ctx))
		talkrpc.RegisterTalkRpcServer(grpcServer, talkrpcServer.NewTalkRpcServer(ctx))
		websiterpc.RegisterWebsiteRpcServer(grpcServer, websiterpcServer.NewWebsiteRpcServer(ctx))
		configrpc.RegisterConfigRpcServer(grpcServer, configrpcServer.NewConfigRpcServer(ctx))
		resourcerpc.RegisterResourceRpcServer(grpcServer, resourcerpcServer.NewResourceRpcServer(ctx))

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
