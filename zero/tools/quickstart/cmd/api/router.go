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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunCommandRouter(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := ParseAPI(f)
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

	var groups []GroupRoute
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
