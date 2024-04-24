package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
	RedisConf RedisConf
}

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Config   string `json:"config"`
}

type RedisConf struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`       // redis的哪个数据库
	Host     string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
