package apidocs

import (
	"log"
	"path"
	"testing"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/inject"
)

func TestSwagger(t *testing.T) {
	converter := &SwaggerApiCollector{}

	converter.ReadSwagJSON(global.GetRuntimeRoot() + "/server/docs/swagger.json")

	log.Println(jsonconv.ObjectToJsonIndent(converter.GetApiTs()))

	converter.toTypeScriptApis("./api", converter.GetApiTs())
}

func TestDst(t *testing.T) {
	root := path.Join(global.GetRuntimeRoot(), "server/api/", "blog")

	cfg := Config{
		OutRoot:        "",
		ApiRoot:        []string{path.Join(root, "controller/logic")},
		ModelRoot:      []string{path.Join(root, "model")},
		ImportPkgPaths: []string{`import http from "@/utils/request"`},
	}

	aad := NewAstApiDoc(cfg)
	aad.Parse()
	aad.GenerateTypeTsFile()
	aad.GenerateApiTsFiles()
}

func TestExtractFieldsAfterDot(t *testing.T) {
	code := `model:=response.Response{data:response.PageResult{list:[]entity.Article}}`

	meta := inject.NewFuncMete("main", code)
	meta.GetNode()
	log.Println("ExtractIdents", jsonconv.ObjectToJsonIndent(inject.ExtractIdents(meta.GetNode())))
	log.Println("ExtractSelectors", jsonconv.ObjectToJsonIndent(inject.ExtractSelectors(meta.GetNode())))
}
