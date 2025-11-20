package logz

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/logz/zaplog"
)

// 同时输出到控制台和文件测试
func TestGlogConsoleAndFile(t *testing.T) {
	logs := zaplog.NewZapLogger(1, zaplog.NewConsoleConfig())
	//logj := zaplog.NewZapLogger(0, zaplog.NewFileConfig())
	Init(logs)

	Debug("This is a debug message")
	Info("This is an info message")
	Warn("This is a warning message")
	Error("This is an error message")
}

// 压力测试
func BenchmarkExample(b *testing.B) {
	logs := zaplog.NewZapLogger(0, zaplog.NewConsoleConfig())
	logj := zaplog.NewZapLogger(0, zaplog.NewFileConfig())
	for i := 0; i < b.N; i++ {
		// 在这里执行需要进行压力测试的代码
		logs.Sugar().Info("hello", i)
		logj.Sugar().Info("hello", i)
	}
}
