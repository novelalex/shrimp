package main

import (
	"github.com/novelalex/shrimp/internal/shell"
)

func main() {

	shell := shell.NewShell()

	shell.REPL()
}
