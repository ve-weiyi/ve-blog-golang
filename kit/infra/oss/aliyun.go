package oss

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Aliyun struct {
	cfg    *Config
	bucket *oss.Bucket
}

func (s *Aliyun) UploadHttpFile(file *multipart.FileHeader, prefix string, filename string) (url string, err error) {
	// 本地文件目录
	dir := path.Join(s.cfg.BasePath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.cfg.BucketUrl + "/" + localPath

	// 读取本地文件
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadHttpFile file.Open() Failed, err:" + err.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	// 上传策略
	bucket := s.bucket
	// 上传文件流。
	err = bucket.PutObject(localPath, f)
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadHttpFile formUploader.Put() Failed, err:" + err.Error())
	}

	return accessPath, nil
}

func (s *Aliyun) UploadLocalFile(filepath string, prefix string, filename string) (url string, err error) {
	// 本地文件目录
	dir := path.Join(s.cfg.BasePath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.cfg.BucketUrl + "/" + localPath

	// 读取本地文件
	f, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadHttpFile file.Open() Failed, err:" + err.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	// 上传策略
	bucket := s.bucket
	// 上传文件流。
	err = bucket.PutObject(localPath, f)
	if err != nil {
		return "", fmt.Errorf("Aliyun.UploadHttpFile formUploader.Put() Failed, err:" + err.Error())
	}

	return accessPath, nil
}

func (s *Aliyun) DeleteFile(key string) (err error) {
	bucket := s.bucket
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		return fmt.Errorf("Aliyun.DeleteFile bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func (s *Aliyun) ListFiles(prefix string, limit int) (keys []string, err error) {
	bucket := s.bucket

	result, err := bucket.ListObjectsV2(oss.Prefix(path.Join(s.cfg.BasePath, prefix)), oss.MaxKeys(limit))
	if err != nil {
		return nil, fmt.Errorf("Aliyun.ListFiles bucketManager.ListFiles() Filed, err:" + err.Error())
	}

	for _, entry := range result.Objects {
		keys = append(keys, entry.Key)
	}

	return keys, nil
}

func NewAliyunOSS(cfg *Config) *Aliyun {

	// 创建OSSClient实例。
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		panic(err)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(cfg.BucketName)
	if err != nil {
		panic(err)
	}

	return &Aliyun{
		cfg:    cfg,
		bucket: bucket,
	}
}
