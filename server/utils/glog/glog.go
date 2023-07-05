package glog

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/glog/zaplog"
)

var _glogger *Glogger

// 默认调用
func init() {
	//包级log函数 skip->2, glog.log函数 skip->1
	_glogger = NewGlogger(2, zaplog.GetDefaultConfig())
}

func ReplaceZapGlobals(cfgs ...zaplog.ZapConfig) {
	if len(cfgs) > 0 {
		zap.ReplaceGlobals(NewGlogger(0, cfgs[0]).Logger())
		return
	}
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(zaplog.GetDefaultZapLogger())
}

func NewGlogger(skip int, cfg zaplog.ZapConfig) *Glogger {
	logger := zaplog.NewZapLogger(skip, cfg)

	//初始化内部类
	mlog := new(Glogger)

	mlog.path = cfg.Director
	mlog.level = zapcore.DebugLevel
	mlog.rotateMu = &sync.Mutex{}
	mlog.rolling = false
	mlog.lastRotate = time.Now()

	mlog.rlog = nil
	mlog.log = logger
	mlog.sugar = logger.Sugar()

	return mlog
}

func Sync() {
	_glogger.log.Sync()
}

// active

func Error(v ...interface{}) {
	_glogger.Error(v...)
}

func Warn(v ...interface{}) {
	_glogger.Warn(v...)
}

func Info(v ...interface{}) {
	_glogger.Info(v...)
}

func Debug(v ...interface{}) {
	_glogger.Debug(v...)
}

func Errorw(format string, v ...interface{}) {
	_glogger.Errorw(format, v...)
}

func Warnw(format string, v ...interface{}) {
	_glogger.Warnw(format, v...)
}

func Infow(format string, v ...interface{}) {
	_glogger.Infow(format, v...)
}

func Debugw(format string, v ...interface{}) {
	_glogger.Debugw(format, v...)
}
