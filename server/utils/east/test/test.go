package test

import (
	"go/ast"

	jsoniter "github.com/json-iterator/go"
)

type ApiGroup struct {
	visitor ast.Visitor  //
	jjson   jsoniter.API //注释
}

type SvGroup struct {
}

func (receiver *ApiGroup) GetName() string {
	//11111

	//注释
	{
	}
	//11112
	return ""
}
