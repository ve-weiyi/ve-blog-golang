package apidocs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/east"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/files"
)

func ParseApiDocsByRoot(root string) []*ApiDeclare {
	apiDocs := make([]*ApiDeclare, 0)
	// 遍历目录下的所有文件
	files.VisitFile(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		// 是目录，则跳过
		if info.IsDir() {
			return nil
		}
		// 是文件，则判断是否是ctl.go文件
		if strings.HasSuffix(path, "ctl.go") {
			// 解析文件
			apiDocs = append(apiDocs, ParseApiDoc(path)...)
		}

		return nil
	})

	return apiDocs
}

func ParseApiDoc(fp string) []*ApiDeclare {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fp, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("解析代码时发生错误:", err)
		return nil
	}

	apiDocs := make([]*ApiDeclare, 0)

	for _, decl := range file.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			name := GetFunctionName(f)
			doc := GetFunctionDoc(f)

			api := &ApiDeclare{}

			api.FunctionName = name
			for _, comment := range doc {
				// 按照空白分割
				content := strings.Fields(comment.Content)
				if len(content) == 0 {
					continue
				}

				switch comment.Tag {
				case "Tags":
					// 按空白分割
					api.Tag = content[0]

				case "Summary":
					api.Summary = content[0]

				case "Param":
					tp := content[1]
					field := &ApiParam{
						Name: content[0],
						Type: content[2],
					}

					switch tp {
					case "header":
						api.Header = append(api.Header, field)
					case "path":
						api.Path = append(api.Path, field)
					case "query":
						api.Query = append(api.Query, field)
					case "formData":
						api.Form = append(api.Form, field)
					case "body":
						api.Body = field
					}

				case "Router":
					api.Url = content[0]
					api.Method = strings.TrimSuffix(strings.TrimPrefix(content[1], "["), "]")

				case "Success":
					api.Response = content[2]
				}
			}

			if api.Tag == "" {
				continue
			}
			apiDocs = append(apiDocs, api)
			// fmt.Println("函数注释:", jsonconv.ObjectToJsonIndent(api))
		}
	}

	return apiDocs
}

func ParseApiModelsByRoot(root string) []*ModelDeclare {
	var models []*ModelDeclare
	// 遍历目录下的所有文件
	files.VisitFile(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		// 是目录，则跳过
		if info.IsDir() {
			return nil
		}
		// 是文件，则判断是否是.go文件
		if strings.HasSuffix(path, ".go") {
			// 解析文件
			model := ParseApiModel(path)
			models = append(models, model...)
		}

		return nil
	})

	return models
}

func ParseApiModel(fp string) []*ModelDeclare {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fp, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("解析代码时发生错误:", err)
		return nil
	}

	var models []*ModelDeclare
	for _, decl := range file.Decls {
		if f, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range f.Specs {
				if t, ok := spec.(*ast.TypeSpec); ok {
					if s, ok := t.Type.(*ast.StructType); ok {
						var modelFields []*ModelField
						var extendFields []*ModelDeclare
						for _, field := range s.Fields.List {
							if len(field.Names) > 0 {
								modelFields = append(modelFields, extractModelField(field))
							} else {
								ext := extractModelField(field)
								extend := &ModelDeclare{
									Pkg:    ext.Name,
									Name:   ext.Type,
									Fields: nil,
								}

								extendFields = append(extendFields, extend)
							}
						}

						modelName := fmt.Sprintf("%s.%s", file.Name.Name, t.Name.Name)
						model := &ModelDeclare{
							Pkg:    file.Name.Name,
							Name:   modelName,
							Extend: extendFields,
							Fields: modelFields,
						}

						models = append(models, model)
					}
				}
			}
		}
	}

	return models
}

func getGoType(name string) string {
	if strings.HasPrefix(name, "*") {
		return GetTypeScriptType(name[1:]) // 指针
	}
	if strings.HasPrefix(name, "[]") {
		return GetTypeScriptType(name[2:]) // 数组
	}
	if strings.LastIndex(name, ".") > 0 {
		return GetTypeScriptType(name[strings.LastIndex(name, ".")+1:]) // 去掉包名
	}

	return name
}

func GetTypeScriptType(name string) string {
	if strings.HasPrefix(name, "*") {
		return GetTypeScriptType(name[1:]) // 指针
	}
	if strings.HasPrefix(name, "[]") {
		return GetTypeScriptType(name[2:]) + "[]" // 数组
	}
	if strings.LastIndex(name, ".") > 0 {
		return GetTypeScriptType(name[strings.LastIndex(name, ".")+1:]) // 去掉包名
	}
	switch name {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "file":
		return "File"
	case "Time":
		return "string"
	case "interface{}":
		return "any"

	default:
		return name
	}
}

func getModelDeclareName(method *ApiDeclare) []string {
	params := make([]string, 0)
	if method.Body != nil {
		params = append(params, method.Body.Type)
	}

	if len(method.Form) > 0 {
		for _, param := range method.Form {
			params = append(params, param.Type)
		}
	}

	if len(method.Path) > 0 {
		for _, param := range method.Path {
			params = append(params, param.Type)
		}
	}

	if len(method.Query) > 0 {
		for _, param := range method.Query {
			params = append(params, param.Type)
		}
	}

	if len(method.Response) > 0 {
		params = append(params, ExtractFieldsByAst(method.Response)...)
	}
	return params
}

func extractModelField(field ast.Node) *ModelField {

	switch node := field.(type) {
	case *ast.ArrayType:
		return extractModelField(node.Elt)
	case *ast.Field:
		if len(node.Names) > 0 {
			name := node.Names[0].Name
			tp := node.Type
			elem := &ModelField{
				Name: name,
				Type: GetNameFromExpr(tp),
			}

			// 读取字段的普通注释
			if node.Doc != nil {
				elem.Comment = strings.TrimSpace(node.Doc.Text())
			}

			// 读取字段的行尾注释
			if node.Comment != nil {
				elem.Comment = strings.TrimSpace(node.Comment.Text())
			}

			return elem
		}
		return extractModelField(node.Type)
	case *ast.StarExpr:
		return extractModelField(node.X)
	case *ast.SelectorExpr:
		if xIdent, ok := node.X.(*ast.Ident); ok {
			elem := &ModelField{
				Name: xIdent.Name,
				Type: node.Sel.Name,
			}

			// 读取字段的行尾注释
			//if field.Comment != nil {
			//	elem.Comment = strings.TrimSpace(field.Comment.Text())
			//}

			return elem
		}
	case *ast.Ident:
		elem := &ModelField{
			Name: node.Name,
			Type: node.Name,
		}
		// 读取字段的行尾注释
		//if node.Comment != nil {
		//	elem.Comment = strings.TrimSpace(field.Comment.Text())
		//}
		return elem
	case *ast.InterfaceType:
		return &ModelField{
			Name: "interface{}",
			Type: "any",
		}

	default:
		ast.Print(nil, field)
	}
	ast.Print(nil, field)

	return nil
}

// response.Response{data=response.PageResult{list=[]entity.Api}} --> Response、PageResult 和 Api
func extractFieldsAfterDot(input string) []string {
	// 定义正则表达式
	re := regexp.MustCompile(`\.(\w+)`)
	// 查找所有匹配的字符串
	matches := re.FindAllStringSubmatch(input, -1)

	// 提取 . 后面的字段并返回切片
	fields := make([]string, len(matches))
	for i, match := range matches {
		fields[i] = match[1]
	}

	return fields
}

func ExtractFieldsByAst(data string) []string {
	if data == "" {
		return nil
	}
	// 使用正则表达式替换字符串中的 "="
	code := fmt.Sprintf("model:=%s", strings.ReplaceAll(data, "=", ":"))
	meta := east.NewFuncMete("main", code)
	var params []string

	// CompositeLit
	nodes := ExtractNodes(meta.GetNode(), &ast.CompositeLit{})
	for _, node := range nodes {
		if len(params) > 0 {
			break
		}
		idents := ExtractNodes(node.Type, &ast.Ident{})
		if len(idents) == 1 {
			params = append(params, idents[0].Name)
		} else {
			params = append(params, idents[0].Name+"."+idents[1].Name)
		}

	}

	// KeyValueExpr要value
	nodes2 := ExtractNodes(meta.GetNode(), &ast.KeyValueExpr{})
	for _, node := range nodes2 {
		switch fmt.Sprintf("%s", node.Key) {
		case "data":
			idents := ExtractNodes(node.Value, &ast.Ident{})
			if len(idents) == 1 {
				params = append(params, idents[0].Name)
			} else {
				params = append(params, idents[0].Name+"."+idents[1].Name)
			}
		case "list":
			idents := ExtractNodes(node.Value, &ast.Ident{})
			if len(idents) == 1 {
				params = append(params, idents[0].Name)
			} else {
				params = append(params, idents[0].Name+"."+idents[1].Name)
			}
		}
	}

	return params
}

func replaceEquals(input string) string {
	// 按照字符等号 "=" 分割字符串
	parts := strings.Split(input, "=")

	// 遍历分割后的部分，并根据条件重新组合字符串
	for i := 1; i < len(parts); i++ {
		// 判断前一个部分是否以冒号 ":" 结尾，如果是则保留等号 "="
		if strings.HasSuffix(parts[i-1], ":") {
			parts[i] = "=" + parts[i]
		} else {
			parts[i] = ":" + parts[i]
		}
	}

	// 重新组合字符串
	output := strings.Join(parts, "")

	return output
}

func getIdentDeclareName(name string) string {
	var englishChars []rune
	if strings.LastIndex(name, ".") > 0 {
		name = name[strings.LastIndex(name, ".")+1:]
	}

	for _, char := range name {
		if unicode.IsLetter(char) {
			englishChars = append(englishChars, char)
		}
	}
	return string(englishChars)
}
