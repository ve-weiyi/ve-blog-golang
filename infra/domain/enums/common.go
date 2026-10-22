package enums

// IsValidIsDelete 判断删除标记是否合法。
func IsValidIsDelete(v int64) bool {
	return v == 0 || v == 1
}
