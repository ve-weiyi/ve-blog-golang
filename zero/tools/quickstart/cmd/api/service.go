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
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandServices(cmd, args)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunCommandServices(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := ParseAPI(f)
	if err != nil {
		panic(err)
	}

	err = generateServices(sp, t, o, n)
	if err != nil {
		panic(err)
	}
}

func generateServices(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(tplPath, "service.tpl"))
	if err != nil {
		return err
	}

	var groups []GroupRoute
	groups = convertGroups(sp)
	pkg, _ := golang.GetParentPackage(outPath)
	// handler
	for _, v := range groups {
		metas = append(metas, invent.TemplateMeta{
			Mode:           invent.ModeOnlyCreate,
			CodeOutPath:    path.Join(outPath, "service", fmt.Sprintf(nameAs, v.Name+".sv")),
			TemplateString: string(handlerTpl),
			Data: map[string]any{
				"Package": "service",
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
