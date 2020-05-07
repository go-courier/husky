package cmd

import (
	"os"

	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/version"
	"github.com/spf13/cobra"
)

var versionOpt = version.VersionOpt{}

func init() {
	CmdRoot.AddCommand(cmdVersion)

	cmdVersion.Flags().StringVarP(&versionOpt.Prerelease, "pre", "", "", "version with pre release. ex. alpha.0 beta.0")
	cmdVersion.Flags().BoolVarP(&versionOpt.SkipPull, "skip-pull", "", false, "skip pull")
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "auto version by conventional commit",
	Run: func(cmd *cobra.Command, args []string) {
		err := version.Version(versionOpt)
		if err != nil {
			fmtx.Fprintln(os.Stderr, err)
		}
	},
}
