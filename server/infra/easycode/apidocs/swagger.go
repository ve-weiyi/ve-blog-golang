package apidocs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

type SwaggerApiCollector struct {
	swagger *SwaggerDefinition
}

func (s *SwaggerApiCollector) ReadSwagJSON(filepath string) error {
	// 读取 JSON 文件内容
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Failed to read JSON file:", err)
		return nil
	}

	// 解析 JSON 数据到结构体
	var swagger SwaggerDefinition
	err = json.Unmarshal(jsonData, &swagger)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return nil
	}

	s.swagger = &swagger

	return nil
}

func (s *SwaggerApiCollector) GetApiTs() map[string]*ApiDoc {
	swagger := s.swagger
	// 生成 TypeScript 代码
	apis := make(map[string]*ApiDoc)
	for path, paths := range swagger.Paths {
		for method, pathItem := range paths {
			if pathItem.Tags == nil {
				continue
			}

			tag := jsonconv.Camel2Case(pathItem.Tags[0])
			if apis[tag] == nil {
				apis[tag] = &ApiDoc{
					Tag:      tag,
					Function: make(map[string]*ApiDeclare),
				}
			}

			apiPath := swagger.BasePath + path

			// var params []string
			// for _, param := range pathItem.Parameters {
			// 	switch param.In {
			// 	case "body":
			// 		body = param.Name
			// 	case "query":
			// 		params = append(params, param.Name)
			// 	case "path":
			// 		params = append(params, param.Name)
			// 	default:
			// 	}
			// }

			apiTs := &ApiDeclare{
				Method: method,
				Url:    apiPath,
				// Body:   body,
				// Params:      params,
				// Description: pathItem.Summary,
			}

			file := apis[tag]
			funcName := getFuncName(path)
			file.Function[funcName] = apiTs
		}

	}

	return apis
}

func (s *SwaggerApiCollector) toTypeScriptApis(root string, apis map[string]*ApiDoc) {

	metas := make([]plate.PlateMeta, 0)
	for _, api := range apis {

		meta := plate.PlateMeta{
			Key:            "api",
			AutoCodePath:   fmt.Sprintf("%s/%s.ts", root, api.Tag),
			Replace:        true,
			TemplateString: ApiTypeScript,
			Data:           api,
		}

		metas = append(metas, meta)
		// fmt.Println(jsonconv.ObjectToJsonIndent(api))
	}

	for _, meta := range metas {
		err := meta.CreateTempFile()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func getParamName(in string) string {
	switch in {
	case "query":
		return "params"
	case "path":
		return "params"
	case "body":
		return "data"
	case "formData":
		return "data"
	case "header":
		return "headers"
	default:
		return "params"
	}
}

func getFuncName(path string) string {
	var name string
	name = strings.ReplaceAll(path, "/", "_")

	// var key []string
	// key = strings.Split(path, "/")
	// for i := len(key) - 1; i >= 0; i-- {
	//	name = name + "_" + key[i]
	// }

	return jsonconv.Case2CamelNotFirst(fmt.Sprintf("%s%s", name, "Api"))
}
