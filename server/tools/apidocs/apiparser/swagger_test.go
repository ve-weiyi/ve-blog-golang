package apiparser

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func TestSwaggerParser(t *testing.T) {
	ac := NewSwaggerParser()
	apis, err := ac.ParseApiDocsByRoots(files.GetRuntimeRoot() + "server/docs")
	if err != nil {
		t.Error(err)
	}

	t.Log(jsonconv.AnyToJsonIndent(apis))

	models, err := ac.ParseModelDocsByRoots(files.GetRuntimeRoot() + "server/docs")
	if err != nil {
		t.Error(err)
	}

	t.Log(jsonconv.AnyToJsonIndent(models))
}

func TestAstParser(t *testing.T) {
	cfg := AstParserConfig{
		ApiBase: "/api/v1",
	}
	ac := NewAstParser(cfg)
	//apis, err := ac.ParseApiDocsByRoots(global.GetRuntimeRoot() + "server/api/blog/controller")
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//t.Log(jsonconv.AnyToJsonIndent(apis))

	models, err := ac.ParseModelDocsByRoots(files.GetRuntimeRoot() + "server/api/blog/model")
	if err != nil {
		t.Error(err)
	}

	t.Log(jsonconv.AnyToJsonIndent(models))
}
