package upload

import (
	"mime/multipart"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
)

// Uploader 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type Uploader interface {
	UploadFile(prefix string, file *multipart.FileHeader) (string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss(cfg *properties.Upload) Uploader {

	switch cfg.Mode {
	case "local":
		return NewLocal(&cfg.Local)
	case "aliyun":
		o, _ := NewAliyunOSS(&cfg.Aliyun)
		return o
	case "qiniu":
		o := NewQiniu(&cfg.Qiniu)
		return o
	default:
		return NewLocal(&cfg.Local)
	}
}

func NewLocal(cfg *properties.Local) *Local {
	return &Local{
		Host:      cfg.Url,
		LocalPath: cfg.Path,
	}
}
