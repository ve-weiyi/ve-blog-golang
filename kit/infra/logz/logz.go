package logz

import (
	"encoding/json"

	"go.uber.org/zap"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/logz/zaplog"
)

// go使用zap + lumberjack重构项目的日志系统 https://blog.csdn.net/weixin_52000204/article/details/126651319
var defaultLogger *zap.Logger

func init() {
	defaultLogger = zaplog.NewZapLogger(1, zaplog.NewConsoleConfig())
}

func Init(lg *zap.Logger) {
	defaultLogger = lg
}

func Default() *zap.Logger {
	return defaultLogger
}

func Error(v ...interface{}) {

	defaultLogger.Sugar().Error(v...)
}

func Warn(v ...interface{}) {

	defaultLogger.Sugar().Warn(v...)
}

func Info(v ...interface{}) {

	defaultLogger.Sugar().Info(v...)
}

func Debug(v ...interface{}) {

	defaultLogger.Sugar().Debug(v...)
}

// print 风格
func Errorf(template string, args ...interface{}) {

	defaultLogger.Sugar().Errorf(template, args...)
}

func Warnf(template string, args ...interface{}) {

	defaultLogger.Sugar().Warnf(template, args...)
}

func Infof(template string, args ...interface{}) {

	defaultLogger.Sugar().Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {

	defaultLogger.Sugar().Debugf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {

	defaultLogger.Sugar().Errorw(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {

	defaultLogger.Sugar().Warnw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {

	defaultLogger.Sugar().Infow(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {

	defaultLogger.Sugar().Debugw(msg, keysAndValues...)
}

func Println(v ...interface{}) {

	defaultLogger.Sugar().Info(v...)
}

func JsonIndent(v ...interface{}) {

	bytes, _ := json.MarshalIndent(v, "", " ")
	defaultLogger.Sugar().Info("--->", string(bytes))
}
