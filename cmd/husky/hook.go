package main

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/go-courier/husky/pkg/log"
	"github.com/go-courier/husky/pkg/scripts"
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
				l := logger.WithName("HOOK").WithName(hook)
				ctx := log.WithLogger(l)(context.Background())

				if err := scripts.RunScripts(ctx, ss); err != nil {
					color.Red(err.Error())
					os.Exit(1)
				}
			}
		}
	},
}
