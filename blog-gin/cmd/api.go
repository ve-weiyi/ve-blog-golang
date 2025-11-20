/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/core"
)

type ApiCmd struct {
	cmd *cobra.Command

	configOption *core.ConfigOption
}

func NewApiCmd() *ApiCmd {
	apiCmd := &ApiCmd{
		configOption: &core.ConfigOption{
			ConfigMode:      "file",
			ConfigType:      "yaml",
			LocalPath:       "config.yaml",
			NacosHost:       "veweiyi.cn",
			NacosPort:       8848,
			NacosNamespace:  "dev",
			NacosDataID:     "ve-blog-golang",
			NacosGroup:      "blog",
			NacosRuntimeDir: "runtime/nacos",
			NacosUsername:   "nacos",
			NacosPassword:   "nacos",
		},
	}
	apiCmd.cmd = &cobra.Command{
		Use:   "api",
		Short: "启动接口服务",
		Long:  `启动接口服务`,
		Run: func(cmd *cobra.Command, args []string) {
			apiCmd.RunApi(cmd, args)
		},
	}

	apiCmd.init()
	return apiCmd
}

func (s *ApiCmd) init() {
	// 设置默认参数
	s.cmd.Flags().StringVarP(&s.configOption.ConfigMode, "config", "c", s.configOption.ConfigMode, "the way of read config file (file|nacos)")
	s.cmd.Flags().StringVarP(&s.configOption.LocalPath, "filepath", "f", s.configOption.LocalPath, "config file path (default is ./config.yaml)")

	s.cmd.Flags().StringVar(&s.configOption.NacosHost, "n-host", s.configOption.NacosHost, "the host for nacos")
	s.cmd.Flags().Uint64Var(&s.configOption.NacosPort, "n-port", s.configOption.NacosPort, "the port for nacos")
	s.cmd.Flags().StringVar(&s.configOption.NacosNamespace, "n-namespace", s.configOption.NacosNamespace, "the namespace for nacos")
	s.cmd.Flags().StringVar(&s.configOption.NacosDataID, "n-data-id", s.configOption.NacosDataID, "the DataId for nacos")
	s.cmd.Flags().StringVar(&s.configOption.NacosGroup, "n-group", s.configOption.NacosGroup, "the group for nacos")
	s.cmd.Flags().StringVar(&s.configOption.NacosUsername, "n-user", s.configOption.NacosUsername, "the user for nacos")
	s.cmd.Flags().StringVar(&s.configOption.NacosPassword, "n-password", s.configOption.NacosPassword, "the password for nacos")
}

func (s *ApiCmd) RunApi(cmd *cobra.Command, args []string) {
	c, err := core.InitConfig(s.configOption)
	if err != nil {
		fmt.Println("initialize config failed:", err)
		return
	}

	// 初始化配置文件
	core.RunWindowsServer(c)
}
