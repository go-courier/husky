package version

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/go-courier/husky/pkg/husky"
	"github.com/go-courier/semver"
)

func ReadOrTouchChangeLogFile() (*os.File, error) {
	changelogFile := filepath.Join(husky.ResolveGitRoot(), "CHANGELOG.md")
	return openOrCreateFile(changelogFile)
}

func ResolveVersionAndCommits() (*semver.Version, []Commit, error) {
	lastVersion, tag, err := LastVersion()
	if err != nil {
		return nil, nil, err
	}

	commitList, err := ListCommit(tag)
	if err != nil {
		return nil, nil, err
	}

	if len(commitList) == 0 {
		return nil, nil, fmt.Errorf("nothing to version")
	}

	return lastVersion, commitList, nil
}

func UpdateChangeLog(file interface {
	io.Reader
	io.Writer
}, nextVer *semver.Version, fromVersion *semver.Version, sections map[string][]*Commit) error {
	list, err := scanExistsChangelogSection(file)
	if err != nil {
		return err
	}

	if err := Truncate(file); err != nil {
		return err
	}

	writeString(file, `# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.
`)

	baseURI, err := resolveBaseURI()

	if err != nil {
		return err
	}

	if err := writeChangelogSection(file, nextVer, fromVersion, sections, baseURI); err != nil {
		return err
	}

	for i := range list {
		s := list[i]

		if s.Version.String() != nextVer.String() {
			writeChangelogSectionSplit(file)
			writeString(file, strings.TrimSpace(s.Lines.String())+"\n")
		}
	}

	return nil
}

type changelogSection struct {
	Version *semver.Version
	Lines   *bytes.Buffer
}

var regVer = regexp.MustCompile(`# \[?([^]]+)\]?`)

func scanExistsChangelogSection(r io.Reader) ([]*changelogSection, error) {
	scanner := bufio.NewScanner(r)

	sections := make([]*changelogSection, 0)

	ver := (*semver.Version)(nil)
	lines := bytes.NewBuffer(nil)

	wrap := func() {
		if ver != nil {
			sections = append(sections, &changelogSection{
				Version: ver,
				Lines:   lines,
			})
		}

		ver = nil
		lines = bytes.NewBuffer(nil)
	}

	for scanner.Scan() {
		line := scanner.Text()

		matches := regVer.FindAllStringSubmatch(line, -1)

		if len(matches) > 0 {
			verStr := matches[0][len(matches[0])-1]

			v, err := semver.ParseVersion(verStr)
			if err == nil {
				wrap()
				ver = v
				lines.WriteString(line + "\n")
				continue
			}
		}

		lines.WriteString(line + "\n")
	}

	wrap()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}

func writeChangelogSection(w io.Writer, nextVer *semver.Version, fromVersion *semver.Version, sections map[string][]*Commit, baseURI string) error {
	titles := make([]string, 0)

	for title := range sections {
		titles = append(titles, title)
	}

	sort.Strings(titles)

	writeChangelogSectionSplit(w)

	writeString(w, "# ")

	if fromVersion == nil {
		writeString(w, nextVer.String())
	} else {
		writeString(w, "[")
		writeString(w, nextVer.String())
		writeString(w, "](")
		writeString(w, baseURI+"/compare/v"+fromVersion.String()+"..."+"v"+nextVer.String())
		writeString(w, ")")
	}

	for _, title := range titles {
		writeString(w, "\n\n")
		writeString(w, "### "+title)
		writeString(w, "\n\n")

		for _, v := range sections[title] {
			if err := writeChangelog(w, v, baseURI); err != nil {
				return err
			}
		}
	}

	return nil
}

func openOrCreateFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return os.Create(filename)
		}
		return nil, err
	}
	return os.OpenFile(filename, os.O_RDWR, os.ModePerm)
}

func writeChangelog(w io.Writer, cm *Commit, baseURI string) error {
	writeString(w, "* ")
	writeString(w, "**")
	writeString(w, cm.Type)

	if cm.Scope != "" {
		writeString(w, "(")
		writeString(w, cm.Scope)
		writeString(w, "):")
	}

	writeString(w, "** ")

	if cm.BreakingChange {
		writeString(w, "**BREAKING CHANGE** ")
	}

	writeString(w, cm.Header)

	writeString(w, " ([")
	writeString(w, cm.Hash[0:7])
	writeString(w, "](")
	writeString(w, baseURI+"/commit/"+cm.Hash)
	writeString(w, "))")
	writeString(w, "\n")

	return nil
}

func writeChangelogSectionSplit(w io.Writer) {
	writeString(w, "\n\n\n")
}

func writeString(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	ignore(err)
}
