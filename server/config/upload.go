package config

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
)

type Upload struct {
	Mode   string               `mapstructure:"mode"`
	Local  *upload.UploadConfig `mapstructure:"local"`
	Aliyun *upload.UploadConfig `mapstructure:"aliyun"`
	Qiniu  *upload.UploadConfig `mapstructure:"qiniu"`
}
