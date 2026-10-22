package config

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/vkit/adapter/storagex"
)

type Config struct {
	rest.RestConf
	RedisConf     RedisConf
	StorageConfig storagex.StorageConfig
	TokenConf     TokenConf

	AppRpcConf zrpc.RpcClientConf
}

type TokenConf struct {
	AccessTokenTTL     time.Duration `json:"access_token_ttl,optional,default=2h"`
	RefreshTokenTTL    time.Duration `json:"refresh_token_ttl,optional,default=168h"` // idle
	RefreshAbsoluteTTL time.Duration `json:"refresh_absolute_ttl,optional,default=720h"`
}

// redis缓存配置
type RedisConf struct {
	Host     string `json:"host,default=127.0.0.1"`
	Port     string `json:"port,default=6379"`
	Password string `json:"password,default=''"`
	DB       int    `json:"db,default=0"`
}
