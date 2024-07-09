/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
)

// migrateCmd represents the migrate command
type ModelDDLCmd struct {
	CMD *cobra.Command
	modelConfig
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

	var tables []*Table
	var err error

	tables, err = ParseTableFromSql(s.SqlFile)
	if err != nil {
		log.Fatal(err)
	}

	err = generateModel(tables, s.modelConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// 从sql文件解析Table
func ParseTableFromSql(sql string) (list []*Table, err error) {
	n := strings.TrimRight(sql, ".sql")

	f := files.ToAbs(sql)
	tables, err := parser.Parse(f, n, false)
	if err != nil {
		return nil, err
	}

	for _, table := range tables {

		fs := make([]*Field, 0)
		for _, field := range table.Fields {
			f := convertFieldToField(field)
			fs = append(fs, &f)
		}

		ufs := make(map[string][]*Field)
		for k, index := range table.UniqueIndex {
			uf := make([]*Field, 0)
			for _, field := range index {
				f := convertFieldToField(field)
				uf = append(uf, &f)
			}
			ufs[k] = uf
		}

		v := &Table{
			Name: table.Name.Source(),
			Db:   table.Db.Source(),
			PrimaryKey: Primary{
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

func convertFieldToField(col *parser.Field) Field {
	f := Field{
		Name:            col.Name.Source(),
		DataType:        col.DataType,
		Comment:         col.Comment,
		SeqInIndex:      col.SeqInIndex,
		OrdinalPosition: col.OrdinalPosition,
	}

	return f
}
