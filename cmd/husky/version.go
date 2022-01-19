package main

import (
	"context"
	"os"

	"github.com/go-courier/husky/pkg/scripts"

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
		ctx := log.WithLogger(logger.WithName(cmd.Use))(context.Background())

		versionOpt.VersionFile = theHusky.VersionFile
		versionOpt.PostVersion = func(version string) error {
			if ss, ok := theHusky.Hooks["post-version"]; ok {
				for _, s := range ss {
					parsed, err := scripts.ParesScriptTemplate(s, map[string]string{
						"Version": version,
					})
					if err != nil {
						return err
					}
					if err := scripts.RunScript(ctx, parsed); err != nil {
						return err
					}
				}
			}
			return nil
		}

		err := version.NewVersionAction(ctx, versionOpt).Do()
		if err != nil {
			logger.Error(err, "")
			os.Exit(1)
		}
	},
}
