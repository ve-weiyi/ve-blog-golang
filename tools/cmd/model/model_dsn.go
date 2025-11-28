/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/model/helper"
)

var dsnCfg = modelConfig{
	SqlFile: "root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local",
	TplFile: "entity.tpl",
	OutPath: "./",
	NameAs:  "%s.go",
}

func NewModelDSNCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dsn",
		Short: "从数据库连接生成 Model",
		RunE:  runModelDSN,
	}

	cmd.Flags().StringVarP(&dsnCfg.SqlFile, "source", "s", dsnCfg.SqlFile, "数据库地址")
	cmd.Flags().StringVarP(&dsnCfg.TplFile, "tpl-file", "t", dsnCfg.TplFile, "模板文件")
	cmd.Flags().StringVarP(&dsnCfg.OutPath, "out-path", "o", dsnCfg.OutPath, "输出路径")
	cmd.Flags().StringVarP(&dsnCfg.NameAs, "name-as", "n", dsnCfg.NameAs, "输出名称")

	return cmd
}

func runModelDSN(cmd *cobra.Command, args []string) error {
	log.Printf("source: %s, tpl-file: %s, out-path: %s, name-as: %s\n",
		dsnCfg.SqlFile, dsnCfg.TplFile, dsnCfg.OutPath, dsnCfg.NameAs)

	tables, err := helper.ParseTableFromDsn(dsnCfg.SqlFile)
	if err != nil {
		return err
	}

	return generateModel(tables, dsnCfg)
}
