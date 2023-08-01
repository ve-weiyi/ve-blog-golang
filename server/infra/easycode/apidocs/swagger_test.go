package apidocs

import (
	"fmt"
	"log"
	"path"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
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
		OutRoot:        "./api",
		ApiRoot:        []string{path.Join(root, "controller/logic")},
		ModelRoot:      []string{path.Join(root, "model")},
		ImportPkgPaths: []string{`import http from "@/utils/request"`},
		IgnoredModels:  []string{"response.PageResult", "response.Response", "request.PageQuery"},
		ReplaceModels: map[string]string{
			"Response": "IApiResponseData",
		},
		ApiFuncNameAs: func(api *ApiDeclare) string {
			return fmt.Sprintf("%vApi", jsonconv.Lcfirst(api.FunctionName))
		},
		ApiFieldNameAs: func(field *ModelField) string {
			return jsonconv.Camel2Case(field.Name)
		},
		ApiFieldTypeAs: func(field *ModelField) string {
			return getTypeScriptType(field.Type)
		},
	}

	aad := NewAstApiDoc(cfg)
	aad.Parse()
	//aad.GenerateTsTypeFile()
	aad.GenerateTsApiFiles()
}

func TestExtractFieldsAfterDot(t *testing.T) {
	fmt.Println(extractFieldsByAst(`response.Response{data=entity.Api}`))

}
