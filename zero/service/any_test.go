package service

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

func Test_VisitFile(t *testing.T) {

	err := files.VisitFile("./rpc/blog/internal/logic", func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err

		}
		// 是目录，则跳过
		if f.IsDir() {
			return nil
		}

		if strings.Contains(f.Name(), "create_") {
			// 添加前缀 "gen_" 到文件名
			newName := strings.Replace(f.Name(), "create_", "add_", 1)

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
