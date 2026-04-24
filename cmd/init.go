package cmd

import (
	"github.com/basti1302/go-husky/internal/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize husky",
	Long: `
husky is a tool to help you manage your git hooks.

For more information, please visit
https://github.com/basti1302/go-husky
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Init(); err != nil {
			exitOnError(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
