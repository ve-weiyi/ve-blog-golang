package config

import (
	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/database/orm"
)

type Config struct {
	JWT      properties.JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      properties.Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis    properties.Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	RabbitMQ properties.RabbitMQ `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	Email    properties.Email    `mapstructure:"email" json:"email" yaml:"email"`
	System   properties.System   `mapstructure:"system" json:"system" yaml:"system"`
	Captcha  properties.Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Oauth    properties.Oauth    `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
	// gorm
	Mysql  orm.Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []orm.SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Upload  properties.Upload  `mapstructure:"upload" json:"upload" yaml:"upload"`
	ChatGPT properties.ChatGPT `mapstructure:"chatgpt" json:"chatgpt" yaml:"chatgpt"`
}
