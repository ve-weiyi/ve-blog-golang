package config

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/vkit/adapter/storagex"
)

type Config struct {
	rest.RestConf

	StorageConfig storagex.StorageConfig

	RedisConf RedisConf

	TokenConf TokenConf

	AppRpcConf zrpc.RpcClientConf
}

type TokenConf struct {
	Secret string `json:"secret"` // HS256 签名密钥，至少 32 字节

	AccessTokenTTL     time.Duration `json:"access_token_ttl,optional,default=2h"`
	RefreshTokenTTL    time.Duration `json:"refresh_token_ttl,optional,default=168h"` // idle
	RefreshAbsoluteTTL time.Duration `json:"refresh_absolute_ttl,optional,default=720h"`
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}
