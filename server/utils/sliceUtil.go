package utils

import jsoniter "github.com/json-iterator/go"

func AppendSlice(ids []string, key string) []string {
	for _, item := range ids {
		if item == key {
			return ids
		}
	}
	return append(ids, key)
}

func AppendJsonSlice(jsonArr string, key string) string {
	var ids []string
	if jsonArr != "" {
		err := jsoniter.Unmarshal([]byte(jsonArr), &ids)
		if err != nil {
			return jsonArr
		}
	}
	ids = AppendSlice(ids, key)
	jb, _ := jsoniter.Marshal(ids)
	return string(jb)
}

func IsExistSlice(ids []string, key string) bool {
	for _, id := range ids {
		if id == key {
			return true
		}
	}
	return false
}

func IsExistJsonSlice(jsonArr string, key string) bool {
	var ids []string
	if jsonArr != "" {
		err := jsoniter.Unmarshal([]byte(jsonArr), &ids)
		if err != nil {
			return false
		}
	}

	return IsExistSlice(ids, key)
}
