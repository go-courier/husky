package cmd

import (
	"os"

	"github.com/go-courier/husky/husky/fmtx"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdLintCommit)
	CmdRoot.AddCommand(cmdLintStaged)
}

var cmdLintCommit = &cobra.Command{
	Use:   "lint-commit",
	Short: "lint commit msg",
	Run: func(cmd *cobra.Command, args []string) {
		catch(theHusky.RunLintCommit())
	},
}

var cmdLintStaged = &cobra.Command{
	Use:   "lint-staged",
	Short: "lint stated files",
	Run: func(cmd *cobra.Command, args []string) {
		catch(theHusky.RunLintStated())
	},
}

func catch(err error) {
	if err != nil {
		fmtx.Fprintln(os.Stderr, os.Stderr, err)
		os.Exit(1)
	}
}
