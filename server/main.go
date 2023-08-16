package main

import (
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// @title						Swagger Example API
// @version					1.0
// @description				This is a sample server celler server.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:9999/
// @BasePath					/api/v1
// @securityDefinitions.basic	BasicAuth
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-token
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log.Println("let's go")
	// 初始化Viper
	initialize.Viper()
	// 初始化zap日志库
	initialize.Zap()
	// 初始化gorm数据库
	initialize.Gorm()
	// 初始化redis服务
	initialize.Redis()
	// 初始化jwt
	initialize.JwtToken()
	// 初始化rbac角色访问控制
	//initialize.RBAC()

	initialize.OtherInit()

	// 创建协程运行rabbitmq订阅消息
	go initialize.RabbitMq()
	// 程序结束前关闭数据库链接
	if global.DB != nil {
		initialize.RegisterTables(global.DB) // 初始化表
		db, _ := global.DB.DB()
		defer db.Close()
	}

	initialize.RunWindowsServer()
}
