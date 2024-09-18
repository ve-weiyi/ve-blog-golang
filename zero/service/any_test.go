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

func TestQS(t *testing.T) {
	arr := []int{4, 2, 5, 1, 6, 5, 7, 5, 9, 0, 11}
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func sort(arr []int, low, high int) {
	if low < high {
		mid := part2(arr, low, high)
		sort(arr, low, mid-1)
		sort(arr, mid+1, high)
	}
}

// 寻找基准:左右指针法
func part(arr []int, low, high int) int {
	pv := arr[low]
	left := low
	right := high

	for left < right {
		// 右边找小
		for left < right && arr[right] >= pv {
			right--
		}

		// 左边找大
		for left < right && arr[left] <= pv {
			left++
		}

		// 交换位置
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	//交换结果与哨兵位置
	arr[low], arr[left] = arr[left], arr[low]

	return left
}

// 寻找基准:前后指针法。前指针找小，后指针交换
func part2(arr []int, low, high int) int {
	pv := arr[low]
	l := low
	for r := low + 1; r <= high; r++ {
		// 前指针找小
		if pv > arr[r] {
			l++
			// 交换与后指针的位置
			arr[r], arr[l] = arr[l], arr[r]
		}

	}

	arr[l], arr[low] = arr[low], arr[l]

	return l
}
