package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/config"
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
	Token  *jjwt.JwtToken

	AuthRpc authrpc.AuthRpc
	ApiRpc  apirpc.ApiRpc
	MenuRpc menurpc.MenuRpc
	RoleRpc rolerpc.RoleRpc
	UserRpc userrpc.UserRpc

	ConfigRpc   configrpc.ConfigRpc
	ArticleRpc  articlerpc.ArticleRpc
	CategoryRpc categoryrpc.CategoryRpc
	TagRpc      tagrpc.TagRpc

	FriendLinkRpc friendlinkrpc.FriendLinkRpc
	RemarkRpc     remarkrpc.RemarkRpc
	CommentRpc    commentrpc.CommentRpc
	PhotoRpc      photorpc.PhotoRpc
	TalkRpc       talkrpc.TalkRpc
	PageRpc       pagerpc.PageRpc

	LogRpc    logrpc.LogRpc
	ChatRpc   chatrpc.ChatRpc
	UploadRpc uploadrpc.UploadRpc

	Uploader upload.Uploader
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Token:   jjwt.NewJwtToken([]byte("ve-weiyi")),
		AuthRpc: authrpc.NewAuthRpc(zrpc.MustNewClient(c.AccountRpcConf)),
		ApiRpc:  apirpc.NewApiRpc(zrpc.MustNewClient(c.ApiRpcConf)),
		MenuRpc: menurpc.NewMenuRpc(zrpc.MustNewClient(c.MenuRpcConf)),
		RoleRpc: rolerpc.NewRoleRpc(zrpc.MustNewClient(c.RoleRpcConf)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),

		ConfigRpc:     configrpc.NewConfigRpc(zrpc.MustNewClient(c.ConfigRpcConf)),
		ArticleRpc:    articlerpc.NewArticleRpc(zrpc.MustNewClient(c.ArticleRpcConf)),
		CategoryRpc:   categoryrpc.NewCategoryRpc(zrpc.MustNewClient(c.CategoryRpcConf)),
		TagRpc:        tagrpc.NewTagRpc(zrpc.MustNewClient(c.TagRpcConf)),
		FriendLinkRpc: friendlinkrpc.NewFriendLinkRpc(zrpc.MustNewClient(c.FriendLinkRpcConf)),
		RemarkRpc:     remarkrpc.NewRemarkRpc(zrpc.MustNewClient(c.RemarkRpcConf)),
		CommentRpc:    commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
		PhotoRpc:      photorpc.NewPhotoRpc(zrpc.MustNewClient(c.PhotoRpcConf)),
		TalkRpc:       talkrpc.NewTalkRpc(zrpc.MustNewClient(c.TalkRpcConf)),
		PageRpc:       pagerpc.NewPageRpc(zrpc.MustNewClient(c.PageRpcConf)),

		LogRpc:    logrpc.NewLogRpc(zrpc.MustNewClient(c.LogRpcConf)),
		ChatRpc:   chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
		UploadRpc: uploadrpc.NewUploadRpc(zrpc.MustNewClient(c.UploadRpcConf)),
		Uploader:  upload.NewQiniu(c.UploadConfig),
	}
}
