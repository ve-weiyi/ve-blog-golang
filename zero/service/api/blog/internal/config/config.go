package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
)

type Config struct {
	rest.RestConf
	RedisConf    RedisConf
	UploadConfig *upload.UploadConfig

	BlogRpcConf zrpc.RpcClientConf
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}
