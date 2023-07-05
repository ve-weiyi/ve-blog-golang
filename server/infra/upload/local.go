package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/crypto"
)

type Local struct {
	Host      string // 访问地址
	LocalPath string // 本地路径
}

func (s *Local) UploadFile(prefix string, file *multipart.FileHeader) (url string, err error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	// 拼接新文件名
	filename := fmt.Sprintf("%s_%s%s", crypto.MD5V([]byte(name)), time.Now().Format("20060102150405"), ext)

	// 本地文件目录
	dir := path.Join(s.LocalPath, prefix)
	// 本地文件路径
	localPath := path.Join(dir, filename)
	// 访问文件路径
	accessPath := s.Host + "/" + localPath

	// 尝试创建上传目录
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Local.UploadFile os.MkdirAll() Filed, err:" + err.Error())
	}
	// 读取文件
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("Local.UploadFile file.Open() Filed, err:" + err.Error())
	}
	defer f.Close()
	// 创建文件
	out, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("Local.UploadFile os.Create() Filed, err:" + err.Error())
	}
	defer out.Close()
	// 传输（拷贝）文件内容
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		return "", fmt.Errorf("Local.UploadFile io.Copy() Filed, err:" + copyErr.Error())
	}

	return accessPath, nil
}

func (s *Local) DeleteFile(key string) error {
	p := s.LocalPath + "/" + key
	if strings.Contains(p, s.LocalPath) {
		if err := os.Remove(p); err != nil {
			return fmt.Errorf("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
