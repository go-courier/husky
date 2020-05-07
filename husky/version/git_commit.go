package version

import (
	"os/exec"
	"time"

	"github.com/go-courier/husky/husky/conventionalcommit"
	"gopkg.in/yaml.v2"
)

type CommitMsg conventionalcommit.CommitMsg

func (v CommitMsg) MarshalText() ([]byte, error) {
	return []byte(conventionalcommit.CommitMsg(v).String()), nil
}

func (v *CommitMsg) UnmarshalText(data []byte) error {
	commitMsg, err := conventionalcommit.ParseCommitMsg(string(data))
	if err != nil {
		return nil
	}
	*v = CommitMsg(*commitMsg)
	return nil
}

type Commit struct {
	Hash   string `yaml:"hash"`
	Author struct {
		Name       string    `yaml:"name,omitempty"`
		Email      string    `yaml:"email,omitempty"`
		AuthoredAt time.Time `yaml:"authoredAt,omitempty"`
	} `yaml:"author,omitempty"`
	Committer struct {
		Name        string    `yaml:"name,omitempty"`
		Email       string    `yaml:"email,omitempty"`
		CommittedAt time.Time `yaml:"committedAt,omitempty"`
	} `yaml:"committer,omitempty"`
	*CommitMsg `yaml:"msg,omitempty"`
}

func ListCommit(from string) ([]Commit, error) {
	args := []string{
		"log", "--date=short", "--no-decorate", "--pretty=" + `
- hash: %H
  author: 
    name: %an
    email: %ae
    authoredAt: %aI
  committer:
    name: %cn
    email: %ce
    committedAt: %cI
  msg: |
    %s
    
    %b
`,
	}

	if from != "" {
		args = append(args, from+"..HEAD")
	}

	ret, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	commits := make([]Commit, 0)
	if err := yaml.Unmarshal(ret, &commits); err != nil {
		return nil, err
	}
	return commits, err
}
