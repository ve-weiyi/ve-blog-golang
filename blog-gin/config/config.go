package config

type Config struct {
	System   SystemConf   `json:"system"`
	JWT      JWTConf      `json:"jwt"`
	Zap      Zap          `json:"zap"`
	Mysql    MysqlConf    `json:"mysql"`
	Redis    RedisConf    `json:"redis"`
	RabbitMQ RabbitMQConf `json:"rabbitmq"`
}

// 系统配置
type SystemConf struct {
	Port        int    `json:"port" example:"8080"`              // 运行端口
	Env         string `json:"env" example:"dev"`                // 运行环境 dev、test、prod
	Version     string `json:"version" example:"v1.0.0"`         // 程序版本
	RuntimePath string `json:"runtime-path" example:"./runtime"` // 运行时目录
}

// jwt鉴权
type JWTConf struct {
	Secret      string `json:"secret" example:"your-secret-key"` // 加密密钥
	ExpiresTime string `json:"expires-time" example:"7d"`        // 过期时间
}

// mysql数据库
type MysqlConf struct {
	Host     string `json:"host" example:"localhost"`
	Port     string `json:"port" example:"3306"`
	Username string `json:"username" example:"root"`
	Password string `json:"password" example:"123456"`
	Dbname   string `json:"dbname" example:"blog"`
	Config   string `json:"config" example:"charset=utf8mb4&parseTime=True&loc=Local"`
}

// redis缓存
type RedisConf struct {
	Host     string `json:"host" example:"localhost"` // 服务器地址
	Port     string `json:"port" example:"6379"`      // 端口
	Password string `json:"password" example:""`      // 密码
	DB       int    `json:"db" example:"0"`           // redis的哪个数据库
}

// rabbitmq消息队列
type RabbitMQConf struct {
	Host     string `json:"host" example:"localhost"`
	Port     string `json:"port" example:"5672"`
	Username string `json:"username" example:"guest"`
	Password string `json:"password" example:"guest"`
}

type Zap struct {
	ServiceName string `json:"service-name,optional" example:"ve-blog-gin"`                   // 服务名称
	Mode        string `json:"mode,default=console,options=[console,file]" example:"console"` // 日志模式 console-控制台输出，file-文件输出
	Encoding    string `json:"encoding,default=json,options=[json,plain]" example:"json"`     // 输出格式 json-JSON格式，plain-纯文本格式
	TimeFormat  string `json:"time-format,optional" example:"2006-01-02T15:04:05.000Z07:00"`  // 时间格式
	Path        string `json:"path,default=logs" example:"logs"`
	Level       string `json:"level,default=info,options=[debug,info,error,severe]" example:"info"`

	KeepDays int `json:"keep-days,optional" example:"7"`
}
