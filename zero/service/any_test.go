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
	fmt.Println(len(`["A3004","A3004311","A3012","A3025","A3027","A3028","A3029","A3029XR","A3030","A3033","A3033EU","A3035","A3040","A3045","A3045ZA1","A3062","A3201","A3600","A3850","A3850Dongle","A3871","A3872","A3873","A3876","A3909","A3910","A3926","A3913","A3930","A3927EU","A3927","A3926Z11","A3931","A3936","A3935W","A3935","A3931XR","A3933","A3944EU","A3944","A3943","A3939WEU","A3939","A3949P25","A3949","A3948A25","A3948","A3947","A3945","A3958","A3955","A3953","A3952","A3951","A3949R50","A3982EU","A3982","A3968","A3961"`))
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
