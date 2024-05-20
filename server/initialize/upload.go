package initialize

import (
	upload2 "github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/copyutil"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

func Upload() {

	var cfg upload2.UploadConfig
	var up upload2.Uploader
	switch global.CONFIG.Upload.Mode {
	case "local":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Local, &cfg)
		if err != nil {
			global.LOG.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload2.NewLocal(&cfg)
	case "qiniu":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Qiniu, &cfg)
		if err != nil {
			global.LOG.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload2.NewQiniu(&cfg)
	case "aliyun":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Aliyun, &cfg)
		global.LOG.JsonIndent("cfg", cfg)
		if err != nil {
			global.LOG.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload2.NewAliyunOSS(&cfg)
	default:
		up = upload2.NewLocal(&cfg)
	}

	global.Uploader = up
	if global.Uploader != nil {
		global.LOG.Infof("上传组件初始化成功！%v", global.CONFIG.Upload.Mode)
	}
}
