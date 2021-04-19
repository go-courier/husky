package main

import (
	"context"
	"flag"
	"path"

	"github.com/go-courier/husky/pkg/husky"
	"github.com/go-courier/husky/pkg/log"
	"github.com/go-courier/husky/internal/version"
	"github.com/go-logr/glogr"
	"github.com/spf13/cobra"
)

var (
	logger      = glogr.New().WithName("husky")
	projectRoot = husky.ResolveGitRoot()
	theHusky    = husky.HuskyFrom(log.WithLogger(logger)(context.Background()), path.Join(projectRoot, ".husky.yaml"))
)

var CmdRoot = &cobra.Command{
	Use:     "husky",
	Short:   "husky " + version.Version,
	Version: version.Version,
}

func init() {
	flag.Parse()
	CmdRoot.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
