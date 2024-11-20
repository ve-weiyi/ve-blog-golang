package oss

import (
	"fmt"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

func FileNameFromHeader(file *multipart.FileHeader) string {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	// 拼接新文件名
	filename := fmt.Sprintf("%s-%s%s", name, time.Now().Format("20060102150405"), ext)

	return filename
}
