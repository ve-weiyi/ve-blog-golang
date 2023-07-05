// Copyright © 2015-2018 Anker Innovations Technology Limited All Rights Reserved.
package glog

import (
	"encoding/json"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// go使用zap + lumberjack重构项目的日志系统 https://blog.csdn.net/weixin_52000204/article/details/126651319
type Glogger struct {
	//rlog  *lumberjack.Logger
	rlog  interface{}
	log   *zap.Logger        //并重性能与易用性，支持结构化和 printf 风格的日志记录。
	sugar *zap.SugaredLogger // 非常强调性能，不提供 printf 风格的 api（减少了 interface{} 与 反射的性能损耗）

	path       string
	level      zapcore.Level
	rotateMu   *sync.Mutex
	rolling    bool
	lastRotate time.Time
}

func (mlog *Glogger) Logger() *zap.Logger {
	return mlog.log
}

func (mlog *Glogger) checkRotate() {
	if !mlog.rolling {
		return
	}
	jack, ok := mlog.rlog.(*lumberjack.Logger)
	if ok {
		n := time.Now()
		if mlog.differentDay(n) {
			mlog.rotateMu.Lock()
			defer mlog.rotateMu.Unlock()

			// 获得锁之后再次检查是否是不同日期
			// 避免上一次调用已经切割日志,
			if mlog.differentDay(n) {
				jack.Rotate()
				mlog.lastRotate = n
			}
		}
	}
}

// 判断是不是换天了，如果换天了就要重新调用rotate()
func (mlog *Glogger) differentDay(t time.Time) bool {
	y, m, d := mlog.lastRotate.Year(), mlog.lastRotate.Month(), mlog.lastRotate.Day()
	return y != t.Year() || m != t.Month() || d != t.Day()
}

func (mlog *Glogger) EnableDailyFile() {
	mlog.rolling = true
}

func (mlog *Glogger) Error(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Error(v...)
}

func (mlog *Glogger) Warn(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warn(v...)
}

func (mlog *Glogger) Info(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Glogger) Debug(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debug(v...)
}

// print 风格
func (mlog *Glogger) Errorf(template string, args ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Errorf(template, args...)
}

func (mlog *Glogger) Warnf(template string, args ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warnf(template, args...)
}

func (mlog *Glogger) Infof(template string, args ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Infof(template, args...)
}

func (mlog *Glogger) Debugf(template string, args ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debugf(template, args...)
}

func (mlog *Glogger) Errorw(msg string, keysAndValues ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Errorw(msg, keysAndValues...)
}

func (mlog *Glogger) Warnw(msg string, keysAndValues ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warnw(msg, keysAndValues...)
}

func (mlog *Glogger) Infow(msg string, keysAndValues ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Infow(msg, keysAndValues...)
}

func (mlog *Glogger) Debugw(msg string, keysAndValues ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debugw(msg, keysAndValues...)
}

func (mlog *Glogger) Println(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Glogger) Printf(template string, args ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Infof(template, args...)
}

func (mlog *Glogger) GetUnderlyingLogger() *zap.Logger {
	return mlog.log
}

func (mlog *Glogger) JsonIndent(v ...interface{}) {
	mlog.checkRotate()
	bytes, _ := json.MarshalIndent(v, "", " ")
	mlog.sugar.Info("--->", string(bytes))
}
