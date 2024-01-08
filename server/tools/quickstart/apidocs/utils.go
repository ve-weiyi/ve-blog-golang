package apidocs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
	"unicode"
)

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
	case "FileHeader":
		return "File"
	case "interface{}", "object":
		return "any"
	default:
		return name
	}
}

func ExtractFieldsByAst(data string) []string {
	if data == "" {
		return nil
	}
	// 使用正则表达式替换字符串中的 "="
	code := fmt.Sprintf("model:=%s", strings.ReplaceAll(data, "=", ":"))
	//meta := east.NewFuncMete("main", code)

	input := fmt.Sprintf(`
package main
func %v(){
   %v 
}`, "main", code)
	fParse, err := parser.ParseFile(token.NewFileSet(), "", input, 0)
	if err != nil {
		fmt.Println("Error parsing file:", err, code)
	}

	var params []string

	// CompositeLit
	nodes := ExtractNodes(&ast.CompositeLit{}, fParse)
	for _, node := range nodes {
		if len(params) > 0 {
			break
		}
		idents := ExtractNodes(&ast.Ident{}, node.Type)
		if len(idents) == 1 {
			params = append(params, idents[0].Name)
		} else {
			params = append(params, idents[0].Name+"."+idents[1].Name)
		}

	}

	// KeyValueExpr要value
	nodes2 := ExtractNodes(&ast.KeyValueExpr{}, fParse)
	for _, node := range nodes2 {
		switch fmt.Sprintf("%s", node.Key) {
		case "data":
			idents := ExtractNodes(&ast.Ident{}, node.Value)
			if len(idents) == 1 {
				params = append(params, idents[0].Name)
			} else if len(idents) > 1 {
				params = append(params, idents[0].Name+"."+idents[1].Name)
			} else {
				fmt.Println("cannot get params for:", idents)
			}
		case "list":
			idents := ExtractNodes(&ast.Ident{}, node.Value)
			if len(idents) == 1 {
				params = append(params, idents[0].Name)
			} else if len(idents) > 1 {
				params = append(params, idents[0].Name+"."+idents[1].Name)
			} else {
				fmt.Println("list cannot get params for:", idents)
			}
		}
	}

	fmt.Println("data:", data, "params:", params)
	return params
}

// 泛型方法 [T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
func ExtractNodes[T any](t T, node ast.Node) []T {
	var idents []T

	if n, ok := node.(T); ok {
		idents = append(idents, n)
	}

	switch n := node.(type) {
	case *ast.AssignStmt:
		return ExtractNodes(t, n.Rhs[0])
	case *ast.ArrayType:
		return ExtractNodes(t, n.Elt)
	case *ast.SelectorExpr:
		//fmt.Println("SelectorExpr", n.Sel.Name)
		idents = append(idents, ExtractNodes(t, n.X)...)
		idents = append(idents, ExtractNodes(t, n.Sel)...)
		break
	case *ast.KeyValueExpr:
		//fmt.Println("KeyValueExpr", n.Key)
		// 判断是否是复合字面值表达式的键值对
		idents = append(idents, ExtractNodes(t, n.Key)...)
		idents = append(idents, ExtractNodes(t, n.Value)...)
		break
	case *ast.CompositeLit:
		//fmt.Println("CompositeLit", n.Type)
		idents = append(idents, ExtractNodes(t, n.Type)...)
		for _, elt := range n.Elts {
			idents = append(idents, ExtractNodes(t, elt)...)
		}
		break
	case *ast.File:
		for _, decl := range n.Decls {
			idents = append(idents, ExtractNodes(t, decl)...)
		}

	case *ast.FuncDecl:
		for _, decl := range n.Body.List {
			idents = append(idents, ExtractNodes(t, decl)...)
		}

	case *ast.Ident:

	default:
		fmt.Printf("ExtractNodes default %T %T\n", t, node)
		ast.Print(token.NewFileSet(), node)
	}

	return idents
}

// response.Response{data=response.PageResult{list=[]entity.User}} --> Response、PageResult 和 User
func ExtractFieldsAfterDot(input string) []string {
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
