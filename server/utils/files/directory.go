package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			fmt.Println("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				fmt.Println("create directory" + v)
				return err
			}
		}
	}
	return err
}

// 深度遍历目录下的所有文件，包括目录和文件
func VisitFile(root string, visitFile func(path string, f os.FileInfo, err error) error) {
	err := filepath.Walk(root, visitFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
