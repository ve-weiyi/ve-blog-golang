package properties

type Email struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 发件人  你自己要发邮件的邮箱
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Deliver  string `mapstructure:"deliver" json:"deliver" yaml:"deliver"`    // 抄送邮箱:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 发件人昵称
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // 是否使用 SSL/TLS
}
