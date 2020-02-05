package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/scripts"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdRun)
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run script",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			scriptName := args[0]

			if ss, ok := theHusky.Scripts[scriptName]; ok {
				fmtx.Fprintln(os.Stdout, color.YellowString(scriptName))
				catch(scripts.RunScripts(ss))
			}
		}
	},
}
