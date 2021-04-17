package main

import (
	"os"

	"github.com/fatih/color"
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
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}
