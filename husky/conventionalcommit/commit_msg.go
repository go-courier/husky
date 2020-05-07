package conventionalcommit

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// https://www.conventionalcommits.org/en/v1.0.0/

var reHeader = regexp.MustCompile("^(?P<type>\\w+)(\\((?P<scope>[\\w/]+)\\))?(?P<breaking>!)?:( +)?(?P<header>.+)")

var Types = map[string]string{
	"feat":     "Features",
	"fix":      "Bug Fixes",
	"build":    "",
	"chore":    "",
	"ci":       "",
	"docs":     "",
	"perf":     "",
	"refactor": "",
	"revert":   "",
	"style":    "",
	"test":     "",
}

func validateType(tpe string) error {
	if _, ok := Types[tpe]; ok {
		return nil
	}
	return fmt.Errorf("invalid type `%s`, see more https://www.conventionalcommits.org/en/v1.0.0", tpe)
}

func ParseCommitMsg(commitMsg string) (*CommitMsg, error) {
	commitLines := strings.Split(commitMsg, "\n\n")

	commitHeader := commitLines[0]

	if !reHeader.MatchString(commitHeader) {
		return nil, fmt.Errorf("invalid header format `%s`, should be %s", commitHeader, reHeader.String())
	}

	groupNames := reHeader.SubexpNames()

	c := &CommitMsg{}

	for _, match := range reHeader.FindAllStringSubmatch(commitHeader, -1) {
		for groupIdx, value := range match {
			name := groupNames[groupIdx]
			switch name {
			case "breaking":
				c.BreakingChange = value == "!"
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
		c.Body = strings.Join(commitLines[1:], "\n\n")
	}

	if strings.HasPrefix(c.Body, "BREAKING CHANGE:") {
		c.BreakingChange = true
	}

	return c, nil
}

type CommitMsg struct {
	Type           string
	BreakingChange bool
	Scope          string
	Header         string
	Body           string
}

func (v CommitMsg) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(v.Type)

	if v.Scope != "" {
		buf.WriteString("(")
		buf.WriteString(v.Scope)
		buf.WriteString(")")
	}

	if v.BreakingChange {
		buf.WriteString("!")
	}

	if v.Header != "" {
		buf.WriteString(" ")
		buf.WriteString(v.Header)
	}

	if v.Body != "" {
		buf.WriteString("\n\n")
		buf.WriteString(v.Body)
	}

	return buf.String()
}

func (v CommitMsg) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v *CommitMsg) UnmarshalText(data []byte) error {
	commitMsg, err := ParseCommitMsg(string(data))
	if err != nil {
		return err
	}
	*v = *commitMsg
	return nil
}
