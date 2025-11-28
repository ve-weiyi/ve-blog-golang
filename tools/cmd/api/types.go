/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/gogen"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"

	"github.com/ve-weiyi/ve-blog-golang/kit/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/api/helper"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/apiparser"
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "生成 Types 类型代码",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generateTypes(cmdVar)
	},
}

func init() {
	ParseFlagVar(typesCmd)
}

func generateTypes(cv *CmdVar) error {
	sp, err := apiparser.ParseApiSpec(cv.VarStringApiFile)
	if err != nil {
		return err
	}

	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(cv.VarStringTplPath, "types.tpl"))
	if err != nil {
		return err
	}

	groupTypes := helper.GroupTypes(sp, "types")
	for k, typeGroup := range groupTypes {

		types := make([]spec.Type, 0, len(typeGroup))
		for _, v := range typeGroup {
			types = append(types, v)
		}

		val, err := gogen.BuildTypes(types)
		if err != nil {
			return err
		}

		meta := invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(cv.VarStringOutPath, "types", fmt.Sprintf(cv.VarStringNameAs, k)),
			TemplateString: string(handlerTpl),
			Data: map[string]any{
				"Package": "types",
				"Imports": []string{},
				"Name":    jsonconv.Case2Camel(k),
				"Types":   val,
			},
		}

		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
