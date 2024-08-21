package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/middlewarex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/apirpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/authrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/commentrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/friendlinkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/logrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/menurpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/pagerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/uploadrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/userrpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/config"
)

type ServiceContext struct {
	Config config.Config

	AuthRpc authrpc.AuthRpc
	ApiRpc  apirpc.ApiRpc
	MenuRpc menurpc.MenuRpc
	RoleRpc rolerpc.RoleRpc
	UserRpc userrpc.UserRpc
	BlogRpc blogrpc.BlogRpc

	ArticleRpc    articlerpc.ArticleRpc
	RemarkRpc     remarkrpc.RemarkRpc
	CommentRpc    commentrpc.CommentRpc
	PhotoRpc      photorpc.PhotoRpc
	TalkRpc       talkrpc.TalkRpc
	PageRpc       pagerpc.PageRpc
	FriendLinkRpc friendlinkrpc.FriendLinkRpc

	ConfigRpc configrpc.ConfigRpc
	LogRpc    logrpc.LogRpc
	ChatRpc   chatrpc.ChatRpc
	UploadRpc uploadrpc.UploadRpc

	Uploader upload.Uploader
	Token    *jtoken.JwtInstance

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
		Config:  c,
		Token:   jwt,
		AuthRpc: authrpc.NewAuthRpc(zrpc.MustNewClient(c.AccountRpcConf, options...)),
		ApiRpc:  apirpc.NewApiRpc(zrpc.MustNewClient(c.ApiRpcConf, options...)),
		MenuRpc: menurpc.NewMenuRpc(zrpc.MustNewClient(c.MenuRpcConf, options...)),
		RoleRpc: rolerpc.NewRoleRpc(zrpc.MustNewClient(c.RoleRpcConf, options...)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf, options...)),
		BlogRpc: blogrpc.NewBlogRpc(zrpc.MustNewClient(c.BlogRpcConf, options...)),

		ConfigRpc:     configrpc.NewConfigRpc(zrpc.MustNewClient(c.ConfigRpcConf, options...)),
		ArticleRpc:    articlerpc.NewArticleRpc(zrpc.MustNewClient(c.ArticleRpcConf, options...)),
		FriendLinkRpc: friendlinkrpc.NewFriendLinkRpc(zrpc.MustNewClient(c.FriendLinkRpcConf, options...)),
		RemarkRpc:     remarkrpc.NewRemarkRpc(zrpc.MustNewClient(c.RemarkRpcConf, options...)),
		CommentRpc:    commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf, options...)),
		PhotoRpc:      photorpc.NewPhotoRpc(zrpc.MustNewClient(c.PhotoRpcConf, options...)),
		TalkRpc:       talkrpc.NewTalkRpc(zrpc.MustNewClient(c.TalkRpcConf, options...)),
		PageRpc:       pagerpc.NewPageRpc(zrpc.MustNewClient(c.PageRpcConf, options...)),

		LogRpc:    logrpc.NewLogRpc(zrpc.MustNewClient(c.LogRpcConf, options...)),
		ChatRpc:   chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf, options...)),
		UploadRpc: uploadrpc.NewUploadRpc(zrpc.MustNewClient(c.UploadRpcConf, options...)),
		Uploader:  upload.NewQiniu(c.UploadConfig),
		JwtToken:  middlewarex.NewJwtTokenMiddleware(jwt, rds).Handle,
		SignToken: middlewarex.NewSignTokenMiddleware().Handle,
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

	client.SetexCtx(context.Background(), fmt.Sprintf("redis:%s", "pong"), time.Now().String(), -1)
	return client, nil
}
