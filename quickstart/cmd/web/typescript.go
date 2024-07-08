/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package web

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/convertx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/quickstart/tools/cobrax"
)

type typescriptFlags struct {
	VarStringModel   string `name:"var_string_model" shorthand:"m"`    // 解析模式 swagger、api、ast
	VarStringTplPath string `name:"var_string_tpl_path" shorthand:"t"` // 模板路径
	VarStringOutPath string `name:"var_string_out_path" shorthand:"o"` // 文件输出路径
	VarStringNameAs  string `name:"var_string_name_as" shorthand:"n"`  // 文件命名模版 %s.go
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
		fmt.Println("typescript called")
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
	//typescriptCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	//typescriptCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	//typescriptCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	//typescriptCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
	cobrax.ParseFlag(typescriptCmd, flag)
}

func RunTypescriptCmd(cmd *cobra.Command, args []string) {
	fmt.Println("typescript called", flag.VarStringModel, flag.VarStringTplPath, flag.VarStringOutPath, flag.VarStringNameAs)

	//var sp *spec.ApiSpec
	switch flag.VarStringModel {
	case "api":
	case "swagger":
	case "ast":
	}

	//
	//err = generateTypescript(sp, t, o, n)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = generateTypesTs(sp, t, o, n)
	//if err != nil {
	//	panic(err)
	//}
}

func generateTypescript(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	apiTemplate, err := os.ReadFile(path.Join(tplPath, "api.ts.tpl"))
	if err != nil {
		return err
	}

	for _, g := range sp.Service.Groups {
		log.Printf("%v", jsonconv.ObjectToJsonIndent(g))

		mmp := make(map[string]spec.Type)
		for _, r := range g.Routes {
			if r.RequestType != nil {
				name := convertx.ConvertGoTypeToTsType(r.RequestType.Name())
				name = strings.ReplaceAll(name, "[]", "")
				mmp[name] = r.RequestType
			}
			if r.ResponseType != nil {
				name := convertx.ConvertGoTypeToTsType(r.ResponseType.Name())
				name = strings.ReplaceAll(name, "[]", "")
				mmp[name] = r.ResponseType
			}
		}
		var models []string
		for k := range mmp {

			models = append(models, k)
		}

		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    fmt.Sprintf("./api/%s.ts", g.Annotation.Properties["group"]),
			TemplateString: string(apiTemplate),
			FunMap: map[string]any{
				"joinArray": utils.JoinArray,
				"convertJson": func(name string) string {
					if name == "ID" {
						return "id"
					}
					return jsonconv.Case2Snake(name)
				},
				"convertTsType": convertx.ConvertGoTypeToTsType,
				"convertHandler": func(name string) string {
					return jsonconv.Case2CamelLowerStart(name)
				},
			},
			Data: map[string]any{
				"Imports": []string{`import http from "@/utils/request"`},
				"Models":  models,
				"Base":    g.Annotation.Properties["prefix"],
				"Routes":  g.Routes,
			},
		}
		err := meta.Execute()
		fmt.Println(err)
	}

	return nil
}

func generateTypesTs(sp *spec.ApiSpec, tplPath, outPath, nameAs string) error {
	tpl, err := os.ReadFile(path.Join(tplPath, "types.ts.tpl"))
	if err != nil {
		return err
	}

	ts := make(map[string]*TypesDeclare)

	for _, v := range sp.Types {
		t := &TypesDeclare{
			Comment: strings.Join(v.Documents(), "\n"),
			Name:    v.Name(),
			Extends: getExtends(v),
			Fields:  getFields(v),
		}

		ts[t.Name] = t
	}

	meta := invent.TemplateMeta{
		Mode:           invent.ModeCreateOrReplace,
		CodeOutPath:    path.Join(outPath, "types.ts"),
		TemplateString: string(tpl),
		FunMap:         nil,
		Data:           ts,
	}

	return meta.Execute()
}

type ApiDeclare struct {
}

type TypesDeclare struct {
	Comment string
	Name    string
	Extends []string
	Fields  []Field
}

type Field struct {
	Comment string
	Name    string
	Type    string
}

func getExtends(st spec.Type) []string {
	var ex []string

	switch t := st.(type) {
	case spec.DefineStruct:
		for _, v := range t.Members {
			if v.Name == "" {
				switch tt := v.Type.(type) {
				case spec.DefineStruct:
					ex = append(ex, tt.RawName)
				}
			}
		}

	}
	return ex
}

func getFields(st spec.Type) []Field {
	var ex []Field

	switch t := st.(type) {
	case spec.DefineStruct:
		for _, v := range t.Members {
			m := Field{
				Comment: v.Comment,
				Name:    v.Name,
				Type:    convertx.ConvertGoTypeToTsType(getType(v.Type)),
			}
			ex = append(ex, m)
		}

	}
	return ex
}

func getType(st spec.Type) string {
	var out string
	switch t := st.(type) {
	case spec.DefineStruct:
		out = t.RawName
	case spec.PrimitiveType:
		out = t.RawName
	case spec.MapType:
		out = t.RawName
	case spec.ArrayType:
		out = t.RawName
	case spec.InterfaceType:
		out = t.RawName
	case spec.PointerType:
		out = t.RawName
	default:
		panic("unknown type")
	}
	return out
}
