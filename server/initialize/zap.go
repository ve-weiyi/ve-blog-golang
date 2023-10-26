package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/zaplog"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/files"
)

// Zap 获取 zap.Logger
func Zap() {
	if ok, _ := files.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	var cfg zaplog.ZapConfig
	cfg = zaplog.NewFileConfig()
	if strings.HasPrefix(global.CONFIG.System.Env, "dev") {
		cfg = zaplog.NewConsoleConfig()
		cfg.Path = global.CONFIG.Zap.Director
		cfg.FileName = "blog.log"
		cfg.Prefix = ""
		cfg.Level = "Info"
		cfg.ShowColor = true
	} else {
		cfg = zaplog.NewFileConfig()
		cfg.Path = global.CONFIG.Zap.Director
		cfg.FileName = "blog.log"
		cfg.Prefix = ""
		cfg.Level = "Info"
		cfg.ShowColor = true
	}

	global.LOG = glog.NewGlogger(1, cfg)
	global.LOG.Println("日志组件初始化成功！")
	return
}
