package quickstart

import (
	"fmt"
	"path"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/apidocs"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/quickstart/apidocs/apiparser"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

func TestApiDocs(t *testing.T) {
	root := path.Join(global.GetRuntimeRoot(), "server/")

	cfg := apidocs.Config{
		OutRoot:        "./api",
		ApiRoot:        []string{path.Join(root, "api/controller/logic")},
		ModelRoot:      []string{path.Join(root, "api/model"), path.Join(root, "infra/chatgpt/chat_model.go")},
		ApiBase:        "/api/v1",
		ImportPkgPaths: []string{`import http from "@/utils/request"`},
		IgnoredModels: []string{
			"response.PageResult", "response.Response", "request.PageQuery",
			"request.Context", "request.Sort", "request.Condition",
		},
		ReplaceModels: map[string]string{
			"Response": "IApiResponseData",
		},
		ApiFuncNameAs: func(api *apiparser.ApiDeclare) string {
			return fmt.Sprintf("%vApi", jsonconv.Lcfirst(api.FunctionName))
		},
		ApiFieldNameAs: func(field *apiparser.ModelField) string {
			return jsonconv.Camel2Case(field.Name)
		},
		ApiFieldTypeAs: func(field *apiparser.ModelField) string {
			return apidocs.GetTypeScriptType(field.Type)
		},
	}

	aad := apidocs.NewAstApiDoc(cfg)
	aad.Parse()
	aad.GenerateTsTypeFile()
	aad.GenerateTsApiFiles()
}

func TestExtractFieldsAfterDot(t *testing.T) {
	fmt.Println(apidocs.ExtractFieldsByAst(`response.Response{data=response.PageResult{list=[]entity.Article}}`))

}
