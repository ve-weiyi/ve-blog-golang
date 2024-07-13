/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import (
	"log"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/convertx"
)

// migrateCmd represents the migrate command
type ModelDSNCmd struct {
	CMD *cobra.Command
	modelConfig
}

func NewModelDSNCmd() *ModelDSNCmd {
	rootCmd := &ModelDSNCmd{}
	rootCmd.CMD = &cobra.Command{
		Use: "dsn",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.RunCommand(cmd, args)
		},
	}

	rootCmd.init()
	return rootCmd
}

func (s *ModelDSNCmd) init() {
	dsn := "root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local"
	s.CMD.PersistentFlags().StringVarP(&s.SqlFile, "source", "s", dsn, "数据库地址")
	s.CMD.PersistentFlags().StringVarP(&s.TplFile, "tpl-file", "t", "entity.tpl", "模板文件")
	s.CMD.PersistentFlags().StringVarP(&s.OutPath, "out-path", "o", "./", "输出路径")
	s.CMD.PersistentFlags().StringVarP(&s.NameAs, "name-as", "n", "%s.go", "输出名称")
}

func (s *ModelDSNCmd) RunCommand(cmd *cobra.Command, args []string) {
	log.Println("run model ddl")
	log.Println("sql-file:", s.SqlFile)
	log.Println("tpl-file:", s.TplFile)
	log.Println("out-path:", s.OutPath)
	log.Println("name-as:", s.NameAs)

	var tables []*Table
	var err error

	tables, err = ParseTableFromDsn(s.SqlFile)
	if err != nil {
		panic(err)
	}

	err = generateModel(tables, s.modelConfig)
	if err != nil {
		panic(err)
	}
}

// 从数据库中解析Table
func ParseTableFromDsn(dsn string) (list []*Table, err error) {
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		return nil, err
	}

	dbName := db.Migrator().CurrentDatabase()
	tables, err := db.Migrator().GetTables()
	for _, tableName := range tables {
		types, err := db.Migrator().ColumnTypes(tableName)
		if err != nil {
			return nil, err
		}

		indexes, err := db.Migrator().GetIndexes(tableName)
		if err != nil {
			return nil, err
		}

		pm := Primary{}
		for _, entity := range types {
			is, ok := entity.PrimaryKey()
			if ok && is {
				pm.Field = convertColumnToField(entity)
				pm.AutoIncrement, _ = entity.AutoIncrement()
			}
		}

		fs := make([]*Field, 0)
		for _, entity := range types {
			f := convertColumnToField(entity)
			fs = append(fs, &f)
		}

		ufs := make(map[string][]*Field)
		for k, index := range GroupByColumn(indexes) {
			uf := make([]*Field, 0)
			for _, field := range index {
				for _, entity := range types {
					if entity.Name() == field {
						f := convertColumnToField(entity)
						uf = append(uf, &f)
					}
				}
			}
			ufs[k] = uf
		}

		v := &Table{
			Name:        tableName,
			Db:          dbName,
			PrimaryKey:  pm,
			UniqueIndex: ufs,
			Fields:      fs,
		}

		list = append(list, v)
	}
	return list, nil
}

func convertColumnToField(col gorm.ColumnType) Field {
	f := Field{}

	f.Name = col.Name()
	f.Comment, _ = col.Comment()
	f.DataType = convertx.ConvertMysqlToGoType(col.DatabaseTypeName())

	// col.DatabaseTypeName() int
	// col.ColumnType() int unsigned
	return f
}

// GroupByColumn group columns
func GroupByColumn(indexList []gorm.Index) map[string][]string {

	ufs := make(map[string][]string)
	if len(indexList) == 0 {
		return ufs
	}

	for _, idx := range indexList {
		if idx == nil {
			continue
		}
		is, ok := idx.PrimaryKey()
		if ok && is {
			continue
		}

		is, ok = idx.Unique()
		if ok && is {
			name := idx.Name()
			for _, col := range idx.Columns() {
				ufs[name] = append(ufs[name], col)
			}
		}
	}
	return ufs
}
