/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/initialize"
)

type RabbitmqCmd struct {
	cmd      *cobra.Command
	filepath string
}

func NewRabbitmqCmd() *RabbitmqCmd {
	rabbitmqCmd := &RabbitmqCmd{}
	rabbitmqCmd.cmd = &cobra.Command{
		Use:   "rabbitmq",
		Short: "运行rabbitmq服务",
		Long:  `运行rabbitmq服务，订阅消息,发送邮件`,
		Run: func(cmd *cobra.Command, args []string) {
			rabbitmqCmd.RunRabbitmq()
		},
	}
	rabbitmqCmd.cmd.PersistentPreRun = rabbitmqCmd.persistentPreRun
	rabbitmqCmd.init()
	return rabbitmqCmd
}

func (s *RabbitmqCmd) init() {
	// 设置默认参数
	s.cmd.PersistentFlags().StringVarP(&s.filepath, "filepath", "f", "config.yaml", "config file path (default is ./config.yaml)")
}

func (s *RabbitmqCmd) persistentPreRun(cmd *cobra.Command, args []string) {

}

func (s *RabbitmqCmd) RunRabbitmq() {
	var c config.Config
	// 初始化Viper
	v := viper.New()
	v.SetConfigFile(s.filepath)
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

	log.Println("rabbitmq服务启动成功", jsonconv.ObjectToJsonIndent(c))
	log.Println("rabbitmq服务启动成功")

	initialize.SubscribeMessage(c)
}
