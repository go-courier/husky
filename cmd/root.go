package cmd

import (
	"path"

	"github.com/go-courier/husky/husky"
	"github.com/spf13/cobra"
)

var (
	gitRoot  = husky.ResolveGitRoot()
	theHusky = husky.HuskyFrom(path.Join(gitRoot, ".husky.yaml"))
)

var CmdRoot = &cobra.Command{
	Use: "husky",
}
