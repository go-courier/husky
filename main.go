package main

import "github.com/go-courier/husky/cmd"

func main() {
	if err := cmd.CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
