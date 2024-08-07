package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

func Test_len(t *testing.T) {
	fmt.Println(len("230f039bbd89dcb83974cac5041a1199d98ba9c1"))
}
func Test_VisitFile(t *testing.T) {

	err := files.VisitFile("./api/blog/internal/logic", func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err

		}
		// 是目录，则跳过
		if f.IsDir() {
			return nil
		}

		if strings.Contains(f.Name(), "list") {
			// 添加前缀 "gen_" 到文件名
			newName := strings.Replace(f.Name(), "get_", "find_", 1)

			// 修改文件名
			err = os.Rename(path, filepath.Join(filepath.Dir(path), newName))
		}
		return err
	})
	t.Log(err)
}

// add
// delete
// update
// find
