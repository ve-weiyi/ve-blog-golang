package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jtoken"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/config"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/middleware"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/apirpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/authrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/categoryrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/commentrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/friendlinkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/logrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/menurpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/pagerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/rolerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/tagrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/uploadrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/userrpc"
)

type ServiceContext struct {
	Config config.Config

	AuthRpc authrpc.AuthRpc
	ApiRpc  apirpc.ApiRpc
	MenuRpc menurpc.MenuRpc
	RoleRpc rolerpc.RoleRpc
	UserRpc userrpc.UserRpc

	ArticleRpc    articlerpc.ArticleRpc
	CategoryRpc   categoryrpc.CategoryRpc
	TagRpc        tagrpc.TagRpc
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

	return &ServiceContext{
		Config:  c,
		Token:   jwt,
		AuthRpc: authrpc.NewAuthRpc(zrpc.MustNewClient(c.AccountRpcConf, options...)),
		ApiRpc:  apirpc.NewApiRpc(zrpc.MustNewClient(c.ApiRpcConf, options...)),
		MenuRpc: menurpc.NewMenuRpc(zrpc.MustNewClient(c.MenuRpcConf, options...)),
		RoleRpc: rolerpc.NewRoleRpc(zrpc.MustNewClient(c.RoleRpcConf, options...)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf, options...)),

		ConfigRpc:     configrpc.NewConfigRpc(zrpc.MustNewClient(c.ConfigRpcConf, options...)),
		ArticleRpc:    articlerpc.NewArticleRpc(zrpc.MustNewClient(c.ArticleRpcConf, options...)),
		CategoryRpc:   categoryrpc.NewCategoryRpc(zrpc.MustNewClient(c.CategoryRpcConf, options...)),
		TagRpc:        tagrpc.NewTagRpc(zrpc.MustNewClient(c.TagRpcConf, options...)),
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
		JwtToken:  middleware.NewJwtTokenMiddleware(jwt, authrpc.NewAuthRpc(zrpc.MustNewClient(c.AccountRpcConf, options...))).Handle,
		SignToken: middleware.NewSignTokenMiddleware().Handle,
	}
}
