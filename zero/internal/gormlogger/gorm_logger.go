package gormlogger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

type GormWriter struct {
	l logger.Writer
}

func NewGormWriter() *GormWriter {
	return &GormWriter{
		l: log.New(os.Stdout, "\r\n", log.LstdFlags),
	}
}

// Printf 格式化打印日志
// message:%s \n[%.3fms] [rows:%v] %s
// data:[line,time,rows,sql]
func (w *GormWriter) Printf(message string, data ...interface{}) {
	data[0] = FileWithLineNum()
	w.l.Printf(message, data...)
}

var gormSourceDir string = "gorm"
var genSourceDir string = "gen"

// 拦截打印行号，过滤gen和gorm的行号
// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, gormSourceDir) && !strings.HasSuffix(file, "gen.go")) {
			return fmt.Sprintf("%v:%v", file, strconv.FormatInt(int64(line), 10))
		}
	}
	return ""
}
