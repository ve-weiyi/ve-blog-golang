package apidocs

import (
	"encoding/json"
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-admin-store/server/infra/easycode/plate"
	"os"
	"strings"
)

type SwaggerConverter struct {
	swagger *SwaggerDefinition
}

func (s *SwaggerConverter) ReadSwagJSON(filepath string) error {
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

func (s *SwaggerConverter) toTypeScriptApis(outpath string) {
	swagger := s.swagger
	// 生成 TypeScript 代码
	//var apiFiles map[string][]PathItem
	apis := make(map[string]*ApiTs)
	for path, paths := range swagger.Paths {
		for method, pathItem := range paths {
			if pathItem.Tags == nil {
				continue
			}

			tag := jsonconv.Camel2Case(pathItem.Tags[0])
			if apis[tag] == nil {
				apis[tag] = &ApiTs{
					FileName: tag,
					Function: make(map[string]*ApiTsMethod),
				}
			}

			apiPath := swagger.BasePath + path

			var body string
			var params []string
			for _, param := range pathItem.Parameters {
				switch param.In {
				case "body":
					body = param.Name
				case "query":
					params = append(params, param.Name)
				case "path":
					params = append(params, param.Name)
				default:
				}
			}

			apiTs := &ApiTsMethod{
				Method:      method,
				Url:         apiPath,
				Body:        body,
				Params:      params,
				Description: pathItem.Summary,
			}

			file := apis[tag]
			funcName := getFuncName(path)
			file.Function[funcName] = apiTs
		}

	}

	metas := make([]plate.PlateMeta, 0)
	for _, api := range apis {

		meta := plate.PlateMeta{
			Key:            "api",
			AutoCodePath:   fmt.Sprintf("%s/%s.ts", outpath, api.FileName),
			Replace:        true,
			TemplateString: ApiTypeScript,
			Data:           api,
		}

		metas = append(metas, meta)
		//fmt.Println(jsonconv.ObjectToJsonIndent(api))
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

	//var key []string
	//key = strings.Split(path, "/")
	//for i := len(key) - 1; i >= 0; i-- {
	//	name = name + "_" + key[i]
	//}

	return jsonconv.Case2CamelNotFirst(fmt.Sprintf("%s%s", name, "Api"))
}
