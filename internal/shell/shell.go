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

		raw_command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
		}

		// remove newline (and carriage return on windows i guess)
		command_string := strings.TrimSpace(raw_command)
		command_list := strings.Fields(command_string)
		command := command_list[0]
		args := command_list[1:]

		if command == "exit" {
			builtins.Exit(shell, args)
		} else if command == "echo" {
			builtins.Echo(shell, args)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}
