package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`       // redis的哪个数据库
	Host     string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
