/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "生成 Model 模型代码",
	}

	cmd.AddCommand(NewModelDDLCmd())
	cmd.AddCommand(NewModelDSNCmd())

	return cmd
}
