package apidocs

import (
	"fmt"
	"path"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/server/tools/apidocs/apiparser"
)

func TestApiDocs(t *testing.T) {
	root := path.Join(files.GetRuntimeRoot(), "server/")

	cfg := Config{
		OutRoot:        "./tmp",
		ApiRoot:        []string{path.Join(root, "api/controller/logic")},
		ModelRoot:      []string{path.Join(root, "api/model"), path.Join(root, "infra/chatgpt/chat_model.go")},
		ApiBase:        "/api/v1",
		ImportPkgPaths: []string{`import http from "@/utils/request"`},
		IgnoredModels: map[string]string{
			"request.Context": "",
			//"request.EmptyReq":      "",
			"request.PageQuery":     "",
			"request.PageLimit":     "",
			"request.PageSort":      "",
			"request.PageCondition": "",
			//"request.IdReq":         "",
			//"request.IdsReq":        "",
			//"response.EmptyResp":    "",
			//"response.BatchResult":  "",
			"response.PageResult": "",
			"response.Response":   "",
		},
		ReplaceModels: map[string]string{
			"Response": "IApiResponseData",
		},
		ApiFuncNameAs: func(api *apiparser.ApiDeclare) string {
			return fmt.Sprintf("%vApi", jsonconv.Lcfirst(api.FunctionName))
		},
		ApiFieldNameAs: func(model *apiparser.ModelField) string {
			if model.JsonTag != "" {
				return model.JsonTag
			}

			return jsonconv.Camel2Case(model.Name)
		},
		ApiFieldTypeAs: func(name string) string {
			return convertGoTypeToTsType(name)
		},
	}

	aad := NewAstApiDoc(cfg)
	aad.Parse()
	// 生成ts api定义文件
	aad.GenerateTsApiFiles()
	// 生成ts type定义文件
	aad.GenerateTsTypeFile()
}

func TestExtractFieldsAfterDot(t *testing.T) {

	t.Log(jsonconv.ObjectToJsonIndent(extractResponseParams("response.Response{data=response.PageResult{list=[]entity.Api}}")))
}
