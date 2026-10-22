package gormlogx

import (
	"slices"
	"strings"
	"testing"
)

// recordWriter 实现 gorm 的 logger.Writer，记下最后一次调用。
type recordWriter struct {
	calls int
	msg   string
	args  []any
}

func (w *recordWriter) Printf(msg string, args ...any) {
	w.calls++
	w.msg = msg
	w.args = args
}

func TestGormWriter(t *testing.T) {
	t.Run("未注入 Writer 时回落标准输出", func(t *testing.T) {
		if w := NewGormWriter(nil, Config{}); w.logger == nil {
			t.Fatal("应回落为一个非空的 logger.Writer")
		}
	})

	t.Run("跳过关键词默认值与覆盖", func(t *testing.T) {
		if got, want := NewGormWriter(&recordWriter{}, Config{}).config.SkipKeywords, []string{"/gen/", "/gorm/"}; !slices.Equal(got, want) {
			t.Errorf("SkipKeywords = %v, want %v", got, want)
		}
		if got, want := NewGormWriter(&recordWriter{}, Config{SkipKeywords: []string{"/x/"}}).config.SkipKeywords, []string{"/x/"}; !slices.Equal(got, want) {
			t.Errorf("SkipKeywords = %v, want %v", got, want)
		}
	})

	t.Run("无参数时原样透传", func(t *testing.T) {
		inner := &recordWriter{}
		NewGormWriter(inner, Config{}).Printf("plain message")

		if inner.calls != 1 || inner.msg != "plain message" || len(inner.args) != 0 {
			t.Errorf("调用 = %d, msg = %q, args = %v", inner.calls, inner.msg, inner.args)
		}
	})

	t.Run("有参数时把 data[0] 改写成业务行号", func(t *testing.T) {
		inner := &recordWriter{}
		// 关键词命中栈上所有帧，行号必为空串，便于确定性断言改写本身。
		NewGormWriter(inner, Config{SkipKeywords: []string{"/"}}).
			Printf("%s [%.3fms] [rows:%v] %s", "gorm 内部行号", 1.5, 1, "SELECT 1")

		if inner.calls != 1 {
			t.Fatalf("调用次数 = %d, want 1", inner.calls)
		}
		if got := inner.args[0]; got != "" {
			t.Errorf("data[0] = %v, want 空串（被业务行号覆盖）", got)
		}
	})

	t.Run("fileLine 定位业务调用方", func(t *testing.T) {
		const selfFile = "writer_test.go"

		if got := callFileLine(0, nil); !strings.Contains(got, selfFile) {
			t.Errorf("fileLine = %q, want 含 %q", got, selfFile)
		}
		if got := callFileLine(0, []string{selfFile}); strings.Contains(got, selfFile) {
			t.Errorf("fileLine = %q, want 跳过 %q", got, selfFile)
		}
		if got := callFileLine(0, []string{"/"}); got != "" {
			t.Errorf("所有帧都命中关键词时 fileLine = %q, want 空串", got)
		}
	})
}

// callFileLine 只是多包一层，让 fileLine 的调用栈里多一帧本文件，
// 这样 skip=0 时它能定位到本测试文件而不是 testing 包的内部帧。
func callFileLine(skip int, keywords []string) string {
	return fileLine(skip, keywords)
}
