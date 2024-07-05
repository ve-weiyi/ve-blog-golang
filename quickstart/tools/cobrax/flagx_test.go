package cobrax

import (
	"testing"

	"github.com/spf13/cobra"
)

type typescriptFlags struct {
	VarStringModel   string `name:"var_string_model" shorthand:"m"`    // 解析模式 swagger、api、ast
	VarStringTplPath string `name:"var_string_tpl_path" shorthand:"t"` // 模板路径
	VarStringOutPath string `name:"var_string_out_path" shorthand:"o"` // 文件输出路径
	VarStringNameAs  string `name:"var_string_name_as" shorthand:"n"`  // 文件命名模版 %s.go
}

func Test_Flags(t *testing.T) {
	res := &typescriptFlags{}
	// 使用反射解析结构体中的字段
	ParseFlag(&cobra.Command{}, res)
}
