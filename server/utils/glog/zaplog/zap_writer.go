package zaplog

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志按大小切割和备份个数、文件有效期，一般选zap+lumberjack
// 日志按天或者小时切割，一般选zap+file-rotatelogs
type WriterConfig struct {
	FileDir    string //日志文件的路径
	MaxSize    int    //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int    //保留旧文件的最大个数
	MaxAge     int    //保留旧文件的最大天数
	Compress   bool   //是否压缩/归档旧文件
	//LogInConsole bool   //是否打印控制台
}

// GetFileWriter 按时间分割日志文件 runtime/log/2022-11
// rotate库于2022.11.9开始不再更新
func GetFileWriter(cfg WriterConfig, level string) *rotatelogs.RotateLogs {
	return GetTimeWriter(cfg, level)
}

// GetTimeWriter 按时间分割日志文件 runtime/log/2022-11
// rotate库于2022.11.9开始不再更新
func GetTimeWriter(cfg WriterConfig, level string) *rotatelogs.RotateLogs {
	fileWriter, err := rotatelogs.New(
		path.Join(cfg.FileDir, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		return nil
	}
	//fileWriter.Rotate()
	return fileWriter
}

// GetSizeWriter 按文件大小分割日志文件
func GetSizeWriter(cfg WriterConfig, level string) *lumberjack.Logger {
	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(cfg.FileDir, "%Y-%m-%d", level+".log"),
		MaxSize:    cfg.MaxSize,    //最大M数，超过则切割
		MaxBackups: cfg.MaxBackups, //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     cfg.MaxAge,     //保存30天
		Compress:   cfg.Compress,   //是否压缩
	}
	//fileWriter.Rotate()
	return fileWriter
}
