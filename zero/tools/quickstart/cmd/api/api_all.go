/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

var apiAllCmd = &cobra.Command{
	Use:   "all",
	Short: "从api文件生成router",
	Run: func(cmd *cobra.Command, args []string) {
		RunRouterCommand(cmd, args)
	},
}

func init() {
	apiAllCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	apiAllCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	apiAllCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	apiAllCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
}

func RunRouterCommand(cmd *cobra.Command, args []string) {
	f := VarStringApiFile
	t := VarStringTplPath
	o := VarStringOutPath
	n := VarStringNameAs

	sp, err := ParseAPI(f)
	if err != nil {
		panic(err)
	}

	//err = GenerateTypes(sp, t, o, n)
	//if err != nil {
	//	panic(err)
	//}

	err = generateServices(sp, t, o, n)
	if err != nil {
		panic(err)
	}
	//
	//err = generateControllers(sp, t, o, n)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = generateRouters(sp, t, o, n)
	//if err != nil {
	//	panic(err)
	//}
}

func ParseAPI(filename string) (out *spec.ApiSpec, err error) {
	if path.IsAbs(filename) {
		return parser.Parse(filename)
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f := path.Join(dir, filename)
	return parser.Parse(f)
}

func GenerateTypes(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	typesTpl, err := os.ReadFile(path.Join(tplPath, "types.tpl"))
	if err != nil {
		return err
	}

	// types
	metas = append(metas, invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(outPath, "types", fmt.Sprintf(nameAs, "types")),
		TemplateString: string(typesTpl),
		Data:           sp.Types,
		FunMap:         invent.StdMapUtils,
	})

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func convertGroups(sp *spec.ApiSpec) (out []GroupRoute) {
	var groups []GroupRoute
	for _, v := range sp.Service.Groups {
		var routes []Route
		for _, r := range v.Routes {

			var doc string
			for _, d := range r.Doc {
				doc = doc + d
			}

			for _, d := range r.HandlerDoc {
				doc = doc + d
			}

			if r.AtDoc.Text != "" {
				doc = doc + strings.Trim(strings.Trim(r.AtDoc.Text, "\\"), "\"")
			}

			var req, resp string
			if r.RequestType != nil {
				req = r.RequestType.Name()
			}

			if r.ResponseType != nil {
				resp = r.ResponseType.Name()
			}

			rt := Route{
				Method:   strings.ToUpper(r.Method),
				Path:     r.Path,
				Handler:  jsonconv.Case2Camel(r.Handler),
				Doc:      doc,
				Request:  req,
				Response: resp,
			}

			routes = append(routes, rt)
		}

		var name = v.Annotation.Properties["group"]
		if name == "" {
			name = "base"
		}

		g := GroupRoute{
			Name:   name,
			Routes: routes,
		}
		groups = append(groups, g)
	}
	return groups
}
