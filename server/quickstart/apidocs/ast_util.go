package apidocs

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

// 获取函数名称
func GetFunctionName(f *ast.FuncDecl) string {
	return f.Name.Name
}

// 获取函数注释
func GetFunctionDoc(f *ast.FuncDecl) []ApiCommentLine {
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
func GetNameFromExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// 处理标识符，直接返回标识符的名称
		return t.Name
	case *ast.SelectorExpr:
		// 处理选择器表达式，格式化输出包名和选择器名称
		return fmt.Sprintf("%v.%v", GetNameFromExpr(t.X), t.Sel.Name)
	case *ast.ArrayType:
		// 处理数组类型，格式化输出数组的元素类型名称
		return fmt.Sprintf("[]%v", GetNameFromExpr(t.Elt))
	case *ast.StarExpr:
		// 处理指针类型，格式化输出指针指向的类型名称
		return fmt.Sprintf("*%v", GetNameFromExpr(t.X))
	case *ast.InterfaceType:
		// 处理接口类型，直接返回 "interface{}"
		return "interface{}"
	default:
		// 其他未处理的表达式类型，返回空字符串
		return ""
	}
}

// 泛型方法 [T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
func ExtractNodes[T any](node ast.Node, t T) []T {
	var idents []T

	if n, ok := node.(T); ok {
		idents = append(idents, n)
	}

	switch n := node.(type) {
	case *ast.AssignStmt:
		return ExtractNodes(n.Rhs[0], t)
	case *ast.ArrayType:
		return ExtractNodes(n.Elt, t)
	case *ast.SelectorExpr:
		//fmt.Println("SelectorExpr", n.Sel.Name)
		idents = append(idents, ExtractNodes(n.X, t)...)
		idents = append(idents, ExtractNodes(n.Sel, t)...)
	case *ast.KeyValueExpr:
		//fmt.Println("KeyValueExpr", n.Key)
		// 判断是否是复合字面值表达式的键值对
		idents = append(idents, ExtractNodes(n.Key, t)...)
		idents = append(idents, ExtractNodes(n.Value, t)...)
	case *ast.CompositeLit:
		//fmt.Println("CompositeLit", n.Type)
		idents = append(idents, ExtractNodes(n.Type, t)...)
		for _, elt := range n.Elts {
			idents = append(idents, ExtractNodes(elt, t)...)
		}
	default:
		//fmt.Printf("default %T\n", n)
	}

	return idents
}
