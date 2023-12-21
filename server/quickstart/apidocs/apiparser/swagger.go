package apiparser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SwaggerParser struct {
}

func NewSwaggerParser() ApiParser {
	return &SwaggerParser{}
}

func (s *SwaggerParser) ParseApiDocsByRoot(root ...string) (out []*ApiDeclare, err error) {
	out = make([]*ApiDeclare, 0)
	for _, v := range root {
		// 遍历目录下的所有文件
		VisitFile(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			// 是目录，则跳过
			if info.IsDir() {
				return nil
			}
			// 是文件，则判断是否是ctl.go文件
			if strings.HasSuffix(path, ".json") {
				// 解析文件
				swagger, err := s.ReadSwagJSON(v)
				if err != nil {
					return err
				}

				out = append(out, s.GetApiDeclares(swagger)...)
			}
			return nil
		})
	}

	return out, nil
}

func (s *SwaggerParser) ReadSwagJSON(filepath string) (out *SwaggerDefinition, err error) {

	// 读取 JSON 文件内容
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file:%v", err)
	}

	// 解析 JSON 数据到结构体
	var swagger SwaggerDefinition
	err = json.Unmarshal(jsonData, &swagger)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON:%v", err)
	}

	//fmt.Println("swagger:", jsonconv.ObjectToJsonIndent(swagger))

	return &swagger, nil
}

// 提取api列表
func (s *SwaggerParser) GetApiDeclares(swagger *SwaggerDefinition) []*ApiDeclare {
	// 生成 TypeScript 代码
	apis := make([]*ApiDeclare, 0)
	for path, paths := range swagger.Paths {

		//fmt.Println("path:", path)

		for method, apiMethod := range paths {

			//fmt.Println("method:", method, "apiMethod:", jsonconv.ObjectToJsonIndent(apiMethod))

			tag := JoinArray(apiMethod.Tags)

			// 未创建tag
			apiPath := swagger.BasePath + path

			var body *ApiParam
			var header []*ApiParam
			var query []*ApiParam
			var params []*ApiParam
			var form []*ApiParam
			for _, param := range apiMethod.Parameters {

				//fmt.Println("param:", jsonconv.ObjectToJsonIndent(param))
				t := getTypeNameFormSchema(param.SchemaObject)

				switch param.In {
				case "header":
					p := &ApiParam{
						Name: param.Name,
						Type: t,
					}
					header = append(header, p)
				case "query":
					p := &ApiParam{
						Name: param.Name,
						Type: t,
					}
					query = append(query, p)
				case "path":
					p := &ApiParam{
						Name: param.Name,
						Type: t,
					}
					params = append(params, p)
				case "form":
					p := &ApiParam{
						Name: param.Name,
						Type: t,
					}
					form = append(form, p)
				case "body":
					body = &ApiParam{
						Name: param.Name,
						Type: getTypeNameFormSchema(param.Schema),
					}
				default:
				}
			}

			var response string
			for _, v := range apiMethod.Responses {
				// 200
				//fmt.Println("in:", jsonconv.ObjectToJsonIndent(v.Schema))
				response = getTypeNameFormSchema(v.Schema)
				//fmt.Println("out:", jsonconv.ObjectToJsonIndent(response))
				break
			}

			apiTs := &ApiDeclare{
				Tag:          tag,
				FunctionName: getFuncNameFormPath(path),
				Summary:      apiMethod.Summary,
				Router:       apiPath,
				Method:       method,
				Header:       header,
				Path:         params,
				Query:        query,
				Form:         form,
				Body:         body,
				Response:     response,
			}

			apis = append(apis, apiTs)
		}

	}

	return apis
}

func getFuncNameFormPath(path string) string {
	if path == "/" || path == "" {
		return ""
	}

	var name string
	name = strings.ReplaceAll(path, "{", "_")
	name = strings.ReplaceAll(path, "}", "_")
	name = strings.ReplaceAll(path, "/", "_")

	var newName string
	var flag bool = false
	for i := 0; i < len(name); i++ {
		if name[i] == '_' {
			flag = true
			continue
		}

		if flag {
			newName = newName + strings.ToUpper(string(byte(rune(name[i]))))
			flag = false
		} else {
			newName = newName + string(byte(rune(name[i])))
		}
	}

	return strings.ToLower(string(byte(rune(newName[0])))) + newName[1:]
}

func getTypeNameFormSchema(in *SchemaObject) string {
	if in == nil {
		return ""
	}

	var t string
	if in.Items != nil {
		t += getTypeNameFormItems(in.Items)
	}

	if in.AllOf != nil {
		for _, v := range in.AllOf {
			t += v.Ref
			for _, p := range v.Properties {
				t += getTypeNameFormSchema(&p)
			}
		}
	}

	return t
}

func getTypeNameFormItems(in *Items) string {
	var t string
	if in.Ref != "" {
		t += in.Ref
	}

	if in.Items != nil {
		t += getTypeNameFormItems(in.Items)
	}

	if t == "" {
		return "#" + in.Type
	}

	return t
}
