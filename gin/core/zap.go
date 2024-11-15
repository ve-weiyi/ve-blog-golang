package core

import (
	"github.com/ve-weiyi/ve-blog-golang/gin/config"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog/zaplog"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

// 初始化日志组件 zap
func SetLog(c config.Zap) {
	err := files.MkDir(c.CacheDir)
	if err != nil {
		panic(err)
	}

	var cfg zaplog.ZapConfig
	cfg = zaplog.NewFileConfig()
	cfg.Mode = c.Mode
	cfg.Encoding = c.Format
	cfg.Prefix = c.Prefix
	cfg.Level = c.Level
	cfg.ShowLine = c.EncodeCaller == "long"
	cfg.ShowColor = c.EncodeColorful

	cfg.FileName = c.Filename + ".log"
	cfg.Path = c.CacheDir
	cfg.KeepDays = c.MaxAge

	glog.Init(1, cfg)
	return
}
