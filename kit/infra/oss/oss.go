package oss

import (
	"io"
	"mime/multipart"
)

type Config struct {
	Zone            string `json:"zone"`
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	BucketName      string `json:"bucket-name"`
	BucketUrl       string `json:"bucket-url"`
	BasePath        string `json:"base-path"`
}

// OSS（Object Storage Service，对象存储服务）
type OSS interface {
	// return BucketUrl+BasePath+prefix+filename
	UploadHttpFile(file *multipart.FileHeader, prefix string, filename string) (url string, err error)
	UploadLocalFile(filepath string, prefix string, filename string) (url string, err error)
	DeleteFile(key string) error
	ListFiles(prefix string, limit int) (urls []string, err error)
}

type File interface {
	Open() (File, error)
	io.Writer
}
