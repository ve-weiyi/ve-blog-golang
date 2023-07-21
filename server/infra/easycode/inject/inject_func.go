package inject

import (
	"fmt"
	"go/token"
	"log"
	"reflect"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
)

const (
	name = token.DEFINE
)

// Variables := IdentName.SelName(Parameters...)
// IdentName.SelName(Parameters...)
type FuncMeta struct {
	FuncName string //插入的方法
	FuncPos  int    //插入的位置

	Variables  []string      //左边的值
	Symbol     string        //符号
	IdentNames []string      //右边的值，当为return时，是右边变量名
	Parameters []interface{} //调用参数
	Comments   string        //注释

	fset *token.FileSet
	node dst.Stmt
}

func (vi *FuncMeta) Visit(node dst.Node) dst.Visitor {
	switch fn := node.(type) {
	case *dst.FuncDecl:
		if fn.Name.Name == vi.FuncName {
			//log.Printf("Func: %s %s\n", fn.Name.Name, vi.FuncName)
			for _, v := range fn.Body.List {
				switch stmt := v.(type) {
				case *dst.AssignStmt: //赋值表达式
					if vi.Symbol != ":=" {
						break
					}
					if len(vi.Variables) == len(stmt.Lhs) {
						for i, lhs := range stmt.Lhs {
							switch nn := lhs.(type) {
							case *dst.Ident:
								if nn.Name == vi.Variables[i] {
									log.Println("hasVar", nn.Name, vi.Variables[i])
									return vi
								}
							}
						}
					}

				case *dst.ExprStmt: //表达语句
					if vi.Symbol != "" {
						break
					}
					switch expr := stmt.X.(type) {
					case *dst.CallExpr: //是调用语句，调用语句不查重
						if call, ok := expr.Fun.(*dst.SelectorExpr); ok {
							_, _ = call.X.(*dst.Ident)
							// 包名xd和方法名都相等
						}
					}
				case *dst.BlockStmt: //空块，注释的效果等于方法名
					if vi.Symbol != "{}" {
						break
					}
					if len(stmt.Decs.Start.All()) != 0 {

					}
				case *dst.IfStmt:
				case *dst.GoStmt:
				case *dst.ReturnStmt:
					if vi.Symbol != "return" {
						break
					}
					break
				}
			}

			insertStmt := vi.GetNode()

			//放到第pos位置
			index := vi.FuncPos
			if vi.FuncPos >= len(fn.Body.List)-1 {
				//除了return语句
				index = len(fn.Body.List) - 1
			}
			for pos, stmt := range fn.Body.List {
				//log.Println("line ", pos, index, vi.IdentNames)
				if pos >= index {
					//dst.Print(stmt)
					switch stmt.(type) {
					case *dst.BlockStmt:
						if vi.Symbol != "{}" {
							break
						}
						//log.Println("insert BlockStmt", vi.IdentNames)
						stmt.(*dst.BlockStmt).List = insertStatements(stmt.(*dst.BlockStmt).List, pos, insertStmt.(dst.Stmt))
					case *dst.ReturnStmt:
						if vi.Symbol != token.RETURN.String() {
							break
						}
						//log.Println("insert ReturnStmt", vi.IdentNames)
						results := stmt.(*dst.ReturnStmt).Results
						if len(results) != len(vi.IdentNames) {
							break
						}
						for i, ret := range results {
							_, ok := ret.(*dst.UnaryExpr)
							if !ok {
								continue
							}
							name1 := ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Type.(*dst.Ident).Name
							//name2 := insertStmt.(*dst.ReturnStmt).Results[0].(*dst.UnaryExpr).X.(*dst.CompositeLit).Type.(*dst.Ident).Name
							if name1 == vi.IdentNames[i] {
								kv := ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts
								addKv := insertStmt.(*dst.ReturnStmt).Results[0].(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts
								//未导入的才添加
								for _, add := range addKv {
									has := false
									for _, exist := range kv {
										if add.(*dst.KeyValueExpr).Key.(*dst.Ident).Name == exist.(*dst.KeyValueExpr).Key.(*dst.Ident).Name {
											has = true
											break
										}
									}
									if !has {
										kv = append(kv, add)
									}
								}
								ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts = kv
							}
						}
					case *dst.AssignStmt:
						if vi.Symbol != token.DEFINE.String() {
							break
						}
						//log.Println("insert AssignStmt", vi.IdentNames)
						fn.Body.List = insertStatements(fn.Body.List, pos, insertStmt.(dst.Stmt))
					default:
						//log.Println("insert default", vi.IdentNames)
					}
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

func (vi *FuncMeta) RollBack(node dst.Node) dst.Visitor {
	if funDecl, ok := node.(*dst.FuncDecl); ok {
		if funDecl.Name.Name == vi.FuncName {
			block := funDecl.Body
			vi.rollBack(block)
		}
	}
	return vi
}

func (vi *FuncMeta) GetNode() dst.Node {
	if vi.node != nil {
		return vi.node
	}
	var insertStmt dst.Stmt
	switch vi.Symbol {
	case ":=":
		if len(vi.Variables) == 0 || len(vi.IdentNames) == 0 {
			panic("赋值表达式左右两边不能为null")
		}
		var rhs dst.Expr
		if len(vi.Parameters) == 0 {
			//是变量 *dst.SelectorExpr
			rhs = vi.GetSelector()
		} else {
			//是调用
			rhs = &dst.CallExpr{
				Fun:  vi.GetSelector(),
				Args: vi.GetParameters(),
			}
		}
		//变量声明
		insertStmt = &dst.AssignStmt{
			// 等式左边
			Lhs: vi.GetVariables(),
			// := 符号
			Tok: token.DEFINE,
			// 等式右边
			Rhs: []dst.Expr{
				rhs,
			},
			Decs: dst.AssignStmtDecorations{
				NodeDecs: dst.NodeDecs{
					Start: dst.Decorations{},
				},
			},
		}
	case "{}":
		log.Println("bbbbb")
		insertStmt = &dst.BlockStmt{
			List:           []dst.Stmt{},
			RbraceHasNoPos: false,
			Decs: dst.BlockStmtDecorations{
				NodeDecs: dst.NodeDecs{
					Before: dst.NewLine,
					Start:  dst.Decorations{"//" + vi.Comments},
				},
				Lbrace: nil,
			},
		}
	case "":
		log.Println("aaaaa", len(vi.IdentNames))
		// 调用语句,没有左边
		insertStmt = &dst.ExprStmt{
			X: &dst.CallExpr{
				Fun:  vi.GetSelector(),
				Args: vi.GetParameters(),
			},
		}
	case token.RETURN.String():

	default:
		log.Println("default", len(vi.IdentNames))
		return nil
	}
	return insertStmt
}

func (vi *FuncMeta) GetCode() string {
	injectCode := fmt.Sprintf("import")

	return injectCode
}

func NewFuncMete(findFuncName, importCode string) *FuncMeta {
	meta := &FuncMeta{}
	input := fmt.Sprintf(`
package main
func %v(){
   %v 
}`, findFuncName, importCode)
	fParse, err := decorator.ParseFile(token.NewFileSet(), "", input, 0)
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}
	dst.Inspect(fParse, func(n dst.Node) bool {
		switch x := n.(type) {
		case *dst.FuncDecl:
			if x.Name.Name != findFuncName {
				break
			}
			//dst.Print(x)
			meta.FuncName = findFuncName
			for _, field := range x.Body.List {
				meta.node = field
				switch s := field.(type) {
				case *dst.AssignStmt:

					variables := []string{}
					for _, lhs := range s.Lhs {
						variables = append(variables, lhs.(*dst.Ident).Name)
					}

					indentNames := []string{}
					parameters := []interface{}{}
					switch callExpr := s.Rhs[0].(type) {
					case *dst.CallExpr:
						idents := extractIdents(callExpr.Fun)
						for _, ident := range idents {
							indentNames = append(indentNames, ident.Name)
						}
						for _, arg := range callExpr.Args {
							switch arg.(type) {
							case *dst.BasicLit:
								value, _ := inferType(arg.(*dst.BasicLit).Value)
								parameters = append(parameters, value)
							case *dst.Ident:
								parameters = append(parameters, arg.(*dst.Ident).Name)
							}
						}
					case *dst.SelectorExpr:
						idents := extractIdents(callExpr)
						for _, ident := range idents {
							indentNames = append(indentNames, ident.Name)
						}
					}

					meta.Variables = variables
					meta.Symbol = s.Tok.String()
					meta.IdentNames = indentNames
					meta.Parameters = parameters

				case *dst.ExprStmt:

					indentNames := []string{}
					callExpr := s.X.(*dst.CallExpr)
					idents := extractIdents(callExpr.Fun)
					for _, ident := range idents {
						indentNames = append(indentNames, ident.Name)
					}

					parameters := []interface{}{}
					for _, arg := range callExpr.Args {
						switch arg.(type) {
						case *dst.BasicLit:
							value, _ := inferType(arg.(*dst.BasicLit).Value)
							parameters = append(parameters, value)
						case *dst.Ident:
							parameters = append(parameters, arg.(*dst.Ident).Name)
						}
					}

					meta.IdentNames = indentNames
					meta.Parameters = parameters

				case *dst.ReturnStmt:
					meta.Symbol = token.RETURN.String()
					//op := s.Results[0].(*dst.UnaryExpr).Op.String()
					ret := s.Results[0].(*dst.UnaryExpr).X.(*dst.CompositeLit)
					ident := ret.Type.(*dst.Ident).Name

					meta.IdentNames = []string{ident}
				}
			}
		}

		return true
	})

	//log.Println("GetCode", meta.GetCode())
	//log.Println("NewFuncMete", jsonconv.ObjectToJsonIndent(meta))
	return meta
}

func extractIdents(node dst.Node) []*dst.Ident {
	var idents []*dst.Ident

	switch n := node.(type) {
	case *dst.SelectorExpr:
		idents = append(idents, extractIdents(n.X)...)
		idents = append(idents, extractIdents(n.Sel)...)
	case *dst.Ident:
		idents = append(idents, n)
	}

	return idents
}

func insertStatements(stmts []dst.Stmt, pos int, toInsert ...dst.Stmt) []dst.Stmt {
	return append(stmts[:pos], append(toInsert, stmts[pos:]...)...)
}

func (vi *FuncMeta) GetVariables() []dst.Expr {
	var variables []dst.Expr
	for _, item := range vi.Variables {
		variables = append(variables, dst.NewIdent(item))
	}
	return variables
}

func (vi *FuncMeta) GetSelector() dst.Expr {
	var selector interface{}
	if len(vi.IdentNames) == 1 {
		//var selector *dst.Ident
		selector = dst.NewIdent(vi.IdentNames[0])
	} else {
		// >=2
		selector = &dst.SelectorExpr{
			// 只有一个 .
			X: &dst.Ident{
				Name: cast.ToString(vi.IdentNames[0]),
			},
			// IdentName.SelName
			Sel: &dst.Ident{
				Name: cast.ToString(vi.IdentNames[1]),
			},
		}
		for _, value := range vi.IdentNames[2:] {
			selector = &dst.SelectorExpr{
				X: selector.(*dst.SelectorExpr),
				//IdentName.SelName.SelName
				Sel: dst.NewIdent(cast.ToString(value)),
			}
		}
	}
	return selector.(dst.Expr)
}

func (vi *FuncMeta) GetParameters() []dst.Expr {
	var varExpr []dst.Expr
	for _, value := range vi.Parameters {
		exp := &dst.BasicLit{
			Kind:  kindToToken(reflect.TypeOf(value).Kind()),
			Value: jsonconv.ObjectToJson(value),
		}
		varExpr = append(varExpr, exp)
	}
	return varExpr
}

func (vi *FuncMeta) rollBack(block *dst.BlockStmt) dst.Visitor {
	for k, v := range block.List {
		switch stmt := v.(type) {
		//赋值表达式
		case *dst.AssignStmt:
			if vi.Symbol != ":=" {
				break
			}
			if len(vi.Variables) == len(stmt.Lhs) {
				for i, lhs := range stmt.Lhs {
					switch nn := lhs.(type) {
					case *dst.Ident:
						if nn.Name == vi.Variables[i] {
							//删除
							block.List = append(append([]dst.Stmt{}, block.List[:k]...), block.List[k+1:]...)
							return vi
						}
					}
				}
			}
		case *dst.ExprStmt: //表达语句
			if vi.Symbol != "" {
				break
			}
			switch expr := stmt.X.(type) {
			case *dst.CallExpr: //是调用语句
				if len(vi.IdentNames) == 0 {
					break
				}
				if call, ok := expr.Fun.(*dst.SelectorExpr); ok {
					x, _ := call.X.(*dst.Ident)
					// 包名和方法名都相等
					if x.Name == vi.IdentNames[0] {
						//删除
						block.List = append(append([]dst.Stmt{}, block.List[:k]...), block.List[k+1:]...)
						return vi
					}
				}
			}
		case *dst.BlockStmt: //{}
			if vi.Symbol != "{}" {
				break
			}
			vi.rollBack(stmt)
			if len(stmt.List) == 0 {
				//删除空块
				block.List = append(append([]dst.Stmt{}, block.List[:k]...), block.List[k+1:]...)
				return vi
			}
		case *dst.ReturnStmt:
			if vi.Symbol != "return" {
				break
			}
			results := stmt.Results
			insertStmt := vi.GetNode()
			if len(results) != len(vi.IdentNames) {
				break
			}

			for i, ret := range results {
				_, ok := ret.(*dst.UnaryExpr)
				if !ok {
					continue
				}
				name1 := ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Type.(*dst.Ident).Name
				//找到需要加入的变量名称
				if name1 == vi.IdentNames[i] {
					kv := ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts
					addKv := insertStmt.(*dst.ReturnStmt).Results[0].(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts
					//未导入的才添加
					for _, add := range addKv {
						for j, exist := range kv {
							//如果已经包含
							if add.(*dst.KeyValueExpr).Key.(*dst.Ident).Name == exist.(*dst.KeyValueExpr).Key.(*dst.Ident).Name {
								kv = append(kv[:j], kv[j+1:]...)
								//break
							}
						}

					}
					ret.(*dst.UnaryExpr).X.(*dst.CompositeLit).Elts = kv
				}
			}
		}
	}

	return vi
}
