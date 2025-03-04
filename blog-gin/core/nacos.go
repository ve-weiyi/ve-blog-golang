package core

import (
	"encoding/json"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/config"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/nacos"
)

func Nacos(nc *nacos.NacosConfig) *config.Config {
	var c config.Config

	// 初始化Nacos
	na := nacos.New(nc)

	// 读取配置文件
	content, err := na.GetConfig()
	if err != nil {
		panic("nacos config read failed " + err.Error())
	}

	// 解析配置文件
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		panic(err)
	}

	// 监听配置文件变化
	na.AddListener(func(content string) {
		err = json.Unmarshal([]byte(content), &c)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(c)
	})

	return &c
}
