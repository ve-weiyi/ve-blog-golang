package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
)

type Config struct {
	rest.RestConf

	BlogRpcConf    zrpc.RpcClientConf
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
	PhotoRpcConf      zrpc.RpcClientConf
	TalkRpcConf       zrpc.RpcClientConf
	PageRpcConf       zrpc.RpcClientConf

	LogRpcConf    zrpc.RpcClientConf
	ChatRpcConf   zrpc.RpcClientConf
	UploadRpcConf zrpc.RpcClientConf

	UploadConfig *upload.UploadConfig
	RedisConf    RedisConf
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}
