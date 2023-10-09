package initialize

import (
	"time"

	"github.com/orca-zhang/ecache"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/upload"
)

func OtherInit() {
	global.BlackCache = ecache.NewLRUCache(16, 200, 10*time.Second).LRU2(1024)
	global.Uploader = upload.NewOss(&global.CONFIG.Upload)
	if global.Uploader != nil {
		global.LOG.Infof("上传组件初始化成功！%v", global.CONFIG.Upload.Mode)
	}
}
