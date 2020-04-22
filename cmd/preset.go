package cmd

import (
	"path"

	"github.com/go-courier/husky/husky"
	"github.com/go-courier/husky/presets"
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
