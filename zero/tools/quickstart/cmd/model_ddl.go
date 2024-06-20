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
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"

	"github.com/ve-weiyi/ve-blog-golang/kit/tools/field"
	"github.com/ve-weiyi/ve-blog-golang/kit/tools/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/tools/quickstart/cmd/api"
)

// migrateCmd represents the migrate command
type ModelDDLCmd struct {
	CMD     *cobra.Command
	SqlFile string
	TplFile string
	OutPath string

	NameAs string
}

func NewModelDDLCmd() *ModelDDLCmd {
	rootCmd := &ModelDDLCmd{}
	rootCmd.CMD = &cobra.Command{
		Use: "ddl",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.RunCommand(cmd, args)
		},
	}

	rootCmd.init()
	return rootCmd
}

func (s *ModelDDLCmd) init() {
	s.CMD.PersistentFlags().StringVarP(&s.SqlFile, "sql-file", "s", "test.sql", "sql文件")
	s.CMD.PersistentFlags().StringVarP(&s.TplFile, "tpl-file", "t", "model.tpl", "模板文件")
	s.CMD.PersistentFlags().StringVarP(&s.OutPath, "out-path", "o", "./", "输出路径")
	s.CMD.PersistentFlags().StringVarP(&s.NameAs, "name-as", "n", "%s.go", "输出名称")
}

func (s *ModelDDLCmd) RunCommand(cmd *cobra.Command, args []string) {
	log.Println("run model ddl")
	log.Println("sql-file:", s.SqlFile)
	log.Println("tpl-file:", s.TplFile)
	log.Println("out-path:", s.OutPath)
	log.Println("name-as:", s.NameAs)

	var metas []invent.TemplateMeta
	var tables []*api.Table
	var err error

	f := s.SqlFile
	t := s.TplFile
	o := s.OutPath
	n := s.NameAs

	tables, err = ParseTableFromSql(f)
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := os.ReadFile(t)
	if err != nil {
		log.Fatal(err)
	}

	for _, table := range tables {
		fmt.Printf("%+v\n", table.Name)

		data := convertTableToData(table)

		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(o, fmt.Sprintf(n, table.Name)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"funcFieldsKey": func(fs []*field.Field) string {
					var name string
					for _, ff := range fs {
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
		}
		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			log.Println(err)
		}
	}
}

// 从sql文件解析Table
func ParseTableFromSql(sql string) (list []*api.Table, err error) {
	n := strings.TrimRight(sql, ".sql")

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f := path.Join(dir, sql)
	tables, err := parser.Parse(f, n, false)
	if err != nil {
		return nil, err
	}

	for _, table := range tables {

		fs := make([]*api.Field, 0)
		for _, field := range table.Fields {
			f := convertFieldToField(field)
			fs = append(fs, &f)
		}

		ufs := make(map[string][]*api.Field)
		for k, index := range table.UniqueIndex {
			uf := make([]*api.Field, 0)
			for _, field := range index {
				f := convertFieldToField(field)
				uf = append(uf, &f)
			}
			ufs[k] = uf
		}

		v := &api.Table{
			Name: table.Name.Source(),
			Db:   table.Db.Source(),
			PrimaryKey: api.Primary{
				AutoIncrement: table.PrimaryKey.AutoIncrement,
				Field:         convertFieldToField(&table.PrimaryKey.Field),
			},
			UniqueIndex: ufs,
			Fields:      fs,
		}

		list = append(list, v)

	}

	return list, nil
}

func convertTableToData(table *api.Table) any {

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
		"TableName":           table.Name,
		"UpperStartCamelName": jsonconv.Case2Camel(table.Name),
		"LowerStartCamelName": jsonconv.Case2CamelLowerStart(table.Name),
		"SnakeName":           jsonconv.Case2Snake(table.Name),
		"Fields":              fs,
		"UniqueFields":        ufs,
	}

	return data
}

func convertField(e *api.Field) *field.Field {

	return &field.Field{
		Name:    jsonconv.Case2Camel(e.Name),
		Type:    strings.TrimPrefix(e.DataType, "u"),
		Comment: e.Comment,
		Tag: []field.Tag{
			{
				Name: "gorm",
				Value: []string{
					fmt.Sprintf("column:%s", e.Name),
				},
			},
			{
				Name:  "json",
				Value: []string{e.Name},
			},
		},
		Docs:     nil,
		IsInline: false,
	}
}

func convertFieldToField(col *parser.Field) api.Field {
	f := api.Field{
		Name:            col.Name.Source(),
		DataType:        col.DataType,
		Comment:         col.Comment,
		SeqInIndex:      col.SeqInIndex,
		OrdinalPosition: col.OrdinalPosition,
	}

	return f
}
