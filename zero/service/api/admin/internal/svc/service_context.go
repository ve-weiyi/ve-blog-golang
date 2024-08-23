package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/messagerpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/friendrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/websiterpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/config"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc    accountrpc.AccountRpc
	PermissionRpc permissionrpc.PermissionRpc
	ArticleRpc    articlerpc.ArticleRpc
	MessageRpc    messagerpc.MessageRpc
	PhotoRpc      photorpc.PhotoRpc
	WebsiteRpc    websiterpc.WebsiteRpc
	TalkRpc       talkrpc.TalkRpc
	FriendRpc     friendrpc.FriendRpc

	Uploader upload.Uploader
	Token    *jtoken.JwtInstance
	Redis    *redis.Redis

	JwtToken  rest.Middleware
	SignToken rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	var options []zrpc.ClientOption
	options = append(options)

	jwt := jtoken.NewJWTInstance([]byte(c.Name))

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		AccountRpc:    accountrpc.NewAccountRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		PermissionRpc: permissionrpc.NewPermissionRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		ArticleRpc:    articlerpc.NewArticleRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		MessageRpc:    messagerpc.NewMessageRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		PhotoRpc:      photorpc.NewPhotoRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		WebsiteRpc:    websiterpc.NewWebsiteRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		TalkRpc:       talkrpc.NewTalkRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		FriendRpc:     friendrpc.NewFriendRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),
		Uploader:      upload.NewQiniu(c.UploadConfig),
		Token:         jwt,
		Redis:         rds,
		JwtToken:      middlewarex.NewJwtTokenMiddleware(jwt, rds).Handle,
		SignToken:     middlewarex.NewSignTokenMiddleware().Handle,
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
