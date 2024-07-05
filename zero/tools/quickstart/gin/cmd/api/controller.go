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

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandController(cmd, args)
	},
}

func init() {
	controllerCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	controllerCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	controllerCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	controllerCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
}

func RunCommandController(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := parsex.ParseAPI(f)
	if err != nil {
		panic(err)
	}

	err = generateControllers(sp, t, o, n)
	if err != nil {
		panic(err)
	}
}

func generateControllers(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(tplPath, "controller.tpl"))
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
			CodeOutPath:    path.Join(outPath, "controller", fmt.Sprintf(nameAs, v.Name+".ctl")),
			TemplateString: string(handlerTpl),
			Data: map[string]any{
				"Package": "controller",
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
