/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

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

func NewRouterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "router",
		Short: "从api文件生成router",
		Run: func(cmd *cobra.Command, args []string) {
			RunRouterCommand(cmd, args)
		},
	}

	cmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	cmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	cmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	cmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
	return cmd
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

	//GenerateTypes(sp, t, o, n)
	err = GenerateRouters(sp, t, o, n)
	if err != nil {
		panic(err)
	}
	err = GenerateControllers(sp, t, o, n)
	if err != nil {
		panic(err)
	}
	err = GenerateServices(sp, t, o, n)
	if err != nil {
		panic(err)
	}
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

func GenerateRouters(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	routesTpl, err := os.ReadFile(path.Join(tplPath, "routes.tpl"))
	if err != nil {
		return err
	}

	var groups []GroupRoute
	groups = convertGroups(sp)

	// route
	metas = append(metas, invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(outPath, fmt.Sprintf(nameAs, "routes")),
		TemplateString: string(routesTpl),
		Data:           groups,
		FunMap:         invent.StdMapUtils,
	})

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateControllers(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(tplPath, "controller.tpl"))
	if err != nil {
		return err
	}

	var groups []GroupRoute
	groups = convertHandlers(sp)

	// handler
	for _, v := range groups {
		metas = append(metas, invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(outPath, "controller", fmt.Sprintf(nameAs, v.Name)),
			TemplateString: string(handlerTpl),
			Data:           v,
			FunMap:         invent.StdMapUtils,
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

func GenerateServices(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	var metas []invent.TemplateMeta

	handlerTpl, err := os.ReadFile(path.Join(tplPath, "service.tpl"))
	if err != nil {
		return err
	}

	var groups []GroupRoute
	groups = convertHandlers(sp)

	// handler
	for _, v := range groups {
		metas = append(metas, invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(outPath, "service", fmt.Sprintf(nameAs, v.Name)),
			TemplateString: string(handlerTpl),
			Data:           v,
			FunMap:         invent.StdMapUtils,
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

type Route struct {
	Doc      []string // 用户接口
	Handler  string   // UserHandler
	Method   string   // POST、GET、PUT、DELETE
	Path     string   // /api/v1/user
	Request  string
	Response string
}

type Handler struct {
	Doc      []string // 用户接口
	Handler  string   // UserHandler
	Method   string   // POST、GET、PUT、DELETE
	Path     string   // /api/v1/user
	Request  string
	Response string
}

type GroupRoute struct {
	Name   string // account
	Routes []Route
}

func convertHandlers(sp *spec.ApiSpec) (out []GroupRoute) {
	var groups []GroupRoute
	for _, v := range sp.Service.Groups {
		var routes []Route
		for _, r := range v.Routes {

			var doc []string
			for _, d := range r.Doc {
				doc = append(doc, d)
			}

			for _, d := range r.HandlerDoc {
				doc = append(doc, d)
			}

			if r.AtDoc.Text != "" {
				doc = append(doc, strings.Trim(r.AtDoc.Text, "\\"))
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
			name = "default"
		}

		g := GroupRoute{
			Name:   name,
			Routes: routes,
		}
		groups = append(groups, g)
	}
	return groups
}

func convertGroups(sp *spec.ApiSpec) (out []GroupRoute) {
	var groups []GroupRoute
	for _, v := range sp.Service.Groups {
		var routes []Route
		for _, r := range v.Routes {

			var doc []string
			for _, d := range r.Doc {
				doc = append(doc, d)
			}

			for _, d := range r.HandlerDoc {
				doc = append(doc, d)
			}

			if r.AtDoc.Text != "" {
				doc = append(doc, strings.Trim(r.AtDoc.Text, "\\"))
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
			name = "default"
		}

		g := GroupRoute{
			Name:   name,
			Routes: routes,
		}
		groups = append(groups, g)
	}
	return groups
}
