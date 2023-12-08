package files

import (
	"fmt"
	"os"
)

// 检查文件夹或者文件是否存在
func IsExist(src string) bool {
	_, err := os.Stat(src)
	return !os.IsNotExist(err)
}

// 检查文件夹是否存在
func IsDirExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true
		}
	}

	return !os.IsNotExist(err)
}

// 检查文件是否存在
func IsFileExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if !fi.IsDir() {
			return true
		}
	}

	return !os.IsNotExist(err)
}

// MkDir 新建文件夹
// `os.Mkdir` 用于创建一个目录。
// `os.MkdirAll` 用于创建一个目录以及它的所有父目录（如果它们不存在的话）
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// IsNotExistMkDir 如果不存在则新建文件夹
func MkDirIfNotExist(src string) error {
	if !IsDirExist(src) {
		if err := MkDir(src); err != nil {
			return fmt.Errorf("create directory '%s' error: %v", src, err)
		}
	}
	return nil
}
