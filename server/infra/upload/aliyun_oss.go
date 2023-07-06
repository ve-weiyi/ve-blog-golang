package upload

import (
	"fmt"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"

	"github.com/ve-weiyi/go-sdk/utils/crypto"
)

type AliyunOSS struct {
	cfg    *properties.Aliyun
	bucket *oss.Bucket
}

func (s *AliyunOSS) UploadFile(prefix string, file *multipart.FileHeader) (url string, err error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	// 拼接新文件名
	filename := fmt.Sprintf("%s_%s%s", crypto.MD5V([]byte(name)), time.Now().Format("20060102150405"), ext)

	// 本地文件目录
	dir := path.Join(s.cfg.BasePath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.cfg.BucketUrl + "/" + localPath

	// 上传策略
	bucket := s.bucket
	// 读取本地文件
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadFile file.Open() Failed, err:" + err.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	// 上传文件流。
	err = bucket.PutObject(localPath, f)
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadFile formUploader.Put() Failed, err:" + err.Error())
	}

	return accessPath, nil
}

func (s *AliyunOSS) DeleteFile(key string) (err error) {
	bucket := s.bucket
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		return fmt.Errorf("Aliyun.DeleteFile bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewAliyunOSS(cfg *properties.Aliyun) (*AliyunOSS, error) {

	// 创建OSSClient实例。
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(cfg.BucketName)
	if err != nil {
		return nil, err
	}

	return &AliyunOSS{
		cfg:    cfg,
		bucket: bucket,
	}, nil
}
