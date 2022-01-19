package husky

import (
	"context"
	"os"
	"path"
	"path/filepath"

	"github.com/go-courier/husky/pkg/lintcommit"
	"github.com/go-courier/husky/pkg/lintstaged"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

func HuskyFrom(ctx context.Context, projectRoot string) *Husky {
	files := []string{
		path.Join(projectRoot, ".husky.toml"),
		path.Join(projectRoot, ".husky.yaml"),
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			panic(err)
		}

		s := NewSpec()

		switch filepath.Ext(f) {
		case ".toml":
			if err := toml.Unmarshal(data, s); err != nil {
				panic(err)
			}
		case ".yaml":
			if err := yaml.Unmarshal(data, s); err != nil {
				panic(err)
			}
		}

		if s.VersionFile == "" {
			s.VersionFile = ".version"
		}

		h := &Husky{Spec: *s}

		h.Init(ctx)

		return h
	}

	panic("missing .husky{.toml,.yaml}")
}

type Husky struct {
	Spec
	RunLintCommit func() error
	RunLintStated func() error
}

func (h *Husky) Init(ctx context.Context) {
	h.RunLintCommit = h.LintCommit.NewLint(ctx)
	h.RunLintStated = h.LintStaged.NewLint(ctx)
}

func NewSpec() *Spec {
	s := &Spec{
		Hooks:      map[string][]string{},
		LintStaged: lintstaged.LintStaged{},
		LintCommit: lintcommit.LintCommit{},
	}
	return s
}

type Spec struct {
	VersionFile string                `toml:"version-file,omitempty" yaml:"version-file,omitempty"`
	Hooks       map[string][]string   `toml:"hooks,omitempty" yaml:"hooks,omitempty"`
	LintStaged  lintstaged.LintStaged `toml:"lint-staged,omitempty" yaml:"lint-staged,omitempty"`
	LintCommit  lintcommit.LintCommit `toml:"lint-commit,omitempty" yaml:"lint-commit,omitempty"`
}
