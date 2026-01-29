package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Fprint(os.Stdout, "$ ")

	raw_command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input: ", err)
	}

	// remove newline (and carriage return on windows i guess)
	command := strings.TrimSpace(raw_command)

	fmt.Fprintf(os.Stdout, "%s: command not found", command)

}
