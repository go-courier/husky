package lintstaged

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/go-courier/husky/pkg/fmtx"
	"github.com/go-courier/husky/pkg/scripts"
	"github.com/gobwas/glob"
)

type LintStaged map[string][]string

func (lintStaged LintStaged) NewLint() func() error {
	testers := make([]tester, 0)

	for g, scripts := range lintStaged {
		testers = append(testers, tester{
			glob:    glob.MustCompile(g),
			scripts: scripts,
		})
	}

	return func() error {
		files, err := listStagedFile()
		if err != nil {
			return err
		}

		if len(files) == 0 {
			fmtx.TopicFprintln("lint-staged", os.Stdout, "No staged files found.")
			return nil
		}

		matchedFiles := make([][]string, len(testers))

		for _, f := range files {
			for i, tester := range testers {
				if tester.glob.Match(f) {
					matchedFiles[i] = append(matchedFiles[i], f)
				}
			}
		}

		for i, tester := range testers {
			if len(matchedFiles[i]) > 0 {
				for _, s := range tester.scripts {
					fmtx.TopicFprintln("lint-staged", os.Stdout, color.MagentaString(s))

					for _, f := range matchedFiles[i] {
						if err := scripts.RunScript(s + " " + f); err != nil {
							return err
						}
						if err := scripts.RunScript("git add " + f); err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}
}

type tester struct {
	glob    glob.Glob
	scripts []string
}

func listStagedFile() ([]string, error) {
	ret, err := exec.Command("git", "diff", "--staged", "--diff-filter=ACMR", "--name-only").CombinedOutput()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes.TrimSpace(ret)), "\n"), nil
}
