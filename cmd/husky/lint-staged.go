package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdLintStaged)
}

var cmdLintStaged = &cobra.Command{
	Use:   "lint-staged",
	Short: "lint stated files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := theHusky.RunLintStated(); err != nil {
			logger.Error(err, "failed.")
			os.Exit(1)
		}
	},
}
