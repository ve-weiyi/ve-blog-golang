package properties

import "fmt"

type RabbitMQ struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // RabbitMQ服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // RabbitMQ端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // RabbitMQ用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // RabbitMQ密码
}

func (r *RabbitMQ) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.Username, r.Password, r.Host, r.Port)
}
