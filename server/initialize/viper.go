package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) {
	var config string

	if len(path) != 0 {
		// 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		log.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	} else {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			// 使用默认的
			config = "config.yaml"
			log.Printf("您正在使用默认的环境变量,config的路径为%s\n", config)
		} else { // 命令行参数不为空 将值赋值于config
			log.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			log.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		log.Println(err)
	}
	log.Printf("配置文件读取成功！")
	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效

	global.VP = v
}
