package initialize

import (
	"fmt"
	"os"

	"github.com/ve-weiyi/ve-blog-golang/server/global"

	"github.com/ve-weiyi/go-sdk/utils/copy"
	"github.com/ve-weiyi/go-sdk/utils/file"
	"github.com/ve-weiyi/go-sdk/utils/glog"
	"github.com/ve-weiyi/go-sdk/utils/glog/zaplog"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() {
	if ok, _ := file.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cfg := zaplog.ZapConfig{}

	copy.DeepCopyByJson(global.CONFIG.Zap, &cfg)

	glog.ReplaceZapGlobals(cfg)
	global.LOG = glog.NewGlogger(1, cfg)

	global.LOG.Printf("日志组件初始化成功！")
	return
}
