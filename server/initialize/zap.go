package initialize

import (
	"log"
	"os"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/glog/zaplog"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/files"
)

// Zap 获取 zap.Logger
func Zap() {
	if ok, _ := files.PathExists(global.CONFIG.Zap.CacheDir); !ok { // 判断是否有Director文件夹
		log.Printf("create directory:%v\n", global.CONFIG.Zap.CacheDir)
		_ = os.Mkdir(global.CONFIG.Zap.CacheDir, os.ModePerm)
	}

	var cfg zaplog.ZapConfig
	cfg = zaplog.NewFileConfig()
	cfg.Mode = global.CONFIG.Zap.Mode
	cfg.Encoding = global.CONFIG.Zap.Format
	cfg.Path = global.CONFIG.Zap.CacheDir
	cfg.FileName = global.CONFIG.Zap.ServerName + ".log"
	cfg.Prefix = global.CONFIG.Zap.Prefix
	cfg.Level = global.CONFIG.Zap.Level
	cfg.ShowLine = global.CONFIG.Zap.EncodeCaller == "long"
	cfg.ShowColor = global.CONFIG.Zap.EncodeColorful

	global.LOG = glog.NewGlogger(1, cfg)
	global.LOG.Println("日志组件初始化成功！")
	return
}
