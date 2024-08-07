/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
type ModelCmd struct {
	CMD     *cobra.Command
	SqlFile string
	TplFile string
	OutPath string

	Style string
}

func NewModelCmd() *ModelCmd {
	rootCmd := &ModelCmd{}
	rootCmd.CMD = &cobra.Command{
		Use:   "model",
		Short: "从sql文件生成go代码",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.RunCommand(cmd, args)
		},
	}

	rootCmd.CMD.AddCommand(NewModelDDLCmd().CMD)
	rootCmd.CMD.AddCommand(NewModelDSNCmd().CMD)
	rootCmd.init()
	return rootCmd
}

func (s *ModelCmd) init() {

}

func (s *ModelCmd) RunCommand(cmd *cobra.Command, args []string) {

}
