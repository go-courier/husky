package husky

import (
	"io/ioutil"
	"os"

	"github.com/go-courier/husky/pkg/lintcommit"
	"github.com/go-courier/husky/pkg/lintstaged"
	"gopkg.in/yaml.v2"
)

func HuskyFrom(huskyFile string) *Husky {
	data, err := ioutil.ReadFile(huskyFile)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	s := NewSpec()

	if err := yaml.Unmarshal(data, s); err != nil {
		panic(err)
	}

	h := &Husky{
		Spec: *s,
	}
	h.Init()

	return h
}

type Husky struct {
	Spec
	RunLintCommit func() error
	RunLintStated func() error
}

func (h *Husky) Init() {
	h.RunLintCommit = h.LintCommit.NewLint()
	h.RunLintStated = h.LintStaged.NewLint()
}

func NewSpec() *Spec {
	return &Spec{
		Hooks:      map[string][]string{},
		LintStaged: lintstaged.LintStaged{},
		LintCommit: lintcommit.LintCommit{},
	}
}

type Spec struct {
	Scripts    map[string][]string   `yaml:"scripts,omitempty"`
	Hooks      map[string][]string   `yaml:"hooks,omitempty"`
	LintStaged lintstaged.LintStaged `yaml:"lint-staged,omitempty"`
	LintCommit lintcommit.LintCommit `yaml:"lint-commit,omitempty"`
}
