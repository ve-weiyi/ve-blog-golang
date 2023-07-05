package inject

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func TestInject(t *testing.T) {
	inject := AstInjectMeta{
		FilePath: "./test/test.go",
		//ImportMetas: []*ImportMeta{
		//	NewImportMete(`jsoniter "github.com/json-iterator/go"`),
		//	NewImportMete(`"go/ast"`),
		//},
		//StructMetas: []*StructMeta{
		//	NewStructMete("Context", `Alias  *dst.Visitor //元素别名`),
		//},
		FuncMetas: []*FuncMeta{
			//&FuncMeta{
			//	FuncName:   "GetName",
			//	FuncPos:    4,
			//	Variables:  []string{"value", "err"},
			//	Symbol:     ":=",
			//	IdentNames: []string{"jsoniter", "AA", "ConfigCompatibleWithStandardLibrary"},
			//	//Parameters: []interface{}{"ss", 11},
			//},
			NewFuncMete("NewApiContext", `json := jsoniter.ConfigCompatibleWithStandardLibrary()`),
			NewFuncMete("NewApiContext", `return &Context{
			visit: ast.NewIdent("hello"),
			}`),
		},
	}
	var err error
	inject.Walk()
	err = inject.RollBack()
	//err = inject.Inject()
	log.Println("-->", err)
	if err != nil {
		return
	}
}
func TestNewAst(t *testing.T) {
	//NewImportMete(`jsoniter "github.com/json-iterator/go"`)
	//NewStructMete("ApiGroup", `Alias   *dst.Visitor //元素别名`)
	//NewFuncMete("NewApiContext", `json := jsoniter.ConfigCompatibleWithStandardLibrary()`)
	NewFuncMete("NewApiContext", `return &Context{
		visitor: ast.NewIdent("hello"),
	}`)
}
