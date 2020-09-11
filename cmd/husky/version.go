package main

import (
	"context"

	"github.com/go-courier/husky/pkg/log"
	"github.com/go-courier/husky/pkg/version"
	"github.com/spf13/cobra"
)

var versionOpt = version.VersionOpt{}

func init() {
	CmdRoot.AddCommand(cmdVersion)

	cmdVersion.Flags().StringVarP(&versionOpt.Prerelease, "pre", "", "", "version with pre release. ex. alpha.0 beta.0")
	cmdVersion.Flags().BoolVarP(&versionOpt.SkipPull, "skip-pull", "", false, "skip pull")
	cmdVersion.Flags().BoolVarP(&versionOpt.SkipCommit, "skip-commit", "", false, "skip commit")
	cmdVersion.Flags().BoolVarP(&versionOpt.SkipTag, "skip-tag", "", false, "skip tag")
	cmdVersion.Flags().BoolVarP(&versionOpt.SkipPush, "skip-push", "", false, "skip push")
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "auto version by conventional commit",
	Run: func(cmd *cobra.Command, args []string) {
		err := version.NewVersionAction(log.WithLogger(logger.WithName(cmd.Use))(context.Background()), versionOpt).Do()
		if err != nil {
			logger.Error(err, "")
		}
	},
}
