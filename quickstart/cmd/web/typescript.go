/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/convertx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/cobrax"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/parserx"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/parserx/aspec"
)

type typescriptFlags struct {
	VarStringTplPath     string `name:"var_string_tpl_path" shorthand:"t"`     // 模板路径
	VarStringOutPath     string `name:"var_string_out_path" shorthand:"o"`     // 文件输出路径
	VarStringNameAs      string `name:"var_string_name_as" shorthand:"n"`      // 文件命名模版 %s.go
	VarStringMode        string `name:"var_string_mode" shorthand:"m"`         // 解析模式 swagger、api、ast
	VarStringFilePath    string `name:"var_string_file_path" shorthand:"f"`    // 文件路径
	VarStringIgnoreModel string `name:"var_string_ignore_model" shorthand:"i"` // 忽略的模型
}

var flag = &typescriptFlags{}

// typescriptCmd represents the typescript command
var typescriptCmd = &cobra.Command{
	Use:   "typescript",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("typescript called")
		RunTypescriptCmd(cmd, args)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// typescriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// typescriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// typescriptCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	// typescriptCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	// typescriptCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	// typescriptCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
	cobrax.ParseFlag(typescriptCmd, flag)
}

func RunTypescriptCmd(cmd *cobra.Command, args []string) {
	RunTypescript(flag)
}

func RunTypescript(conf *typescriptFlags) {
	log.Println("typescript called", jsonconv.ObjectToJsonIndent(conf))
	var err error
	var sp *aspec.ApiSpec
	switch conf.VarStringMode {
	case "api":
		sp, err = parserx.NewSpecParser().ParseApi(conf.VarStringFilePath)
	case "swagger":
		sp, err = parserx.NewSwaggerParser().ParseApi(conf.VarStringFilePath)
	case "ast":
	}

	if err != nil {
		panic(err)
	}

	err = generateApiTs(sp, conf)
	if err != nil {
		panic(err)
	}

	err = generateTypesTs(sp, conf)
	if err != nil {
		panic(err)
	}
}

func generateApiTs(sp *aspec.ApiSpec, conf *typescriptFlags) error {
	t := path.Join(conf.VarStringTplPath, "api.ts.tpl")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	gps := make(map[string][]aspec.Route)
	for _, g := range sp.Service.Groups {
		group := g.Annotation.Properties["group"]
		gps[group] = append(gps[group], g.Routes...)
	}

	ims := strings.Split(conf.VarStringIgnoreModel, ",")
	for n, g := range gps {

		mt := make(map[string]aspec.Type)
		for _, r := range g {
			if r.RequestType != nil {
				name := r.RequestType.Name()
				mt[name] = r.RequestType
			}
			if r.ResponseType != nil {
				name := r.ResponseType.Name()
				mt[name] = r.ResponseType
			}
		}

		var tns []string
		for k := range mt {
			if !slices.Contains(ims, k) {
				tns = append(tns, k)
			}
		}

		var trs []TsApiRoute
		for _, r := range g {
			req := "any"
			resp := "any"
			if r.RequestType != nil {
				req = r.RequestType.Name()
			}

			if r.ResponseType != nil {
				resp = r.ResponseType.Name()
			}

			tr := TsApiRoute{
				Summery:  r.AtDoc.Text,
				Path:     r.Path,
				Method:   r.Method,
				Handler:  jsonconv.Lcfirst(r.Handler) + "Api",
				Request:  req,
				Response: resp,
			}
			trs = append(trs, tr)
		}

		var fn = n
		if fn == "" {
			fn = strings.Split(path.Base(conf.VarStringFilePath), ".")[0]
		}

		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(conf.VarStringOutPath, fmt.Sprintf("%s.ts", fn)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"Join": func(s []string) string {
					return strings.Join(s, ", ")
				},
			},
			Data: TsApi{
				ImportPkgPaths: []string{`import request from "@/utils/request"`},
				ImportTypes:    tns,
				Routes:         trs,
			},
		}
		err := meta.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func generateTypesTs(sp *aspec.ApiSpec, conf *typescriptFlags) error {
	t := path.Join(conf.VarStringTplPath, "types.ts.tpl")
	o := path.Join(conf.VarStringOutPath, "types.ts")

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	ims := strings.Split(conf.VarStringIgnoreModel, ",")
	ts := make(map[string]TsType)

	for _, v := range sp.Types {
		if !slices.Contains(ims, v.Name()) {
			ts[v.Name()] = convertTypeTs(v)
		}
	}

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

type TsApi struct {
	ImportPkgPaths []string
	ImportTypes    []string
	Routes         []TsApiRoute
}
type TsApiRoute struct {
	Summery  string
	Path     string
	Method   string
	Handler  string
	Request  string
	Response string
}

type TsType struct {
	Comment string
	Name    string
	Extends []string
	Fields  []TsTypeField
}

type TsTypeField struct {
	Name    string
	Type    string
	Comment string
}

func convertTypeTs(st aspec.Type) TsType {
	var ts TsType

	switch t := st.(type) {
	case aspec.DefineStruct:

		var ex []string
		var tfs []TsTypeField
		for _, v := range t.Members {
			if v.Name == "" {
				ex = append(ex, v.Type.Name())
			} else {
				m := TsTypeField{
					Comment: v.Comment,
					Name:    jsonconv.Case2Snake(v.Name),
					Type:    convertx.ConvertGoTypeToTsType(v.Type.Name()),
				}
				tfs = append(tfs, m)
			}
		}

		ts = TsType{
			Comment: strings.Join(t.Comments(), "\n"),
			Name:    t.Name(),
			Extends: ex,
			Fields:  tfs,
		}

	case aspec.PrimitiveType:
	case aspec.MapType:
	case aspec.ArrayType:
	case aspec.InterfaceType:
	case aspec.PointerType:
	default:
		panic("unknown type")
	}

	return ts
}
