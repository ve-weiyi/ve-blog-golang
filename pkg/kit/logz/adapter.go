package logz

import "go.uber.org/zap"

// zapLoggerAdapter 是 zap.Logger 到 Logger 接口的适配器
type zapLoggerAdapter struct {
	logger *zap.SugaredLogger
}

// NewZapLogger 创建基于 zap 的 Logger 实现
func NewZapLogger(logger *zap.Logger) Logger {
	return &zapLoggerAdapter{
		logger: logger.Sugar(),
	}
}

// NewDefaultLogger 创建默认的 Logger 实现
func NewDefaultLogger() Logger {
	return &zapLoggerAdapter{
		logger: S(),
	}
}

// Debug 实现 Logger 接口
func (l *zapLoggerAdapter) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *zapLoggerAdapter) Debugf(fmt string, args ...interface{}) {
	l.logger.Debugf(fmt, args...)
}

// Info 实现 Logger 接口
func (l *zapLoggerAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *zapLoggerAdapter) Infof(fmt string, args ...interface{}) {
	l.logger.Infof(fmt, args...)
}

// Warn 实现 Logger 接口
func (l *zapLoggerAdapter) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *zapLoggerAdapter) Warnf(fmt string, args ...interface{}) {
	l.logger.Warnf(fmt, args...)
}

// Error 实现 Logger 接口
func (l *zapLoggerAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *zapLoggerAdapter) Errorf(fmt string, args ...interface{}) {
	l.logger.Errorf(fmt, args...)
}

// Fatal 实现 Logger 接口
func (l *zapLoggerAdapter) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *zapLoggerAdapter) Fatalf(fmt string, args ...interface{}) {
	l.logger.Fatalf(fmt, args...)
}

// zapStructuredLoggerAdapter 是支持结构化日志的适配器
type zapStructuredLoggerAdapter struct {
	zapLoggerAdapter
}

// NewStructuredLogger 创建结构化日志 Logger
func NewStructuredLogger(logger *zap.Logger) StructuredLogger {
	return &zapStructuredLoggerAdapter{
		zapLoggerAdapter: zapLoggerAdapter{
			logger: logger.Sugar(),
		},
	}
}

// WithFields 实现 StructuredLogger 接口
func (l *zapStructuredLoggerAdapter) WithFields(fields map[string]interface{}) Logger {
	var args []interface{}
	for k, v := range fields {
		args = append(args, k, v)
	}
	return &zapLoggerAdapter{
		logger: l.logger.With(args...),
	}
}

// WithField 实现 StructuredLogger 接口
func (l *zapStructuredLoggerAdapter) WithField(key string, value interface{}) Logger {
	return &zapLoggerAdapter{
		logger: l.logger.With(key, value),
	}
}

// WithError 实现 StructuredLogger 接口
func (l *zapStructuredLoggerAdapter) WithError(err error) Logger {
	return &zapLoggerAdapter{
		logger: l.logger.With("error", err),
	}
}
