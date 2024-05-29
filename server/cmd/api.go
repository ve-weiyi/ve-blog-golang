/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/core"
)

type ApiCmd struct {
	cmd        *cobra.Command
	configFile string
	nacosCfg   *nacos.NacosConfig
}

func NewApiCmd() *ApiCmd {
	apiCmd := &ApiCmd{}
	apiCmd.cmd = &cobra.Command{
		Use:   "api",
		Short: "启动接口服务",
		Long:  `启动接口服务`,
		Run: func(cmd *cobra.Command, args []string) {
			apiCmd.RunApi()
		},
	}
	apiCmd.cmd.PersistentPreRun = apiCmd.persistentPreRun
	apiCmd.init()
	return apiCmd
}

func (s *ApiCmd) init() {
	nacosCfg := s.GetDefaultNacosConfig()
	s.nacosCfg = nacosCfg
	// 设置默认参数
	s.cmd.PersistentFlags().StringVarP(&s.configFile, "config", "c", "config.yaml", "config file (default is $HOME/.config.yaml)")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.IP, "n-ip", nacosCfg.IP, "the ip for nacos")
	s.cmd.PersistentFlags().Uint64Var(&s.nacosCfg.Port, "n-port", nacosCfg.Port, "the port for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.UserName, "n-user", nacosCfg.UserName, "the user for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.Password, "n-password", nacosCfg.Password, "the password for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.DataId, "n-data-id", nacosCfg.DataId, "the DataId for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.Group, "n-group", nacosCfg.Group, "the group for nacos")
	s.cmd.PersistentFlags().StringVar(&s.nacosCfg.NameSpaceId, "n-namespace", nacosCfg.NameSpaceId, "the namespace for nacos")
}

func (s *ApiCmd) GetDefaultNacosConfig() *nacos.NacosConfig {
	return &nacos.NacosConfig{
		IP:          "120.79.136.81",
		Port:        8848,
		UserName:    "nacos",
		Password:    "nacos",
		NameSpaceId: "dev",
		Group:       "veweiyi.cn",
		DataId:      "ve-blog-golang",
		RuntimeDir:  "runtime/nacos",
		LogLevel:    "warn",
		Timeout:     5000,
	}
}

func (s *ApiCmd) persistentPreRun(cmd *cobra.Command, args []string) {

}

func (s *ApiCmd) RunApi() {
	var c config.Config

	if s.configFile != "" {
		log.Println("读取配置文件..使用文件路径")
		// 初始化Viper
		v := viper.New()
		v.SetConfigFile(s.configFile)
		v.SetConfigType("yaml")
		// 读取配置文件
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s \n", err))
		}
		// 修改解析的tag（默认是mapstructure）
		withJsonTag := func(c *mapstructure.DecoderConfig) {
			c.TagName = "json"
		}
		// 解析配置文件
		if err = v.Unmarshal(&c, withJsonTag); err != nil {
			panic(err)
		}

		// 监听配置文件变化
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			log.Println("config file changed:", e.Name)
			if err = v.Unmarshal(&c, withJsonTag); err != nil {
				log.Println(err)
			}
		})

	} else {
		log.Println("读取配置文件...使用nacos")
		// 初始化Nacos
		nc := nacos.New(s.nacosCfg)

		// 读取配置文件
		content, err := nc.GetConfig()
		if err != nil {
			panic("nacos config read failed " + err.Error())
		}
		// 解析配置文件
		if err = json.Unmarshal([]byte(content), &c); err != nil {
			panic(err)
		}

		// 监听配置文件变化
		err = nc.AddListener(func(content string) error {
			log.Println("更新配置文件...")
			// 解析配置文件
			if err = json.Unmarshal([]byte(content), &c); err != nil {
				panic(err)
			}
			return nil
		})
		if err != nil {
			panic("nacos config listener failed " + err.Error())
		}
	}

	log.Println("let's go")
	// 初始化配置文件
	core.RunWindowsServer(&c)
}
