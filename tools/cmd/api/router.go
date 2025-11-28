/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"

	"github.com/ve-weiyi/ve-blog-golang/kit/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/api/helper"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/apiparser"
)

var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "生成 Router 路由代码",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := generateRouters(cmdVar); err != nil {
			return err
		}
		return generateRoutes(cmdVar)
	},
}

func init() {
	ParseFlagVar(routerCmd)
}

func generateRouters(cv *CmdVar) error {
	sp, err := apiparser.NewSpecParser().ParseApi(cv.VarStringApiFile)
	if err != nil {
		return err
	}

	var metas []invent.TemplateMeta

	tpl, err := os.ReadFile(path.Join(cv.VarStringTplPath, "router.tpl"))
	if err != nil {
		return err
	}

	pkg, _, _ := golang.GetParentPackage(cv.VarStringOutPath)

	var groups map[string][]helper.GroupRoute
	groups = helper.ConvertRouteGroups(sp)

	for k, v := range groups {

		m := invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(cv.VarStringOutPath, "router", fmt.Sprintf(cv.VarStringNameAs, k)),
			TemplateString: string(tpl),
			Data: map[string]any{
				"Package": "router",
				"Imports": []string{
					fmt.Sprintf(`"%s"`, cv.VarContextPackage),
					fmt.Sprintf(`"%s/types"`, pkg),
					fmt.Sprintf(`"%s/handler"`, pkg),
				},
				"Group":       jsonconv.Case2Camel(k),
				"GroupRoutes": v,
			},
		}

		metas = append(metas, m)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRoutes(cv *CmdVar) error {
	sp, err := apiparser.NewSpecParser().ParseApi(cv.VarStringApiFile)
	if err != nil {
		return err
	}

	var metas []invent.TemplateMeta

	tpl, err := os.ReadFile(path.Join(cv.VarStringTplPath, "routes.tpl"))
	if err != nil {
		return err
	}

	pkg, _, _ := golang.GetParentPackage(cv.VarStringOutPath)

	var groups map[string][]helper.GroupRoute
	groups = helper.ConvertRouteGroups(sp)

	var gps []string
	for k, _ := range groups {
		gps = append(gps, jsonconv.Case2Camel(k))
	}

	sort.Slice(gps, func(i, j int) bool {
		return gps[i] < gps[j]
	})

	m := invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(cv.VarStringOutPath, "routes.go"),
		TemplateString: string(tpl),
		Data: map[string]any{
			"Package": filepath.Base(cv.VarStringOutPath),
			"Imports": []string{
				fmt.Sprintf(`"%s"`, cv.VarContextPackage),
				fmt.Sprintf(`"%s/router"`, pkg),
			},
			"Groups": gps,
		},
		FunMap: invent.StdMapUtils,
	}

	metas = append(metas, m)

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
