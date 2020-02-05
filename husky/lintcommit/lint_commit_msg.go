package lintcommit

import (
	"fmt"
	"regexp"
	"strings"
)

// https://www.conventionalcommits.org/en/v1.0.0-beta.2/

var (
	reHeader    = regexp.MustCompile("^(?P<type>\\w+)(\\((?P<scope>\\w+)\\))?:( +)?(\\[(?P<rel>.+)\\])?(?P<header>.+)")
	reHeaderRel = regexp.MustCompile("([A-Z]+)-([0-9]+)")

	types = []string{
		"build",
		"chore",
		"ci",
		"docs",
		"feat",
		"fix",
		"perf",
		"refactor",
		"revert",
		"style",
		"test",
	}
)

func validateType(tpe string) error {
	for _, t := range types {
		if tpe == t {
			return nil
		}
	}

	return fmt.Errorf("invalid type `%s`, type should be one of %v", tpe, types)
}

func validateRel(rel string) error {
	if reHeaderRel.MatchString(rel) {
		return nil
	}
	return fmt.Errorf("invalid rel", )
}

type CommitHeader struct {
	Type   string
	Scope  string
	Rel    string
	Header string
}

func parseCommitHeader(commitHeader string) (*CommitHeader, error) {
	if !reHeader.MatchString(commitHeader) {
		return nil, fmt.Errorf("invalid header format `%s`, should be %s", commitHeader, reHeader.String())
	}

	groupNames := reHeader.SubexpNames()

	ch := &CommitHeader{}

	for _, match := range reHeader.FindAllStringSubmatch(commitHeader, -1) {
		for groupIdx, value := range match {
			name := groupNames[groupIdx]
			switch name {
			case "type":
				ch.Type = strings.TrimSpace(value)
			case "scope":
				ch.Scope = strings.TrimSpace(value)
			case "rel":
				ch.Rel = strings.TrimSpace(value)
			case "header":
				ch.Header = strings.TrimSpace(value)
			}
		}
	}

	return ch, nil
}

func LintCommitMsg(commitMsg string) error {
	commitHeader := strings.Split(commitMsg, "\n")[0]

	ch, err := parseCommitHeader(commitHeader)
	if err != nil {
		return err
	}

	if err := validateType(ch.Type); err != nil {
		return err
	}

	if ch.Rel != "" {
		if err := validateRel(ch.Rel); err != nil {
			return err
		}
	}

	return nil
}
