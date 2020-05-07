package version

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"regexp"

	"github.com/go-courier/husky/husky/fmtx"
	"github.com/go-courier/husky/husky/scripts"
	"github.com/go-courier/semver"
)

func ResolveRepoName() (string, error) {
	gitURI, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
	if err != nil {
		return "", err
	}
	return GetRepoName(string(bytes.TrimSpace(gitURI))), nil
}

var regGit = regexp.MustCompile(`^((git@([^/]+):)|(https?://([^/]+)/))(.+).git$`)

func GetRepoName(gitURI string) string {
	ss := regGit.FindAllStringSubmatch(gitURI, -1)
	if len(ss) == 0 {
		return ""
	}
	return ss[0][len(ss[0])-1]
}

func Truncate(v interface{}) error {
	if f, ok := v.(interface {
		Truncate(size int) error
	}); ok {
		if err := f.Truncate(0); err != nil {
			return err
		}
	}

	if seeker, ok := v.(io.Seeker); ok {
		if _, err := seeker.Seek(0, 0); err != nil {
			return err
		}
	}
	return nil
}

func IsCleanWorkingDir() (bool, error) {
	ret, err := exec.Command("git", "status", "-s").CombinedOutput()
	if err != nil {
		return false, err
	}
	return len(bytes.TrimSpace(ret)) == 0, nil
}

func GitUpAll() error {
	return scripts.StdRun("git pull --rebase && git pull --tags")
}

func GitTagVersion(ver *semver.Version) error {
	v := ver.String()
	defer fmtx.Fprintln(os.Stdout, v)
	return scripts.StdRun(`git add . && git commit --no-verify -m "chore(release): v` + v + `" && git tag --force v` + v)
}
