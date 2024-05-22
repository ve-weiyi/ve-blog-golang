package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/copyutil"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

func Upload() {

	var cfg upload.UploadConfig
	var up upload.Uploader
	switch global.CONFIG.Upload.Mode {
	case "local":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Local, &cfg)
		if err != nil {
			glog.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload.NewLocal(&cfg)
	case "qiniu":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Qiniu, &cfg)
		if err != nil {
			glog.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload.NewQiniu(&cfg)
	case "aliyun":
		err := copyutil.DeepCopyByJson(global.CONFIG.Upload.Aliyun, &cfg)
		if err != nil {
			glog.Errorf("上传组件初始化失败！%v", err)
		}
		up = upload.NewAliyunOSS(&cfg)
	default:
		up = upload.NewLocal(&cfg)
	}

	global.Uploader = up
	if global.Uploader != nil {
		glog.Infof("上传组件初始化成功！%v", global.CONFIG.Upload.Mode)
	}
}
