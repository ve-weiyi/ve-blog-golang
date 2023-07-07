package inject

import (
	"go/token"
	"reflect"
	"strconv"
)

func tokenToKind(t token.Token) reflect.Kind {
	switch t {
	case token.INT:
		return reflect.Int
	case token.FLOAT:
		return reflect.Float64
	case token.STRING:
		return reflect.String
	case token.CHAR:
		return reflect.Int32
	default:
		return reflect.Invalid
	}
}

func kindToToken(k reflect.Kind) token.Token {
	switch k {
	case reflect.Int:
		return token.INT
	case reflect.Float64:
		return token.FLOAT
	case reflect.String:
		return token.STRING
	case reflect.Int32:
		return token.CHAR
	default:
		return token.ILLEGAL
	}
}

func inferType(str string) (interface{}, error) {
	// 尝试将字符串解析为int
	i, err := strconv.Atoi(str)
	if err == nil {
		return i, nil
	}

	// 尝试将字符串解析为float
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return f, nil
	}

	// 尝试将字符串解析为带引号的string
	s, err := strconv.Unquote(str)
	if err == nil {
		return s, nil
	}

	// 如果都不匹配，则返回原始字符串
	return str, nil
}
