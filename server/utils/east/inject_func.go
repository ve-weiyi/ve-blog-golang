package east

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strconv"

	"github.com/spf13/cast"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

const (
	name = token.DEFINE
)

// Variables := IdentName.SelName(Parameters...)
// IdentName.SelName(Parameters...)
type FuncMeta struct {
	fset       *token.FileSet
	InjectCode string //注入的代码
	FuncName   string //插入的方法
	FuncPos    int    //插入的位置

	//ValueName string //声明
	//IdentName string //引用名
	//SelName   string //元素struct名

	Variables  []string
	Symbol     string
	IdentNames []string
	Parameters []interface{}
}

func (vi FuncMeta) Visit(node ast.Node) ast.Visitor {
	switch fn := node.(type) {
	case *ast.FuncDecl:
		log.Printf("Func: %s %s\n", fn.Name.Name, vi.FuncName)
		if fn.Name.Name == vi.FuncName {

			hasVar := false
			for _, v := range fn.Body.List {
				switch varSpec := v.(type) {
				case *ast.AssignStmt: //赋值表达式
					for i := range varSpec.Lhs {
						switch nn := varSpec.Lhs[i].(type) {
						case *ast.Ident:
							if nn.Name == vi.Variables[i] {
								hasVar = true
							}
						}
					}

				case *ast.ExprStmt: //表达语句
					switch expr := varSpec.X.(type) {
					case *ast.CallExpr: //是调用语句
						if call, ok := expr.Fun.(*ast.SelectorExpr); ok {
							x, _ := call.X.(*ast.Ident)
							// 包名和方法名都相等
							if x.Name == vi.IdentNames[0] {
								hasVar = true
							}
						}
					}
				case *ast.IfStmt:
				case *ast.GoStmt:
				case *ast.ReturnStmt:
					break
				}
			}

			var insertStmt ast.Stmt
			//lline := vi.fset.Position(fn.Body.Lbrace).Line
			//rline := vi.fset.Position(fn.Body.Rbrace).Line
			if !hasVar && vi.Symbol == ":=" {
				var rhs ast.Expr
				if len(vi.Parameters) == 0 {
					//是变量 *ast.SelectorExpr
					rhs = vi.GetSelector()
				} else {
					//是调用
					rhs = &ast.CallExpr{
						Fun:  vi.GetSelector(),
						Args: vi.GetParameters(),
					}
				}
				//变量声明
				insertStmt = &ast.AssignStmt{
					// 等式左边
					Lhs: []ast.Expr{
						ast.NewIdent(vi.IdentNames[0]),
					},
					// := 符号
					Tok: token.DEFINE,
					// 等式右边
					Rhs: []ast.Expr{
						rhs,
					},
				}
			}

			if !hasVar && vi.Symbol == "" {
				// 调用语句
				insertStmt = &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun:      vi.GetSelector(),
						Lparen:   0,
						Args:     vi.GetParameters(),
						Ellipsis: 0,
						Rparen:   0,
					},
				}

			}

			//放到第pos位置
			index := vi.FuncPos
			if vi.FuncPos >= len(fn.Body.List)-1 {
				index = len(fn.Body.List) - 1
			}
			// Find the line number after the third line
			for i, stmt := range fn.Body.List {
				ast.Print(vi.fset, fn)
				pos := index
				log.Println("line ", stmt)
				log.Println("line ", vi.fset.Position(stmt.Pos()).Line)
				log.Println("line ", pos)
				log.Println("line ", vi.fset.Position(fn.Body.Lbrace).Line)
				log.Println("line ", vi.fset.Position(fn.Body.Rbrace).Line)
				if vi.fset.Position(stmt.Pos()).Line > pos {
					fn.Body.List = insertStatements(fn.Body.List, i, insertStmt)
					break
				}
			}
			////放到第pos位置
			//fn.Body.List = append(fn.Body.List, fn.Body.List[index])
			//copy(fn.Body.List[index+1:], fn.Body.List[index:])
			//fn.Body.List[index] = assignStmt
		}
	}
	return vi
}

func (vi FuncMeta) GetSelector() ast.Expr {
	var selector interface{}
	if len(vi.IdentNames) == 1 {
		//var selector *ast.Ident
		selector = ast.NewIdent(vi.IdentNames[0])
	} else {
		// >=2
		selector = &ast.SelectorExpr{
			// 只有一个 .
			X: &ast.Ident{
				Name: cast.ToString(vi.IdentNames[0]),
			},
			// IdentName.SelName
			Sel: &ast.Ident{
				Name: cast.ToString(vi.IdentNames[1]),
			},
		}
		for _, value := range vi.IdentNames[2:] {
			selector = &ast.SelectorExpr{
				X: selector.(*ast.SelectorExpr),
				//IdentName.SelName.SelName
				Sel: ast.NewIdent(cast.ToString(value)),
			}
		}
	}
	return selector.(ast.Expr)
}

func (vi FuncMeta) GetParameters() []ast.Expr {
	var varExpr []ast.Expr
	for _, value := range vi.Parameters {
		exp := &ast.BasicLit{
			Kind:  kindToToken(reflect.TypeOf(value).Kind()),
			Value: jsonconv.ObjectToJson(value),
		}
		varExpr = append(varExpr, exp)
	}
	return varExpr
}

func (vi FuncMeta) RollBack(node ast.Node) ast.Visitor {
	if funDecl, ok := node.(*ast.FuncDecl); ok {
		if funDecl.Name.Name == vi.FuncName {
			for k, v := range funDecl.Body.List {
				switch varSpec := v.(type) {
				//赋值表达式
				case *ast.AssignStmt:
					for i := range varSpec.Lhs {
						switch nn := varSpec.Lhs[i].(type) {
						case *ast.Ident:
							if nn.Name == vi.IdentNames[0] {
								//删除
								block := funDecl.Body
								block.List = append(append([]ast.Stmt{}, block.List[:k]...), block.List[k+1:]...)
								return vi
							}
						}
					}
				case *ast.ExprStmt: //表达语句
					switch expr := varSpec.X.(type) {
					case *ast.CallExpr: //是调用语句
						if call, ok := expr.Fun.(*ast.SelectorExpr); ok {
							x, _ := call.X.(*ast.Ident)
							// 包名和方法名都相等
							if x.Name == vi.IdentNames[0] {
								//删除
								block := funDecl.Body
								block.List = append(append([]ast.Stmt{}, block.List[:k]...), block.List[k+1:]...)
								return vi
							}
						}
					}
				}
			}
		}
	}
	return vi
}

func NewFuncMete(importCode string) {
	meta := FuncMeta{}
	input := "package main\nfunc main() {\n" + importCode + "\n}"
	fParse, err := parser.ParseFile(token.NewFileSet(), "", input, 0)
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}
	ast.Inspect(fParse, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			ast.Print(token.NewFileSet(), x)

			variables := []string{}
			for _, lhs := range x.Lhs {
				variables = append(variables, lhs.(*ast.Ident).Name)
			}

			indentNames := []string{}
			parameters := []interface{}{}
			switch callExpr := x.Rhs[0].(type) {
			case *ast.CallExpr:
				idents := extractIdents(callExpr.Fun)
				for _, ident := range idents {
					indentNames = append(indentNames, ident.Name)
				}
				for _, arg := range callExpr.Args {
					switch arg.(type) {
					case *ast.BasicLit:
						value, _ := InferType(arg.(*ast.BasicLit).Value)
						parameters = append(parameters, value)
					case *ast.Ident:
						parameters = append(parameters, arg.(*ast.Ident).Name)
					}
				}
			case *ast.SelectorExpr:
				idents := extractIdents(callExpr)
				for _, ident := range idents {
					indentNames = append(indentNames, ident.Name)
				}
			}

			meta.Variables = variables
			meta.Symbol = x.Tok.String()
			meta.IdentNames = indentNames
			meta.Parameters = parameters

		case *ast.ExprStmt:
			ast.Print(token.NewFileSet(), x)

			indentNames := []string{}
			callExpr := x.X.(*ast.CallExpr)
			idents := extractIdents(callExpr.Fun)
			for _, ident := range idents {
				indentNames = append(indentNames, ident.Name)
			}

			parameters := []interface{}{}
			for _, arg := range callExpr.Args {
				switch arg.(type) {
				case *ast.BasicLit:
					value, _ := InferType(arg.(*ast.BasicLit).Value)
					parameters = append(parameters, value)
				case *ast.Ident:
					parameters = append(parameters, arg.(*ast.Ident).Name)
				}
			}

			meta.IdentNames = indentNames
			meta.Parameters = parameters
		}
		return true
	})

	log.Println("InjectCode", meta.InjectCode)
	log.Println("NewStructMete", jsonconv.ObjectToJsonIndent(meta))

}

func parseStatements(fset *token.FileSet, src string) ([]ast.Stmt, error) {
	node, err := parser.ParseFile(fset, "", "package main; func _() {"+src+"}", 0)
	if err != nil {
		return nil, err
	}

	fn := node.Decls[0].(*ast.FuncDecl)
	return fn.Body.List, nil
}

func extractIdents(node ast.Node) []*ast.Ident {
	var idents []*ast.Ident

	switch n := node.(type) {
	case *ast.SelectorExpr:
		idents = append(idents, extractIdents(n.X)...)
		idents = append(idents, extractIdents(n.Sel)...)
	case *ast.Ident:
		idents = append(idents, n)
	}

	return idents
}

func insertStatements(stmts []ast.Stmt, pos int, toInsert ...ast.Stmt) []ast.Stmt {
	return append(stmts[:pos], append(toInsert, stmts[pos:]...)...)
}

func InferType(str string) (interface{}, error) {
	// 尝试将字符串解析为int
	i, err := strconv.Atoi(str)
	if err == nil {
		return i, nil
	}

	// 尝试将字符串解析为float
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return f, nil
	}

	// 尝试将字符串解析为带引号的string
	s, err := strconv.Unquote(str)
	if err == nil {
		return s, nil
	}

	// 如果都不匹配，则返回原始字符串
	return str, nil
}
