package slicex

import (
	"fmt"
)

func AppendSlice(keys []string, key string) []string {
	for _, item := range keys {
		if item == key {
			return keys
		}
	}
	return append(keys, key)
}

func IsExistSlice(keys []string, key string) bool {
	for _, id := range keys {
		if id == key {
			return true
		}
	}
	return false
}

// RemoveStringSliceDuplicate O(n) 复杂度去重.
func RemoveStringSliceDuplicate(arr []string) (ret []string) {
	exists := make(map[string]struct{})
	for _, s := range arr {
		if _, ok := exists[s]; !ok {
			ret = append(ret, s)
			exists[s] = struct{}{}
		}
	}

	return
}

func Join[S ~[]E, E any](elems S, sep string) string {
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
