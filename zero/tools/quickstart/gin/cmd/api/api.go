/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
type ApiCmd struct {
	CMD     *cobra.Command
	SqlFile string
	TplFile string
	OutPath string

	Style string
}

func NewApiCmd() *ApiCmd {
	rootCmd := &ApiCmd{}
	rootCmd.CMD = &cobra.Command{
		Use:   "api",
		Short: "从api文件生成",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.RunCommand(cmd, args)
		},
	}

	rootCmd.CMD.AddCommand(apiAllCmd)
	rootCmd.CMD.AddCommand(routerCmd)
	rootCmd.CMD.AddCommand(serviceCmd)
	rootCmd.CMD.AddCommand(controllerCmd)
	rootCmd.init()
	return rootCmd
}

func (s *ApiCmd) init() {

}

func (s *ApiCmd) RunCommand(cmd *cobra.Command, args []string) {

}
