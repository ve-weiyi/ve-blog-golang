package logws

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(l *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: l}
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *ZapLogger) Debugf(format string, value ...interface{}) {
	l.logger.Sugar().Debugf(format, value...)
}

func (l *ZapLogger) Warningf(format string, value ...interface{}) {
	l.logger.Sugar().Warnf(format, value...)
}

func (l *ZapLogger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *ZapLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *ZapLogger) Warning(message string) {
	l.logger.Warn(message)
}

func (l *ZapLogger) Error(message string) {
	l.logger.Error(message)
}
