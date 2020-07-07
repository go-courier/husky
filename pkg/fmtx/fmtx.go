package fmtx

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

func Fprintln(w io.Writer, a ...interface{}) {
	TopicFprintln("husky", w, a...)
}

func TopicFprintln(topic string, w io.Writer, a ...interface{}) {
	fmt.Fprintf(w, "%s%s\n", color.BlueString("%s > ", topic), fmt.Sprint(a...))
}
