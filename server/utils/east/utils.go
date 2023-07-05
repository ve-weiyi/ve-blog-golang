package east

import (
	"go/token"
	"reflect"
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
