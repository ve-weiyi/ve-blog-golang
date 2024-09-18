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

func TestSort(t *testing.T) {
	arr := []int{3, 1, 4, 5, 6, 3, 2}
	quickSort(arr, 0, 6)

	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := part(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

// 分区函数：将数组划分为两部分，左边部分小于基准，右边部分大于基准
func part(arr []int, low, high int) int {
	pv := arr[low] // 选择第一个元素为基准
	left, right := low, high

	for left < right {
		// 从右往左找第一个小于基准的元素
		for left < right && arr[right] >= pv {
			right--
		}
		// 从左往右找第一个大于基准的元素
		for left < right && arr[left] <= pv {
			left++
		}
		// 交换两个指针指向的元素
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	// 将基准元素放到正确的位置
	arr[low], arr[left] = arr[left], arr[low]
	return left
}
