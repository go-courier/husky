package version

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"

	"github.com/go-courier/husky/pkg/husky"

	"github.com/go-courier/husky/pkg/scripts"
	"github.com/go-courier/semver"
)

func resolveBaseURI() (string, error) {
	gitURI, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
	if err != nil {
		return "", err
	}
	return getBaseURI(string(bytes.TrimSpace(gitURI))), nil
}

var regGit = regexp.MustCompile(`^(git@|(https?://))(?P<host>[^:/]+)[:/]?(?P<repoName>.+)\.git$`)

func getBaseURI(gitURI string) string {
	matched := regGit.FindAllStringSubmatch(gitURI, -1)
	if len(matched) == 0 {
		return ""
	}

	host := ""
	repoName := ""

	for i, n := range regGit.SubexpNames() {
		v := matched[0][i]

		switch n {
		case "host":
			host = v
		case "repoName":
			repoName = v
		}
	}

	return "https://" + host + "/" + repoName
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

func GitUpAll(ctx context.Context) error {
	return scripts.RunScript(ctx, "git pull --rebase && git pull --tags --force")
}

func GitPushFollowTags(ctx context.Context) error {
	return scripts.RunScript(ctx, "git push --follow-tags")
}

func GitTagVersion(ctx context.Context, ver *semver.Version, skipCommit bool, skipTag bool, versionFile string) error {
	_ = husky.WriteFile(versionFile, []byte(ver.String()))

	if skipCommit {
		return nil
	}

	_ = scripts.RunScript(ctx, fmt.Sprintf(`git add . && git commit --no-verify --message "chore(release): v%s"`, ver))

	if skipTag {
		return nil
	}

	return scripts.RunScript(ctx, fmt.Sprintf(`git tag --force --annotate v%s --message "v%s"`, ver, ver))
}
