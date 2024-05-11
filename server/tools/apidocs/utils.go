package apidocs

import (
	"strings"
)

func convertGoTypeToTsType(name string) string {
	if strings.HasPrefix(name, "*") {
		return convertGoTypeToTsType(name[1:]) // 指针
	}
	if strings.HasPrefix(name, "[]") {
		return convertGoTypeToTsType(name[2:]) + "[]" // 数组
	}
	if strings.LastIndex(name, ".") > 0 {
		return convertGoTypeToTsType(name[strings.LastIndex(name, ".")+1:]) // 去掉包名
	}
	switch name {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "file":
		return "File"
	case "Time":
		return "string"
	case "FileHeader":
		return "File"
	case "interface{}", "object":
		return "any"
	default:
		return name
	}
}
