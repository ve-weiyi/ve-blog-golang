/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

type ServerCmd struct {
	cmd        *cobra.Command
	configFile string
}

func NewServerCmd() *ServerCmd {
	serverCmd := &ServerCmd{}
	serverCmd.cmd = &cobra.Command{
		Use:   "server",
		Short: "启动接口服务",
		Long:  `启动接口服务`,
		Run: func(cmd *cobra.Command, args []string) {
			serverCmd.OnInitialize()
		},
	}
	serverCmd.init()
	return serverCmd
}

func (s *ServerCmd) init() {
	// 设置默认参数
	s.cmd.PersistentFlags().StringVar(&s.configFile, "config", "config.yaml", "config file (default is $HOME/.config.yaml)")
}

func (s *ServerCmd) OnInitialize() {
	log.Println("let's go")

	// 初始化Viper
	initialize.Viper(s.configFile)
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

	// 文件上传组件
	initialize.Upload()

	initialize.OtherInit()

	// 程序结束前关闭数据库链接
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}

	initialize.RunWindowsServer()
}
