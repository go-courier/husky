package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/scripts"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdHook)
}

var cmdHook = &cobra.Command{
	Use:   "hook",
	Short: "run hook <hookname>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			hook := args[0]

			if ss, ok := theHusky.Hooks[hook]; ok {
				fmtx.Fprintln(os.Stdout, color.YellowString(hook))
				catch(scripts.RunScripts(ss))
			}
		}
	},
}
