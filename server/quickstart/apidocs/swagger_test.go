package apidocs

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/quickstart/apidocs/apiparser"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

func TestSwaggerParser(t *testing.T) {
	ac := apiparser.NewSwaggerParser()
	apis, err := ac.ParseApiDocsByRoot(global.GetRuntimeRoot() + "server/docs/swagger.json")
	if err != nil {
		t.Error(err)
	}

	t.Log(jsonconv.ObjectToJsonIndent(apis))
}

func TestAstParser(t *testing.T) {
	ac := apiparser.NewAstParser()
	apis, err := ac.ParseApiDocsByRoot(global.GetRuntimeRoot() + "server/controllers")
	if err != nil {
		t.Error(err)
	}

	t.Log(jsonconv.ObjectToJsonIndent(apis))
}
