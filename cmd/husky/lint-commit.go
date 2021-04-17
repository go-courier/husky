package main

import (
	"os"

	"github.com/fatih/color"
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
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}
