package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/go-courier/husky/husky"
	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/version"
	"github.com/spf13/cobra"
)

var (
	projectRoot = husky.ResolveGitRoot()
	theHusky    = husky.HuskyFrom(path.Join(projectRoot, ".husky.yaml"))
)

var CmdRoot = &cobra.Command{
	Use:     "husky",
	Version: version.Version,
}

func init() {
	Init(projectRoot)
}

func Init(root string) {
	githooks, _ := husky.ListGithookName(root)

	for _, githook := range githooks {
		ignore(ioutil.WriteFile(path.Join(root, ".git/hooks", githook), []byte(`#!/bin/sh

husky hook $(basename "$0") $*
`), os.ModePerm))
	}
}

func ignore(err error) {

}

func catch(err error) {
	if err != nil {
		fmtx.Fprintln(os.Stderr, os.Stderr, err)
		os.Exit(1)
	}
}
