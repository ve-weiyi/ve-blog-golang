package invent

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

var StdMapUtils = map[string]any{
	"Case2Camel":           jsonconv.Case2Camel,
	"Case2Snake":           jsonconv.Case2Snake,
	"Case2CamelLowerStart": jsonconv.Case2CamelLowerStart,
	"ToUpper":              strings.ToUpper,
	"ToLower":              strings.ToLower,
}
