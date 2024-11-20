package oss

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Local struct {
	cfg *Config
}

func (s *Local) UploadHttpFile(file *multipart.FileHeader, prefix string, filename string) (url string, err error) {
	// 本地文件目录
	dir := path.Join(s.cfg.BasePath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.cfg.BucketUrl + "/" + localPath

	// 尝试创建上传目录
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile os.MkdirAll() Filed, err:" + err.Error())
	}

	// 创建目标文件
	out, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile os.Create() Filed, err:" + err.Error())
	}
	defer out.Close()

	// 读取文件
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile file.Open() Filed, err:" + err.Error())
	}
	defer f.Close()

	// 传输（拷贝）文件内容
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		return "", fmt.Errorf("Local.UploadHttpFile io.Copy() Filed, err:" + copyErr.Error())
	}

	return accessPath, nil
}

func (s *Local) UploadLocalFile(filepath string, prefix string, filename string) (url string, err error) {
	// 本地文件目录
	dir := path.Join(s.cfg.BasePath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.cfg.BucketUrl + "/" + localPath

	// 尝试创建上传目录
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile os.MkdirAll() Filed, err:" + err.Error())
	}

	// 创建目标文件
	out, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile os.Create() Filed, err:" + err.Error())
	}
	defer out.Close()

	// 读取文件
	f, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("Local.UploadHttpFile file.Open() Filed, err:" + err.Error())
	}
	defer f.Close()

	// 传输（拷贝）文件内容
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		return "", fmt.Errorf("Local.UploadHttpFile io.Copy() Filed, err:" + copyErr.Error())
	}

	return accessPath, nil
}

func (s *Local) DeleteFile(key string) error {
	p := s.cfg.BucketUrl + "/" + key
	if strings.Contains(p, s.cfg.BasePath) {
		if err := os.Remove(p); err != nil {
			return fmt.Errorf("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

func (s *Local) ListFiles(prefix string, limit int) (keys []string, err error) {
	// 获取指定目录下的所有文件
	var keysList []string

	err = filepath.Walk(path.Join(s.cfg.BasePath, prefix), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果遇到错误，返回
		}

		// 跳过目录，只处理文件
		if info.IsDir() {
			return nil
		}

		keysList = append(keysList, path)

		// 如果已经达到限制数量，提前返回
		if len(keysList) >= limit {
			return fmt.Errorf("limit reached") // 触发 Walk 停止
		}

		return nil
	})

	if err != nil && err.Error() != "limit reached" {
		return nil, err // 如果是其他错误，返回
	}

	// 返回符合条件的文件列表
	return keysList, nil
}

func NewLocal(cfg *Config) *Local {
	return &Local{
		cfg: cfg,
	}
}
