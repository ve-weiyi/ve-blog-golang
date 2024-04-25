package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
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

	PhotoRpcConf zrpc.RpcClientConf
	TalkRpcConf  zrpc.RpcClientConf
	LogRpcConf   zrpc.RpcClientConf
}
