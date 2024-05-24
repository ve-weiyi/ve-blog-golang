package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlConf    MysqlConf
	RedisConf    RedisConf
	RabbitMQConf RabbitMQConf
	EmailConf    EmailConf
	OauthConf    map[string]AuthConf
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

type RabbitMQConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmailConf struct {
	Host     string   `json:"host"`     // 服务器地址
	Port     int      `json:"port"`     // 端口
	Username string   `json:"username"` // 发件人
	Password string   `json:"password"` // 密钥
	Nickname string   `json:"nickname"` // 发件人昵称
	Deliver  []string `json:"deliver"`  // 抄送邮箱:多个以英文逗号分隔
	IsSSL    bool     `json:"is_ssl"`   // 是否使用 SSL/TLS
}

// 基本配置
type AuthConf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
}
