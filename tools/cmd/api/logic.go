/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/golang"

	"github.com/ve-weiyi/ve-blog-golang/kit/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/api/helper"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/apiparser"
)

var logicCmd = &cobra.Command{
	Use:   "logic",
	Short: "生成 Logic 代码",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generateLogics(cmdVar)
	},
}

func init() {
	ParseFlagVar(logicCmd)
}

func generateLogics(cv *CmdVar) error {
	sp, err := apiparser.NewSpecParser().ParseApi(cv.VarStringApiFile)
	if err != nil {
		return err
	}

	var metas []invent.TemplateMeta

	tpl, err := os.ReadFile(path.Join(cv.VarStringTplPath, "logic.tpl"))
	if err != nil {
		return err
	}

	pkg, _, _ := golang.GetParentPackage(cv.VarStringOutPath)

	var groups map[string][]helper.GroupRoute
	groups = helper.ConvertRouteGroups(sp)

	for k, v := range groups {

		m := invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(cv.VarStringOutPath, "logic", fmt.Sprintf(cv.VarStringNameAs, k)),
			TemplateString: string(tpl),
			Data: map[string]any{
				"Package": "logic",
				"Imports": []string{
					fmt.Sprintf(`"%s"`, cv.VarContextPackage),
					fmt.Sprintf(`"%s/types"`, pkg),
				},
				"Group":       jsonconv.Case2Camel(k),
				"GroupRoutes": v,
			},
			FunMap: map[string]any{
				"pkgTypes": func(input string) string {
					// 使用正则表达式匹配单词
					re := regexp.MustCompile(`\w+`)

					// 替换每个单词，前面添加 'types.'
					result := re.ReplaceAllString(input, "types.$0")

					if strings.HasPrefix(result, "[]") {
						return result
					}

					return fmt.Sprintf("*%v", result)
				},
			},
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
