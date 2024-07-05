/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tsCmd represents the ts command
var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ts called")
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tsCmd.PersistentFlags().StringVarP(&VarStringApiFile, "api-file", "f", "test.api", "sql文件")
	tsCmd.PersistentFlags().StringVarP(&VarStringTplPath, "tpl-path", "t", "test.tpl", "模板文件")
	tsCmd.PersistentFlags().StringVarP(&VarStringOutPath, "out-path", "o", "./", "输出路径")
	tsCmd.PersistentFlags().StringVarP(&VarStringNameAs, "name-as", "n", "%s.go", "输出名称")
}
