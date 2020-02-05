package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/go-courier/husky/husky"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(cmdInit)
}

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "init githooks",
	Run: func(cmd *cobra.Command, args []string) {
		Init(gitRoot)
	},
}

func Init(root string) {
	githooks, _ := husky.ListGithookName(root)

	for _, githook := range githooks {
		ioutil.WriteFile(path.Join(root, ".git/hooks", githook), []byte(`#!/bin/sh

husky run $(basename "$0") $*
`), os.ModePerm)
	}
}
