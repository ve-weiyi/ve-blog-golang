/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/zero/tools/parsex"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandRouter(cmd, args)
	},
}

func init() {
	routerCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	routerCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	routerCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	routerCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
}

func RunCommandRouter(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := parsex.ParseAPI(f)
	if err != nil {
		panic(err)
	}

	err = generateRouters(sp, t, o, n)
	if err != nil {
		panic(err)
	}
}

func generateRouters(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	routesTpl, err := os.ReadFile(path.Join(tplPath, "routes.tpl"))
	if err != nil {
		return err
	}

	var groups []parsex.GroupRoute
	groups = convertGroups(sp)

	pkg, _ := golang.GetParentPackage(outPath)
	// route
	metas = append(metas, invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(outPath, fmt.Sprintf(nameAs, "routes")),
		TemplateString: string(routesTpl),
		Data: map[string]any{
			"Package": path.Base(outPath),
			"Imports": []string{
				fmt.Sprintf(`"%s/types"`, pkg),
				fmt.Sprintf(`"%s/controller"`, pkg),
			},
			"Groups": groups,
		},
		FunMap: invent.StdMapUtils,
	})

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
