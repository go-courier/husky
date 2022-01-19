package scripts

import (
	"bytes"
	"context"
	"os"
	"strings"
	"text/template"

	"github.com/go-courier/husky/pkg/log"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

func ParesScriptTemplate(s string, data interface{}) (string, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return "", err
	}
	var buf = bytes.NewBuffer(nil)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RunScripts(ctx context.Context, scripts []string) error {
	for i := range scripts {
		if err := RunScript(ctx, scripts[i]); err != nil {
			return err
		}
	}
	return nil
}

func RunScript(ctx context.Context, script string) (e error) {
	logger := log.LoggerFromContext(ctx).WithName("RunScript")

	logger.V(1).Info(script)

	sh, err := syntax.NewParser().Parse(strings.NewReader(script), "")
	if err != nil {
		return err
	}

	runner, err := interp.New(interp.StdIO(os.Stdin, os.Stdout, os.Stderr))
	if err != nil {
		return err
	}

	return runner.Run(context.Background(), sh)
}
