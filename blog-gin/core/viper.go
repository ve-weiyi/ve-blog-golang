package core

import (
	"github.com/mitchellh/mapstructure"

	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/config"
)

func Viper(filename string) *config.Config {
	var c config.Config

	// 初始化Viper
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(filename)
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
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&c, withJsonTag); err != nil {
			fmt.Println(err)
		}

		fmt.Println(c)
	})

	return &c
}
