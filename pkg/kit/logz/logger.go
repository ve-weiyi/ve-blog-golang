package logz

// Logger 日志接口，定义了常用的日志方法
type Logger interface {
	// Debug 调试级别日志
	Debug(args ...interface{})
	Debugf(fmt string, args ...interface{})

	// Info 信息级别日志
	Info(args ...interface{})
	Infof(fmt string, args ...interface{})

	// Warn 警告级别日志
	Warn(args ...interface{})
	Warnf(fmt string, args ...interface{})

	// Error 错误级别日志
	Error(args ...interface{})
	Errorf(fmt string, args ...interface{})

	// Fatal 致命错误级别日志（会导致程序退出）
	Fatal(args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

// StructuredLogger 结构化日志接口，支持字段化日志
type StructuredLogger interface {
	Logger

	// WithFields 添加字段到日志上下文
	WithFields(fields map[string]interface{}) Logger

	// WithField 添加单个字段到日志上下文
	WithField(key string, value interface{}) Logger

	// WithError 添加错误到日志上下文
	WithError(err error) Logger
}
