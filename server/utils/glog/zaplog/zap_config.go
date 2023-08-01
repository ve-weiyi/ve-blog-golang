package zaplog

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

// zapcore.Core 需要三个配置
// ——Encoder
// ——WriteSyncer
// ——LogLevel

type ZapConfig struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

func (cfg *ZapConfig) GetEncoderConfig() FormatConfig {
	cfgFormat := FormatConfig{
		Prefix:        cfg.Prefix,
		StacktraceKey: cfg.StacktraceKey,
		LogInConsole:  cfg.LogInConsole,
		ShowLine:      cfg.ShowLine,
		EncodeLevel:   cfg.ZapEncodeLevel(),
		Format:        cfg.Format,
	}
	return cfgFormat
}

func (cfg *ZapConfig) GetWriterConfig() WriterConfig {
	cfgWriter := WriterConfig{
		FileDir:    cfg.Director,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     cfg.MaxAge,
		Compress:   false,
	}
	return cfgWriter
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (cfg *ZapConfig) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case cfg.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case cfg.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case cfg.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case cfg.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (cfg *ZapConfig) TransportLevel() zapcore.Level {
	cfg.Level = strings.ToLower(cfg.Level)
	switch cfg.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
