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

func removeDuplicates(keys []string) {

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
