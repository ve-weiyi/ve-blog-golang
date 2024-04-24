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

	ConfigRpcConf  zrpc.RpcClientConf
	ArticleRpcConf zrpc.RpcClientConf
}
