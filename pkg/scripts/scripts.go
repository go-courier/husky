package scripts

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/go-courier/husky/pkg/fmtx"
)

func RunScripts(scripts []string) error {
	for _, s := range scripts {
		cmd := s

		fmtx.Fprintln(os.Stdout, color.MagentaString(cmd))

		if err := StdRun(cmd); err != nil {
			return err
		}
	}
	return nil
}

func StdRun(cmd string) error {
	sh := "sh"
	if runtime.GOOS == "windows" {
		sh = "bash"
	}
	return stdRun(exec.Command(sh, "-c", cmd))
}

func stdRun(cmd *exec.Cmd) error {
	{
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		go scanAndStdout(bufio.NewScanner(stdoutPipe))
	}

	{
		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		go scanAndStderr(bufio.NewScanner(stderrPipe))
	}

	return cmd.Run()
}

func scanAndStdout(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmtx.Fprintln(os.Stderr, scanner.Text())
	}
}

func scanAndStderr(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmtx.Fprintln(os.Stderr, color.RedString("%s", scanner.Text()))
	}
}
