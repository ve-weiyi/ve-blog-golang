package config

type Config struct {
	System System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  MysqlConf `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  RedisConf `mapstructure:"redis" json:"redis" yaml:"redis"`
	// 消息队列、邮件发送服务
	RabbitMQ RabbitMQConf `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	Email    EmailConf    `mapstructure:"email" json:"email" yaml:"email"`
	// 鉴权
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Oauth Oauth `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 文件上传
	Upload  Upload  `mapstructure:"upload" json:"upload" yaml:"upload"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	ChatGPT ChatGPT `mapstructure:"chatgpt" json:"chatgpt" yaml:"chatgpt"`
}

// mysql数据库配置
type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Config   string `json:"config"`
}

// redis缓存配置
type RedisConf struct {
	DB       int    `json:"db" yaml:"db"`     // redis的哪个数据库
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
}

// rabbitmq配置
type RabbitMQConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 邮件配置
type EmailConf struct {
	Host     string   `json:"host" yaml:"host"`         // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Port     int      `json:"port" yaml:"port"`         // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	Username string   `json:"username" yaml:"username"` // 发件人  你自己要发邮件的邮箱
	Password string   `json:"password" yaml:"password"` // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Deliver  []string `json:"deliver" yaml:"deliver"`   // 抄送邮箱:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	Nickname string   `json:"nickname" yaml:"nickname"` // 发件人昵称
	IsSSL    bool     `json:"isSSL" yaml:"is-ssl"`      // 是否使用 SSL/TLS
}

// oauth配置
type OauthConf struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
}

type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // 验证码长度
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // 验证码宽度
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // 验证码高度
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}

type ChatGPT struct {
	ApiHost string `mapstructure:"api-host" json:"api-host" yaml:"api-host"` // host
	ApiKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`    // 秘钥
	Model   string `mapstructure:"model" json:"model" yaml:"model"`          // 模型
}
