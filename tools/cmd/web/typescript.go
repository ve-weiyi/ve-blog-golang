package web

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/kit/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/web/helper"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/apiparser"
	"github.com/ve-weiyi/ve-blog-golang/tools/parserx/apiparser/aspec"
)

var typescriptCmd = &cobra.Command{
	Use:   "typescript",
	Short: "生成 TypeScript API 代码",
	RunE:  runTypescript,
}

func runTypescript(cmd *cobra.Command, args []string) error {
	log.Printf("generating TypeScript code from %s\n", cmdVar.VarStringApiFile)
	var sp *aspec.ApiSpec
	var err error

	switch cmdVar.VarStringMode {
	case "api":
		sp, err = apiparser.NewSpecParser().ParseApi(cmdVar.VarStringApiFile)
	case "swagger":
		sp, err = apiparser.NewSwaggerParser().ParseApi(cmdVar.VarStringApiFile)
	default:
		return fmt.Errorf("unsupported mode: %s", cmdVar.VarStringMode)
	}

	if err != nil {
		return err
	}

	if err = generateApiTs(sp, cmdVar); err != nil {
		return err
	}

	if err = generateTypesTs(sp, cmdVar); err != nil {
		return err
	}

	log.Println("TypeScript code generated successfully")
	return nil
}

func generateApiTs(sp *aspec.ApiSpec, conf *CmdVar) error {
	t := path.Join(conf.VarStringTplPath, "api.ts.tpl")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	groups := helper.ConvertApiTs(sp)

	var metas []invent.TemplateMeta
	for k, v := range groups {

		mt := make(map[string]string)
		for _, g := range v {
			for _, r := range g.Routes {
				if r.Request != "" {
					name := strings.Trim(r.Request, "[]")
					mt[name] = name
				}
				if r.Response != "" {
					name := strings.Trim(r.Response, "[]")
					mt[name] = name
				}
			}
		}

		var ims []string
		for kk, _ := range mt {
			if kk == "any" {
				continue
			}
			ims = append(ims, kk)
		}

		sort.Slice(ims, func(i, j int) bool {
			return ims[i] < ims[j]
		})

		meta := invent.TemplateMeta{
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(conf.VarStringOutPath, fmt.Sprintf("%s.ts", k)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"Join": func(s []string) string {
					return strings.Join(s, ", ")
				},
			},
			Data: map[string]any{
				"ImportPkgPaths": []string{
					`import request from "@/utils/request";`,
					fmt.Sprintf(`import type { %s } from "./types";`, strings.Join(ims, ", ")),
				},
				"Name":        jsonconv.Case2Camel(k),
				"TsApiGroups": v,
			},
		}

		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func generateTypesTs(sp *aspec.ApiSpec, conf *CmdVar) error {
	t := path.Join(conf.VarStringTplPath, "types.ts.tpl")
	o := path.Join(conf.VarStringOutPath, "types.ts")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	//ims := strings.Split(conf.VarStringIgnoreModel, ",")
	ts := helper.ConvertTypeTs(sp)

	meta := invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    o,
		TemplateString: string(tpl),
		FunMap: map[string]any{
			"Join": func(s []string) string {
				return strings.Join(s, ", ")
			},
		},
		Data: ts,
	}

	return meta.Execute()
}

func generatePermsTs(sp *aspec.ApiSpec, conf *CmdVar) error {
	t := path.Join(conf.VarStringTplPath, "perms.ts.tpl")
	o := path.Join(conf.VarStringOutPath, "perms.ts")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	groups := helper.ConvertApiTs(sp)

	meta := invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    o,
		TemplateString: string(tpl),
		FunMap:         nil,
		Data: map[string]any{
			"Name":   jsonconv.Case2Camel(sp.Service.Name),
			"Groups": groups,
		},
	}

	return meta.Execute()
}
