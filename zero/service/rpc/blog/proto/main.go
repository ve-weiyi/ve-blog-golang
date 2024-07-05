package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/field"
	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

var sqlFile = flag.String("f", "test.sql", "the sql file")
var tplFile = flag.String("t", "model.tpl", "the template file")
var outPath = flag.String("o", "./", "the output path")

func main() {
	flag.Parse()

	fmt.Println(*sqlFile)
	fmt.Println(*tplFile)
	fmt.Println(*outPath)

	c := Config{
		SqlFile: *sqlFile,
		TplFile: *tplFile,
		OutPath: *outPath,
	}
	GenerateModel(c)
}

type Config struct {
	SqlFile string
	TplFile string
	OutPath string
}

func GenerateModel(c Config) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	f := path.Join(dir, c.SqlFile)
	t := path.Join(dir, c.TplFile)
	o := path.Join(dir, c.OutPath)

	tables, err := parser.Parse(f, "go_zero", false)
	if err != nil {
		return err
	}

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	for _, table := range tables {
		fmt.Printf("%+v\n", table.Name)

		var fs []*field.Field
		for _, e := range table.Fields {
			//fmt.Printf("%+v", jsonconv.ObjectToJsonIndent(e))
			fs = append(fs, convertField(e))
		}

		var ufs [][]*field.Field
		for _, e := range table.UniqueIndex {
			var u []*field.Field
			for _, f := range e {
				u = append(u, convertField(f))
			}
			ufs = append(ufs, u)
		}

		data := map[string]any{
			"tableName":             table.Name.Source(),
			"upperStartCamelObject": jsonconv.Case2Camel(table.Name.Source()),
			"lowerStartCamelObject": jsonconv.Case2CamelLowerStart(table.Name.Source()),
			"snakeName":             jsonconv.Case2Snake(table.Name.Source()),
			"fields":                fs,
			"uniqueFields":          ufs,
		}

		var metas []invent.TemplateMeta
		metas = append(metas, invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(o, fmt.Sprintf("%v_model_gen.go", table.Name.Source())),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"funcFieldsKey": func(fs []*field.Field) string {
					var name string
					for _, ff := range fs {
						//if name != "" {
						//	name += "And"
						//}
						name += ff.Name
					}

					return name
				},
				"funcFieldsKeyVar": func(fs []*field.Field) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						tp := ff.Type
						if name != "" {
							name += ", "
						}
						name += fmt.Sprintf("%s %s", v, tp)
					}

					return name
				},
				"funcFieldsKeyCond": func(fs []*field.Field) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += " and "
						}
						name += fmt.Sprintf("`%s` = ?", v)
					}

					return name
				},
				"funcFieldsKeyCondVar": func(fs []*field.Field) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += ", "
						}
						name += v
					}

					return name
				},
			},
			Data: data,
		})

		for _, m := range metas {
			err = m.Execute()
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func convertField(e *parser.Field) *field.Field {

	return &field.Field{
		Name:    jsonconv.Case2Camel(e.Name.Source()),
		Type:    strings.TrimPrefix(e.DataType, "u"),
		Comment: e.Comment,
		Tag: []field.Tag{
			{
				Name:  "json",
				Value: []string{e.Name.Source()},
			},
			{
				Name: "gorm",
				Value: []string{
					fmt.Sprintf("column:%s", e.Name.Source()),
				},
			},
		},
		Docs:     nil,
		IsInline: false,
	}
}
