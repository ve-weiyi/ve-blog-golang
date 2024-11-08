package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

func Test_len(t *testing.T) {
	tx := "哈哈zz"
	fmt.Println(len(tx))
	fmt.Println(maskNickName(tx))

}

func maskNickName(nickName string) string {
	if len(nickName) <= 3 {
		return nickName
	}

	return nickName[:3] + "***"
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

func TestMerge(t *testing.T) {
	nums1 := []int{1, 0}
	m := 1
	nums2 := []int{2}
	n := 1
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		fmt.Println(nums1)
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		fmt.Println(nums1)
		return
	}

	i := 0
	l := 0
	for i+l < m+n {
		a := nums1[i+l]
		b := nums2[l]

		if a <= b && a != 0 {
			// 加入a
			i++
		} else {
			// 加入b
			nums1 = slices.Insert(nums1, i+l, b)[0 : m+n]
			l++
		}
	}
	fmt.Println(nums1)
}
