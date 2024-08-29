package utils

import (
	"fmt"
)

// [a,b,c] -> a, b, c
func JoinArray[S ~[]E, E any](elems S, sep string) string {
	var out string

	for _, e := range elems {
		if out == "" {
			out = out + fmt.Sprintf("%v", e)
		} else {
			out = out + sep
			out = out + fmt.Sprintf("%v", e)
		}
	}

	return out
}

func AppendSlice(keys []string, key string) []string {
	for _, item := range keys {
		if item == key {
			return keys
		}
	}
	return append(keys, key)
}

func ExistSlice(keys []string, key string) bool {
	for _, id := range keys {
		if id == key {
			return true
		}
	}
	return false
}

// 删除数组中的重复元素
func RemoveSliceDuplicates(arr []string) (ret []string) {
	exists := make(map[string]struct{})
	for _, s := range arr {
		if _, ok := exists[s]; !ok {
			ret = append(ret, s)
			exists[s] = struct{}{}
		}
	}

	return
}

func MapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
