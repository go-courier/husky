package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/scripts"
	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run script",
}

func init() {
	CmdRoot.AddCommand(cmdRun)

	for name := range theHusky.Scripts {
		n := name
		ss := theHusky.Scripts[n]

		c := &cobra.Command{
			Use:   n,
			Short: strings.Join(ss, " && "),
			Run: func(cmd *cobra.Command, args []string) {

				fmtx.Fprintln(os.Stdout, color.YellowString(n))
				catch(scripts.RunScripts(ss))
			},
		}

		cmdRun.AddCommand(c)
		CmdRoot.AddCommand(c)
	}
}
