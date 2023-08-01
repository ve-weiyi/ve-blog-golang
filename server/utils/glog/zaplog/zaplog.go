package zaplog

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger = NewZapLogger(0, GetDefaultConfig())

func GetDefaultZapLogger() *zap.Logger {
	return _logger
}

func GetDefaultConfig() ZapConfig {
	return ZapConfig{
		Level:         "debug",
		Prefix:        "[Default]",
		Format:        "json",
		Director:      fmt.Sprintf("runtime/log"),
		StacktraceKey: "stacktrace",
		MaxAge:        30,
		LogInConsole:  true,
		ShowLine:      true,
		EncodeLevel:   "CapitalLevelEncoder",
	}
}

// 默认的zap不需要增加skip
func NewZapLogger(skip int, cfg ZapConfig) *zap.Logger {
	// 使用了core的NewTee
	cores := zapcore.NewTee(
		GetEncoderCore(cfg)...,
	)
	// 创建一个将日志写入 WriteSyncer 的核心。
	// Glogger.Debug->skip1  glog.Debug->skip2
	logger := zap.New(
		cores,
	)

	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
		logger = logger.WithOptions(zap.AddCallerSkip(skip))
	}

	return logger
}

func GetEncoderCore(cfg ZapConfig) []zapcore.Core {

	cfgFormat := cfg.GetEncoderConfig()
	cfgWriter := cfg.GetWriterConfig()

	//使用不同后缀收集日志
	lvCores := make([]zapcore.Core, 0, 7)
	for level := cfg.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		encoder := GetJsonEncoder(cfgFormat)
		writer := zapcore.AddSync(GetFileWriter(cfgWriter, level.String()))
		lvCores = append(lvCores, zapcore.NewCore(encoder, writer, GetLevelPriority(level)))
	}

	if cfg.LogInConsole {
		//控制台打印、console格式
		encoder := GetConsoleEncoder(cfgFormat)
		writer := zapcore.AddSync(os.Stderr)
		lvCores = append(lvCores, zapcore.NewCore(encoder, writer, cfg.TransportLevel()))
	}

	return lvCores
}

// GetCoreSimple 获取Encoder的 zapcore.Core
// 只分为 info 和 error两个等级
func GetCoreSimple(cfg ZapConfig) []zapcore.Core {
	cfgFormat := cfg.GetEncoderConfig()
	cfgWriter := cfg.GetWriterConfig()
	//文件打印、json格式
	jsonEncoder := GetJsonEncoder(cfgFormat)
	//控制台打印、console格式
	consoleEncoder := GetConsoleEncoder(cfgFormat)

	//保活信息:debug，info ,交由运维同学监控
	saveLv := GetLevelLowThan(zapcore.InfoLevel)
	//错误日志:Warn，Error，Fatal，Panic ,开发成员关注
	errorLv := GetLevelUpThan(zapcore.WarnLevel)

	cores := []zapcore.Core{
		// 保活日志
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(GetFileWriter(cfgWriter, "info")), saveLv),
		// 错误日志:输入到文件中，使用json格式，无颜色
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(GetFileWriter(cfgWriter, "error")), errorLv),

		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), zapcore.DebugLevel),
	}

	return cores
}
