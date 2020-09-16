package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdLintCommit)
}

var cmdLintCommit = &cobra.Command{
	Use:   "lint-commit",
	Short: "lint commit msg",
	Run: func(cmd *cobra.Command, args []string) {
		if err := theHusky.RunLintCommit(); err != nil {
			logger.Error(err, "failed")
			os.Exit(1)
		}
	},
}
