package conventionalcommit

import (
	"fmt"
	"regexp"
	"strings"
)

// https://www.conventionalcommits.org/en/v1.0.0-beta.2/

var reHeader = regexp.MustCompile("^(?P<type>\\w+)(\\((?P<scope>[\\w/]+)\\))?:( +)?(?P<header>.+)")

var breakingChange = "BREAKING CHANGE"

var types = []string{
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

type Commit struct {
	Type   string
	Scope  string
	Header string
	Body   string
	Footer string
}

func validateType(tpe string) error {
	for _, t := range types {
		if tpe == t {
			return nil
		}
	}

	return fmt.Errorf("invalid type `%s`, type should be one of %v", tpe, types)
}

func ParseCommit(commitMsg string) (*Commit, error) {
	commitLines := strings.Split(commitMsg, "\n\n")

	commitHeader := commitLines[0]

	if !reHeader.MatchString(commitHeader) {
		return nil, fmt.Errorf("invalid header format `%s`, should be %s", commitHeader, reHeader.String())
	}

	groupNames := reHeader.SubexpNames()

	c := &Commit{}

	for _, match := range reHeader.FindAllStringSubmatch(commitHeader, -1) {
		for groupIdx, value := range match {
			name := groupNames[groupIdx]
			switch name {
			case "type":
				c.Type = strings.TrimSpace(value)
				if err := validateType(c.Type); err != nil {
					return nil, err
				}
			case "scope":
				c.Scope = strings.TrimSpace(value)
			case "header":
				c.Header = strings.TrimSpace(value)
			}
		}
	}

	if len(commitLines) > 1 {
		if len(commitLines) > 2 {
			c.Footer = commitLines[len(commitLines)-1]
			c.Body = strings.Join(commitLines[1:len(commitLines)-1], "\n\n")
		} else {
			c.Body = strings.Join(commitLines[1:], "\n\n")
		}
	}

	return c, nil
}
