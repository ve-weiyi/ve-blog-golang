package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/upload"
)

type Config struct {
	rest.RestConf
	UploadConfig *upload.UploadConfig

	AccountRpcConf zrpc.RpcClientConf
	ApiRpcConf     zrpc.RpcClientConf
	MenuRpcConf    zrpc.RpcClientConf
	RoleRpcConf    zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf

	ConfigRpcConf   zrpc.RpcClientConf
	ArticleRpcConf  zrpc.RpcClientConf
	CategoryRpcConf zrpc.RpcClientConf
	TagRpcConf      zrpc.RpcClientConf

	FriendLinkRpcConf zrpc.RpcClientConf
	RemarkRpcConf     zrpc.RpcClientConf
	CommentRpcConf    zrpc.RpcClientConf

	PhotoRpcConf zrpc.RpcClientConf
	TalkRpcConf  zrpc.RpcClientConf
	LogRpcConf   zrpc.RpcClientConf
	ChatRpcConf  zrpc.RpcClientConf

	UploadRpcConf zrpc.RpcClientConf
}
