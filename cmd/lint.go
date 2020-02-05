package cmd

import (
	"os"

	"github.com/go-courier/husky/husky/fmtx"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdLint)

	cmdLint.AddCommand(cmdLintCommit)
	cmdLint.AddCommand(cmdLintStaged)
}

var cmdLint = &cobra.Command{
	Use: "lint",
}

var cmdLintCommit = &cobra.Command{
	Use: "commit",
	Run: func(cmd *cobra.Command, args []string) {
		catch(theHusky.RunLintCommit())
	},
}

var cmdLintStaged = &cobra.Command{
	Use: "staged",
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
