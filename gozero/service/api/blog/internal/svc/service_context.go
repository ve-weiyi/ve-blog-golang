package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tokenx"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/config"
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
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc    accountrpc.AccountRpc
	PermissionRpc permissionrpc.PermissionRpc
	ArticleRpc    articlerpc.ArticleRpc
	MessageRpc    messagerpc.MessageRpc
	PhotoRpc      photorpc.PhotoRpc
	TalkRpc       talkrpc.TalkRpc
	SyslogRpc     syslogrpc.SyslogRpc
	WebsiteRpc    websiterpc.WebsiteRpc
	ConfigRpc     configrpc.ConfigRpc
	ResourceRpc   resourcerpc.ResourceRpc

	Redis            *redis.Redis
	TokenHolder      tokenx.TokenHolder
	Uploader         oss.OSS
	WebsocketManager ws.ClientManager

	TimeToken rest.Middleware
	SignToken rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	var options []zrpc.ClientOption
	options = append(options)

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	th := tokenx.NewSignTokenHolder(c.Name, c.Name, rds)

	return &ServiceContext{
		Config:        c,
		AccountRpc:    accountrpc.NewAccountRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		PermissionRpc: permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		ArticleRpc:    articlerpc.NewArticleRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		MessageRpc:    messagerpc.NewMessageRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		PhotoRpc:      photorpc.NewPhotoRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		TalkRpc:       talkrpc.NewTalkRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		SyslogRpc:     syslogrpc.NewSyslogRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		WebsiteRpc:    websiterpc.NewWebsiteRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		ConfigRpc:     configrpc.NewConfigRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		ResourceRpc:   resourcerpc.NewResourceRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),

		Uploader:         oss.NewQiniu(c.UploadConfig),
		Redis:            rds,
		TokenHolder:      th,
		WebsocketManager: ws.NewDefaultClientManager(),

		TimeToken: middlewarex.NewTimeTokenMiddleware().Handle,
		SignToken: middlewarex.NewSignTokenMiddleware(th).Handle,
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

	client.SetexCtx(context.Background(), fmt.Sprintf("redis:blog:%s", "PONG"), time.Now().String(), -1)
	return client, nil
}
