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
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/tools/parsex"
)

// typescriptCmd represents the typescript command
var typescriptCmd = &cobra.Command{
	Use:   "typescript",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandTypescript(cmd, args)
	},
}

func init() {
	typescriptCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	typescriptCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	typescriptCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	typescriptCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
}

func RunCommandTypescript(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := parsex.ParseAPI(f)
	if err != nil {
		panic(err)
	}

	err = generateTypescripts(sp, t, o, n)
	if err != nil {
		panic(err)
	}
}

func generateTypescripts(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(tplPath, "typescript.tpl"))
	if err != nil {
		return err
	}

	var groups []parsex.GroupRoute
	groups = convertGroups(sp)

	pkg, _ := golang.GetParentPackage(outPath)
	// handler
	for _, v := range groups {

		metas = append(metas, invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(outPath, "typescript", fmt.Sprintf(nameAs, v.Name+".ctl")),
			TemplateString: string(handlerTpl),
			Data: map[string]any{
				"Package": "typescript",
				"Imports": []string{
					fmt.Sprintf(`"%s/types"`, pkg),
					fmt.Sprintf(`"%s/service"`, pkg),
				},
				"Name":   jsonconv.Case2Camel(v.Name),
				"Routes": v.Routes,
			},
			FunMap: invent.StdMapUtils,
		})
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
