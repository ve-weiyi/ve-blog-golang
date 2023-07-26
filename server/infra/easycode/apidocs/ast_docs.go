package apidocs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/ve-weiyi/go-sdk/utils/files"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/tmpl"
)

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
			name := getFunctionName(f)
			doc := getFunctionDoc(f)

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
					case "form":
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
						for _, field := range s.Fields.List {
							if len(field.Names) > 0 {
								name := field.Names[0].Name
								tp := field.Type
								elem := &ModelField{
									Name: name,
									Type: getNameFromExpr(tp),
								}

								// 读取字段的普通注释
								//if field.Doc != nil {
								//	elem.Comment = strings.TrimSpace(field.Doc.Text())
								//}

								// 读取字段的行尾注释
								if field.Comment != nil {
									elem.Comment = strings.TrimSpace(field.Comment.Text())
								}

								modelFields = append(modelFields, elem)
								//fmt.Println("name:", name, "tp:", tp)
							}
						}

						// modelName := fmt.Sprintf("%s.%s", file.Name.Name, t.Name.Name)
						model := &ModelDeclare{
							Pkg:    file.Name.Name,
							Name:   t.Name.Name,
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

// 收集api需要的导入的model（去重）
func CollectApiModel(apis []*ApiDeclare) []string {
	params := make(map[string]string)

	for _, method := range apis {
		if method.Body != nil {
			params[method.Body.Type] = method.Body.Type
		}

		if len(method.Form) > 0 {
			for _, param := range method.Form {
				params[param.Type] = param.Type
			}
		}

		if len(method.Path) > 0 {
			for _, param := range method.Path {
				params[param.Type] = param.Type
			}
		}

		if len(method.Query) > 0 {
			for _, param := range method.Query {
				params[param.Type] = param.Type
			}
		}
	}

	var out []string
	for _, param := range params {
		out = append(out, param)
	}

	return out
}

func GenerateApiTypeScript(root string, apis map[string][]*ApiDeclare) {

	metas := make([]plate.PlateMeta, 0)
	for tag, api := range apis {

		meta := plate.PlateMeta{
			Key:            "api",
			AutoCodePath:   fmt.Sprintf("%s/%s.ts", root, tag),
			Replace:        true,
			TemplateString: tmpl.Api,
			Data:           api,
		}

		metas = append(metas, meta)
	}

	for _, meta := range metas {
		err := meta.CreateTempFile()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func convertTsModelFields(fields []*ModelField) []*ModelField {
	tsFields := make([]*ModelField, 0)
	for _, field := range fields {
		tsField := &ModelField{
			Name:    jsonconv.Camel2Case(field.Name),
			Type:    getTypeScriptType(field.Type),
			Comment: field.Comment,
		}

		tsFields = append(tsFields, tsField)
	}

	return tsFields
}

// 获取函数名称
func getFunctionName(f *ast.FuncDecl) string {
	return f.Name.Name
}

// 获取函数注释
func getFunctionDoc(f *ast.FuncDecl) []ApiCommentLine {
	if f.Doc != nil {
		// return f.Doc.Text()
		// 定义一个map来存储注释
		comments := make([]ApiCommentLine, 0)
		// 定义正则表达式，匹配形式为 @标签名 内容
		re := regexp.MustCompile(`@(\w+)\s+(.+?)\s*$`)
		// 获取函数节点上方的注释
		for _, comment := range f.Doc.List {
			commentText := strings.TrimSpace(comment.Text)
			// 使用正则表达式提取注释中的标签和内容
			matches := re.FindAllStringSubmatch(commentText, -1)
			for _, match := range matches {
				tag := strings.TrimSpace(match[1])
				comments = append(comments, ApiCommentLine{Tag: tag, Content: strings.TrimSpace(match[2])})
			}
		}
		return comments
	}

	return nil
}

// 从 ast.Expr 中取出名称。
func getNameFromExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// 处理标识符，直接返回标识符的名称
		return t.Name
	case *ast.SelectorExpr:
		// 处理选择器表达式，格式化输出包名和选择器名称
		return fmt.Sprintf("%v.%v", getNameFromExpr(t.X), t.Sel.Name)
	case *ast.ArrayType:
		// 处理数组类型，格式化输出数组的元素类型名称
		return fmt.Sprintf("[]%v", getNameFromExpr(t.Elt))
	case *ast.StarExpr:
		// 处理指针类型，格式化输出指针指向的类型名称
		return fmt.Sprintf("*%v", getNameFromExpr(t.X))
	case *ast.InterfaceType:
		// 处理接口类型，直接返回 "interface{}"
		return "interface{}"
	default:
		// 其他未处理的表达式类型，返回空字符串
		return ""
	}
}

func getTypeScriptType(name string) string {
	switch name {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "time.Time":
		return "string"
	default:
		return "any"
	}

}
