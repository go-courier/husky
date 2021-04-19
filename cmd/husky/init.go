package main

import (
	"os"
	"path"

	"github.com/go-courier/husky/pkg/husky"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdInit)
}

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "init git hooks",
	RunE: func(cmd *cobra.Command, args []string) error {

		githooks, err := husky.ListGithookName(projectRoot)
		if err != nil {
			return err
		}

		for _, githook := range githooks {
			if err := os.WriteFile(path.Join(projectRoot, ".git/hooks", githook), []byte(`#!/bin/sh

husky hook $(basename "$0") $*
`), os.ModePerm); err != nil {
				return err
			}
		}
		return nil
	},
}
