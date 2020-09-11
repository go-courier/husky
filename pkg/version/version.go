package version

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/go-courier/husky/pkg/log"
	"github.com/go-logr/logr"

	"github.com/go-courier/husky/pkg/conventionalcommit"
	"github.com/go-courier/semver"
)

type verIncr int

const (
	verIncrPatch verIncr = iota
	verIncrMinor
	verIncrMajor
)

func incr(v *semver.Version, incr verIncr) *semver.Version {
	switch incr {
	case verIncrMajor:
		return v.IncrMajor()
	case verIncrMinor:
		return v.IncrMinor()
	default:
		return v.IncrPatch()
	}
}

func CalcNextVer(list []Commit, fromVersion *semver.Version) (*semver.Version, map[string][]*Commit) {
	sections := map[string][]*Commit{}
	verIncr := verIncrPatch
	ver := semver.MustParseVersion("0.0.0")

	add := func(cm *Commit) {
		sectionTitle := conventionalcommit.Types[cm.Type]
		if sectionTitle == "" {
			return
		}

		if cm.Type == "feat" {
			verIncr = verIncrMinor
		}

		if cm.BreakingChange {
			verIncr = verIncrMajor
		}

		sections[sectionTitle] = append(sections[sectionTitle], cm)
	}

	for i := range list {
		c := list[i]

		if c.CommitMsg == nil {
			continue
		}

		add(&c)
	}

	if fromVersion != nil {
		ver = incr(fromVersion, verIncr)
	}

	return ver, sections
}

func LastVersion() (ver *semver.Version, tag string, err error) {
	cmd := exec.Command("git", "tag", "--list", "--sort", "v:refname", "--merge", "HEAD")
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, "", err
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		s := bufio.NewScanner(stdoutPipe)

		for s.Scan() {
			t := s.Text()
			v, err := semver.ParseVersion(t)
			if err == nil {
				if v.Prerelease() == "" && v.Metadata() == "" {
					ver = v
					tag = t
				}
			}
		}
	}()

	if err := cmd.Run(); err != nil {
		return nil, "", err
	}

	wg.Wait()

	return
}

type VersionOpt struct {
	Prerelease string
	SkipPull   bool
	SkipCommit bool
	SkipTag    bool
	SkipPush   bool
}

func NewVersionAction(ctx context.Context, opt VersionOpt) *VersionAction {
	return &VersionAction{logger: log.LoggerFromContext(ctx).WithName("Version"), opt: opt}
}

type VersionAction struct {
	logger logr.Logger
	opt    VersionOpt
}

func (a *VersionAction) Do() error {
	ok, err := IsCleanWorkingDir()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("files should be committed before version")
	}

	ctx := log.WithLogger(a.logger)(context.Background())

	if !a.opt.SkipPull {
		if err := GitUpAll(ctx); err != nil {
			return err
		}
	}

	lastVer, commitList, err := ResolveVersionAndCommits()
	if err != nil {
		return err
	}

	a.logger.Info("calc version...")

	nextVer, sections := CalcNextVer(commitList, lastVer)

	// no change log when pre release
	if a.opt.Prerelease == "" {
		file, err := ReadOrTouchChangeLogFile()
		if err != nil {
			return err
		}
		a.logger.Info("updating changelog...")
		if err := UpdateChangeLog(file, nextVer, lastVer, sections); err != nil {
			return err
		}
	} else {
		v, err := nextVer.WithPrerelease(a.opt.Prerelease)
		if err != nil {
			return err
		}
		nextVer = v
	}

	if err := GitTagVersion(ctx, nextVer, a.opt.SkipCommit, a.opt.SkipTag); err != nil {
		return err
	}

	if !a.opt.SkipPush {
		return GitPushFollowTags(ctx)
	}

	return nil
}
