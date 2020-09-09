package scripts

import (
	"context"
	"os"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

func RunScripts(scripts []string) error {
	for i := range scripts {
		if err := RunScript(scripts[i]); err != nil {
			return err
		}
	}
	return nil
}

func RunScript(script string) error {
	sh, err := syntax.NewParser().Parse(strings.NewReader(script), "")
	if err != nil {
		return err
	}

	runner, err := interp.New(
		interp.StdIO(os.Stdin, os.Stdout, os.Stderr),
	)

	if err != nil {
		return err
	}

	return runner.Run(context.Background(), sh)
}
