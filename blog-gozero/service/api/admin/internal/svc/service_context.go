package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/loads"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/rbacx"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tokenx"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/docs"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oss"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc    accountrpc.AccountRpc
	PermissionRpc permissionrpc.PermissionRpc
	ArticleRpc    articlerpc.ArticleRpc
	MessageRpc    messagerpc.MessageRpc
	PhotoRpc      photorpc.PhotoRpc
	ResourceRpc   resourcerpc.ResourceRpc
	TalkRpc       talkrpc.TalkRpc
	WebsiteRpc    websiterpc.WebsiteRpc
	ConfigRpc     configrpc.ConfigRpc
	SyslogRpc     syslogrpc.SyslogRpc

	Redis            *redis.Redis
	PermissionHolder rbacx.PermissionHolder
	TokenHolder      tokenx.TokenHolder
	Uploader         oss.OSS

	TimeToken    rest.Middleware
	JwtToken     rest.Middleware
	Permission   rest.Middleware
	OperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	var options []zrpc.ClientOption
	options = append(options)

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	uploader := oss.NewQiniu(c.UploadConfig)

	th := tokenx.NewJwtTokenHolder(c.Name, c.Name, rds)

	doc, err := loads.Analyzed(json.RawMessage(docs.Docs), "")
	if err != nil {
		panic(err)
	}

	accountRpc := accountrpc.NewAccountRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	permissionRpc := permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	articleRpc := articlerpc.NewArticleRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	messageRpc := messagerpc.NewMessageRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	photoRpc := photorpc.NewPhotoRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	resourceRpc := resourcerpc.NewResourceRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	talkRpc := talkrpc.NewTalkRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	websiteRpc := websiterpc.NewWebsiteRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	configRpc := configrpc.NewConfigRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))
	syslogRpc := syslogrpc.NewSyslogRpc(zrpc.MustNewClient(c.BlogRpcConf, options...))

	ph := rbacx.NewCasbinHolder(fmt.Sprintf("%s:%s", c.RedisConf.Host, c.RedisConf.Port), permissionRpc)
	err = ph.LoadPolicy()
	if err != nil {
		panic(err)
	}
	ph.StartAutoLoadPolicy()

	return &ServiceContext{
		Config:        c,
		AccountRpc:    accountRpc,
		PermissionRpc: permissionRpc,
		ArticleRpc:    articleRpc,
		MessageRpc:    messageRpc,
		PhotoRpc:      photoRpc,
		ResourceRpc:   resourceRpc,
		TalkRpc:       talkRpc,
		WebsiteRpc:    websiteRpc,
		ConfigRpc:     configRpc,
		SyslogRpc:     syslogRpc,

		Redis:            rds,
		Uploader:         uploader,
		TokenHolder:      th,
		PermissionHolder: ph,
		TimeToken:        middlewarex.NewTimeTokenMiddleware().Handle,
		JwtToken:         middlewarex.NewJwtTokenMiddleware(th).Handle,
		//Permission:       middlewarex.NewPermissionMiddleware(ph).Handle,
		Permission:   middlewarex.NewMiddleware().Handle, // 不使用接口权限控制
		OperationLog: middlewarex.NewOperationLogMiddleware(doc.Spec(), syslogRpc).Handle,
	}
}

func ConnectRedis(c config.RedisConf) (*redis.Redis, error) {
	address := c.Host + ":" + c.Port
	client, err := redis.NewRedis(redis.RedisConf{
		Host: address,
		Type: redis.NodeType,
		Pass: c.Password,
		Tls:  false,
	})

	if err != nil {
		return nil, fmt.Errorf("redis 连接失败: %v", err)
	}

	client.SetexCtx(context.Background(), fmt.Sprintf("redis:admin:%s", "PONG"), time.Now().String(), -1)
	return client, nil
}
