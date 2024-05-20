package utils

// [a,b,c] -> a, b, c
func JoinArray(arr []string) string {
	var result string
	for i, v := range arr {
		result += v
		if i < len(arr)-1 {
			result += ", "
		}
	}
	return result
}

// 数组 [a, b, c] 转换为字符串 a<b<c>>
func JoinArrayBracket(params []string) string {
	var result string
	for i, val := range params {
		if i > 0 {
			result += "<"
		}
		result += val
	}
	for i := 0; i < len(params)-1; i++ {
		result += ">"
	}

	return result
}

// 删除数组中的重复元素
func RemoveDuplicates(input []string) []string {
	encountered := map[string]bool{}
	var result []string

	for _, v := range input {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
