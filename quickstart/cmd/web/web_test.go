package web

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/quickstart/resource"
)

func Test_Typescript(t *testing.T) {
	f := &typescriptFlags{
		VarStringMode:     "api",
		VarStringFilePath: "/Users/weiyi/Github/ve-blog-golang/zero/service/api/blog/proto/admin.api",
		VarStringTplPath:  resource.GetTemplateRoot() + "/web",
		VarStringOutPath:  "./runtime/api",
		VarStringNameAs:   "",
	}
	RunTypescript(f)
}
