package apiparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type AstParser struct {
}

func NewAstParser() ApiParser {
	return &AstParser{}
}

func (s *AstParser) ParseApiDocsByRoot(root ...string) (out []*ApiDeclare, err error) {
	out = make([]*ApiDeclare, 0)
	for _, v := range root {
		swagger, err := s.ReadApiDocsComment(v)
		if err != nil {
			return nil, err
		}

		out = append(out, swagger...)
	}

	return out, nil
}

func (s *AstParser) ReadApiDocsComment(root string) (out []*ApiDeclare, err error) {
	apiDocs := make([]*ApiDeclare, 0)
	// 遍历目录下的所有文件
	VisitFile(root, func(path string, info os.FileInfo, err error) error {
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

	return apiDocs, nil
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
					api.Router = content[0]
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
