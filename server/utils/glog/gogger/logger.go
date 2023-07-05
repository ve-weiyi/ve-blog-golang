package gogger

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type gormWriter struct {
	logger.Writer
	ZapLogger *zap.Logger
	UseZapLog bool
}

// NewWriter gormWriter 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *gormWriter {
	return &gormWriter{
		Writer:    w,
		ZapLogger: zap.L(),
		UseZapLog: false,
	}
}

// Printf 格式化打印日志
// message:%s \n[%.3fms] [rows:%v] %s
// data:[line,time,rows,sql]
func (w *gormWriter) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = w.UseZapLog
	data[0] = FileWithLineNum()
	if logZap {
		w.Writer.Printf(message, data...)
		//global.GVA_LOG.Info(fmt.Sprintf(message, data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

var gormSourceDir string = "gorm"
var genSourceDir string = "gen"

// 拦截打印行号，过滤gen和gorm的行号
// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 6; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		//log.Println(file)
		if ok && (!strings.Contains(file, gormSourceDir) && !strings.Contains(file, genSourceDir) || strings.HasSuffix(file, "_test.go")) {
			return fmt.Sprintf("'%v:%v'", file, strconv.FormatInt(int64(line), 10))
		}
	}

	return ""
}
