/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
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

func ParseApiSpec(filename string) (out *spec.ApiSpec, err error) {
	if path.IsAbs(filename) {
		return parser.Parse(filename)
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f := path.Join(dir, filename)
	return parser.Parse(f)
}
