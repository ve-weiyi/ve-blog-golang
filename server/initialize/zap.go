package initialize

import (
	"log"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog/zaplog"
	"github.com/ve-weiyi/ve-blog-golang/server/global"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

// Zap 获取 zap.Logger
func Zap() {
	err := files.MkDir(global.CONFIG.Zap.CacheDir)
	if err != nil {
		log.Println(err)
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

	glog.Init(1, cfg)
	glog.Println("日志组件初始化成功！")
	return
}
