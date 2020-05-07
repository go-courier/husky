package main

import (
	"github.com/go-courier/husky/cmd"
)

var Version string

func main() {
	cmd.SetVersion(Version)

	if err := cmd.CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
