package web

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "web",
		Short: "生成 Web 前端代码",
	}

	ParseFlagVar(cmd)
	cmd.AddCommand(typescriptCmd)

	return cmd
}
