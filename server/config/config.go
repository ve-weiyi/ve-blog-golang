package config

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
)

type Config struct {
	System   SystemConf           `json:"system" yaml:"system"`
	JWT      JWTConf              `json:"jwt" yaml:"jwt"`
	Zap      Zap                  `json:"zap" yaml:"zap"`
	Mysql    MysqlConf            `json:"mysql" yaml:"mysql"`
	Redis    RedisConf            `json:"redis" yaml:"redis"`
	RabbitMQ RabbitMQConf         `json:"rabbitmq" yaml:"rabbitmq"`
	Email    EmailConf            `json:"email" yaml:"email"`
	Oauth    map[string]OauthConf `json:"oauth" yaml:"oauth"`
	Upload   UploadConf           `json:"upload" yaml:"upload"`
	Captcha  CaptchaConf          `json:"captcha" yaml:"captcha"`
	ChatGPT  ChatGPTConf          `json:"chatgpt" yaml:"chatgpt"`
}

// 系统配置
type SystemConf struct {
	Version      string `json:"version" yaml:"version"`             // 程序版本
	Env          string `json:"env" yaml:"env"`                     // 运行环境
	Port         int    `json:"port" yaml:"port"`                   // 运行端口
	RouterPrefix string `json:"router-prefix" yaml:"router-prefix"` // 路由前缀
	RuntimePath  string `json:"runtime-path" yaml:"runtime-path"`   // 运行时目录
}

// jwt鉴权
type JWTConf struct {
	SigningKey  string `json:"signing-key" yaml:"signing-key"`   // jwt签名
	ExpiresTime string `json:"expires-time" yaml:"expires-time"` // 过期时间
	Issuer      string `json:"issuer" yaml:"issuer"`             // 签发者
	Type        string `json:"type" yaml:"type"`                 // 类型 例如：Bearer
}

// mysql数据库
type MysqlConf struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Dbname   string `json:"dbname" yaml:"dbname"`
	Config   string `json:"config" yaml:"config"`
}

// redis缓存
type RedisConf struct {
	Host     string `json:"host" yaml:"host"` // 服务器地址:端口
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"` // 密码
	DB       int    `json:"db" yaml:"db"`             // redis的哪个数据库
}

// rabbitmq消息队列
type RabbitMQConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// email邮件发送服务
type EmailConf struct {
	Host     string   `json:"host" yaml:"host"`         // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Port     int      `json:"port" yaml:"port"`         // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	Username string   `json:"username" yaml:"username"` // 发件人  你自己要发邮件的邮箱
	Password string   `json:"password" yaml:"password"` // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Deliver  []string `json:"deliver" yaml:"deliver"`   // 抄送邮箱:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	Nickname string   `json:"nickname" yaml:"nickname"` // 发件人昵称
	IsSSL    bool     `json:"is-ssl" yaml:"is-ssl"`     // 是否使用 SSL/TLS
}

// oauth配置
type OauthConf struct {
	ClientId     string `json:"client-id"`
	ClientSecret string `json:"client-secret"`
	RedirectUri  string `json:"redirect-uri"`
}

type CaptchaConf struct {
	KeyLong            int `json:"key-long" yaml:"key-long"`                         // 验证码长度
	ImgWidth           int `json:"img-width" yaml:"img-width"`                       // 验证码宽度
	ImgHeight          int `json:"img-height" yaml:"img-height"`                     // 验证码高度
	OpenCaptcha        int `json:"open-captcha" yaml:"open-captcha"`                 // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}

type ChatGPTConf struct {
	ApiHost string `json:"api-host" yaml:"api-host"` // host
	ApiKey  string `json:"api-key" yaml:"api-key"`   // 秘钥
	Model   string `json:"model" yaml:"model"`       // 模型
}

type UploadConf struct {
	Mode   string               `json:"mode"`
	Local  *upload.UploadConfig `json:"local"`
	Aliyun *upload.UploadConfig `json:"aliyun"`
	Qiniu  *upload.UploadConfig `json:"qiniu"`
}

type Zap struct {
	Mode           string `json:"mode" yaml:"mode"`                       // 模式
	Format         string `json:"format" yaml:"format"`                   // 输出
	Level          string `json:"level" yaml:"level"`                     // 级别
	Prefix         string `json:"prefix" yaml:"prefix"`                   // 日志前缀
	EncodeCaller   string `json:"encode-caller" yaml:"encode-caller"`     // 编码调用者
	EncodeColorful bool   `json:"encode-colorful" yaml:"encode-colorful"` // 编码调用者

	Filename string `json:"filename" yaml:"filename"`   // 日志文件名称
	CacheDir string `json:"cache-dir" yaml:"cache-dir"` // 日志文件夹
	MaxAge   int    `json:"max-age" yaml:"max-age"`     // 日志留存时间
}
