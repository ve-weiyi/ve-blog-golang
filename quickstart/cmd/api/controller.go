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
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/cmd/api/helper"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/parserx"
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
		err := generateControllers(cmdVar)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ParseFlagVar(controllerCmd)
}

func generateControllers(cv *CmdVar) error {
	sp, err := parserx.NewSpecParser().ParseApi(cv.VarStringApiFile)
	if err != nil {
		return err
	}

	var metas []invent.TemplateMeta

	tpl, err := os.ReadFile(path.Join(cv.VarStringTplPath, "controller.tpl"))
	if err != nil {
		return err
	}

	pkg, _ := golang.GetParentPackage(cv.VarStringOutPath)

	var groups map[string][]helper.GroupRoute
	groups = helper.ConvertRouteGroups(sp)

	for k, v := range groups {

		m := invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(cv.VarStringOutPath, "controller", fmt.Sprintf(cv.VarStringNameAs, k)),
			TemplateString: string(tpl),
			Data: map[string]any{
				"Package": "controller",
				"Imports": []string{
					fmt.Sprintf(`"%s"`, cv.VarContextPackage),
					fmt.Sprintf(`"%s/dto"`, pkg),
					fmt.Sprintf(`"%s/service"`, pkg),
				},
				"Group":       jsonconv.Case2Camel(k),
				"GroupRoutes": v,
			},
			FunMap: invent.StdMapUtils,
		}

		metas = append(metas, m)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
