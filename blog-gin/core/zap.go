package core

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/config"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/logz"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/logz/zaplog"
)

// 初始化日志组件 zap
func SetLog(c config.Zap) {
	var cfg zaplog.ZapConfig

	switch c.Mode {
	case "console":
		cfg = zaplog.NewConsoleConfig()
	case "file":
		cfg = zaplog.NewFileConfig()
	}

	cfg.Mode = c.Mode
	cfg.Encoding = c.Encoding
	cfg.Prefix = c.ServiceName
	cfg.Level = c.Level
	cfg.KeepDays = c.KeepDays

	logz.Init(zaplog.NewZapLogger(1, cfg))
	logz.Infof("zap log init success. mode:%v, level:%v", c.Mode, c.Level)
	return
}
