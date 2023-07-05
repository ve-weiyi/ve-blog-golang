package glog

import (
	"testing"

	"go.uber.org/zap"
)

func TestName(t *testing.T) {
	Error("hello")
	Debug("123")
	Info("info")
	ReplaceZapGlobals()
	zap.L().Debug("hello world")
}
