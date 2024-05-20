/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/nacos"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

type ApiCmd struct {
	cmd        *cobra.Command
	configFile string
	useNacos   bool
	nacosCfg   *nacos.NacosConfig
}

func NewApiCmd() *ApiCmd {
	serverCmd := &ApiCmd{}
	serverCmd.cmd = &cobra.Command{
		Use:   "api",
		Short: "启动接口服务",
		Long:  `启动接口服务`,
		Run: func(cmd *cobra.Command, args []string) {
			serverCmd.RunApi()
		},
	}
	serverCmd.cmd.PersistentPreRun = serverCmd.persistentPreRun
	serverCmd.init()
	return serverCmd
}

func (s *ApiCmd) init() {
	nacosCfg := s.GetDefaultNacosConfig()
	s.nacosCfg = nacosCfg
	// 设置默认参数
	s.cmd.PersistentFlags().StringVarP(&s.configFile, "config", "c", "config.yaml", "config file (default is $HOME/.config.yaml)")
	s.cmd.PersistentFlags().BoolVar(&s.useNacos, "use-nacos", false, "service conf content from nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.IP, "n-ip", nacosCfg.IP, "the ip for nacos")
	s.cmd.PersistentFlags().Uint64Var(&s.nacosCfg.Port, "n-port", nacosCfg.Port, "the port for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.UserName, "n-user", nacosCfg.UserName, "the user for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.Password, "n-password", nacosCfg.Password, "the password for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.DataID, "n-data-id", nacosCfg.DataID, "the DataId for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.Group, "n-group", nacosCfg.Group, "the group for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.NameSpaceID, "n-ns", nacosCfg.NameSpaceID, "the namespace for nacos")
}

func (s *ApiCmd) GetDefaultNacosConfig() *nacos.NacosConfig {
	return &nacos.NacosConfig{
		IP:          "120.79.136.81",
		Port:        8848,
		UserName:    "nacos",
		Password:    "nacos",
		NameSpaceID: "dev",
		Group:       "veweiyi.cn",
		DataID:      "ve-blog-golang",
		RuntimeDir:  "runtime/nacos",
		LogLevel:    "warn",
		Timeout:     5000,
	}
}

func (s *ApiCmd) persistentPreRun(cmd *cobra.Command, args []string) {

}

func (s *ApiCmd) RunApi() {
	if s.useNacos {
		log.Println("读取配置文件...使用nacos")
		nc := nacos.New(s.nacosCfg)

		// 获取配置文件
		content, err := nc.GetConfig()
		if err != nil {
			panic("nacos config read failed " + err.Error())
		}

		initialize.InitConfigByContent(content)
		s.OnInitialize()

		// 监听配置文件变化
		err = nc.AddListener(func(content string) error {
			log.Println("更新配置文件...")
			initialize.InitConfigByContent(content)
			s.OnInitialize()
			return nil
		})
		if err != nil {
			panic("nacos config listener failed " + err.Error())
		}
	} else {
		log.Println("读取配置文件..使用文件路径")
		// 初始化Viper
		initialize.InitConfigByFile(s.configFile)
		s.OnInitialize()
	}
}

func (s *ApiCmd) OnInitialize() {
	log.Println("let's go")

	// 初始化zap日志库
	initialize.Zap()
	// 初始化gorm数据库
	initialize.Gorm()
	// 初始化redis服务
	initialize.Redis()
	// 初始化jwt
	initialize.JwtToken()
	// 初始化rbac角色访问控制
	initialize.RBAC()

	// 文件上传组件
	initialize.Upload()

	// 消息队列
	initialize.RabbitMq()

	initialize.OtherInit()

	// 程序结束前关闭数据库链接
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}

	initialize.RunWindowsServer()
}
