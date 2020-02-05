package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/scripts"
	"github.com/spf13/cobra"
)

var cmdHook = &cobra.Command{
	Use:   "hook",
	Short: "run hook <hookname>",
}

func init() {
	CmdRoot.AddCommand(cmdHook)

	for hook := range theHusky.Hooks {
		ss := theHusky.Hooks[hook]

		cmdHook.AddCommand(&cobra.Command{
			Use: hook,
			Run: func(cmd *cobra.Command, args []string) {
				fmtx.Fprintln(os.Stdout, color.YellowString(hook))
				catch(scripts.RunScripts(ss))
			},
		})
	}
}
