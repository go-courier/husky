package main

import (
	"path"

	"github.com/go-courier/husky/cmd/husky/presets"
	"github.com/go-courier/husky/pkg/husky"
	"github.com/spf13/cobra"
)

func init() {
	var cmdPreset = &cobra.Command{
		Use:   "preset",
		Short: "preset for setup",
	}

	CmdRoot.AddCommand(cmdPreset)

	for name := range presets.Presets {
		p := presets.Presets[name]

		c := &cobra.Command{
			Use: name,
			Run: func(cmd *cobra.Command, args []string) {
				for f, data := range p {
					err := husky.WriteFile(path.Join(projectRoot, f), data)
					if err != nil {
						panic(err)
					}
				}
			},
		}

		cmdPreset.AddCommand(c)
	}
}
