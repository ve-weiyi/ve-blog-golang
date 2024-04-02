package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/server/config"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
)

func Upload(c config.Upload) (upload.Uploader, error) {

	var cfg upload.UploadConfig
	var up upload.Uploader
	switch c.Mode {
	case "local":
		up = upload.NewLocal(c.Local)
	case "qiniu":
		up = upload.NewQiniu(&cfg)
	case "aliyun":
		up = upload.NewAliyunOSS(&cfg)
	default:
		up = upload.NewLocal(&cfg)
	}

	return up, nil
}
