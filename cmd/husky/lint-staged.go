package main

import (
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdLintStaged)
}

var cmdLintStaged = &cobra.Command{
	Use:   "lint-staged",
	Short: "lint stated files",
	Run: func(cmd *cobra.Command, args []string) {
		catch(theHusky.RunLintStated())
	},
}
