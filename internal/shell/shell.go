package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/novelalex/shrimp/internal/builtins"
)

type Shell struct {
	IsRunning bool
}

func (s *Shell) SetIsRunning(b bool) {
	s.IsRunning = b
}

func NewShell() *Shell {
	return &Shell{
		IsRunning: true,
	}
}

func (shell *Shell) REPL() {
	reader := bufio.NewReader(os.Stdin)
	for shell.IsRunning {
		fmt.Fprint(os.Stdout, "$ ")

		rawCmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
		}

		// remove newline (and carriage return on windows i guess)
		cmdString := strings.TrimSpace(rawCmd)

		cmd, args := splitCommandAndArgs(cmdString)

		if builtinFunction, ok := builtins.BuiltIns[cmd]; ok {
			out, err := builtinFunction(shell, args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error running %s command: %s\n", cmd, err)
			}
			fmt.Fprintf(os.Stdout, "%s\n", out)

		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}

func splitCommandAndArgs(cmdString string) (cmd string, args []string) {
	cmdList := strings.Fields(cmdString)
	cmd = cmdList[0]
	args = cmdList[1:]
	return cmd, args
}
