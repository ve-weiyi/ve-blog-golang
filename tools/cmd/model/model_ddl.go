/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/model/helper"
)

var ddlCfg = modelConfig{
	SqlFile: "test.sql",
	TplFile: "model.tpl",
	OutPath: "./",
	NameAs:  "%s.go",
}

func NewModelDDLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ddl",
		Short: "从 SQL DDL 文件生成 Model",
		RunE:  runModelDDL,
	}

	cmd.Flags().StringVarP(&ddlCfg.SqlFile, "sql-file", "s", ddlCfg.SqlFile, "sql文件")
	cmd.Flags().StringVarP(&ddlCfg.TplFile, "tpl-file", "t", ddlCfg.TplFile, "模板文件")
	cmd.Flags().StringVarP(&ddlCfg.OutPath, "out-path", "o", ddlCfg.OutPath, "输出路径")
	cmd.Flags().StringVarP(&ddlCfg.NameAs, "name-as", "n", ddlCfg.NameAs, "输出名称")

	return cmd
}

func runModelDDL(cmd *cobra.Command, args []string) error {
	log.Printf("sql-file: %s, tpl-file: %s, out-path: %s, name-as: %s\n",
		ddlCfg.SqlFile, ddlCfg.TplFile, ddlCfg.OutPath, ddlCfg.NameAs)

	tables, err := helper.ParseTableFromSql(ddlCfg.SqlFile)
	if err != nil {
		return err
	}

	return generateModel(tables, ddlCfg)
}
