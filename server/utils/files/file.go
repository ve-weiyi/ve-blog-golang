package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

// Size 获取文件大小
func Size(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

// Ext 获取文件后缀
func Ext(fileName string) string {
	return path.Ext(fileName)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// openExistFile 打开文件 不存在则创建
func OpenExistFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取目录: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("无此权限: %s", src)
	}

	err = MkDirIfNotExist(src)
	if err != nil {
		return nil, fmt.Errorf("创建文件夹错误: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败 :%v", err)
	}

	return f, nil
}

// 删除文件
func DeLFile(filePath string) error {
	return os.RemoveAll(filePath)
}

// @description: 文件移动供外部调用
// @param: src string, dst string(src: 源位置,绝对路径or相对路径, dst: 目标位置,绝对路径or相对路径,必须为文件夹)
func MoveFile(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	revoke := false
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}

// 深度遍历目录下的所有文件，包括目录和文件
func VisitFile(root string, visitFile func(path string, f os.FileInfo, err error) error) {
	err := filepath.Walk(root, visitFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// 向文件中写入内容
func WriteContentToFile(filename string, content string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
