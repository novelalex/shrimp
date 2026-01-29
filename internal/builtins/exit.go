package builtins

import (
	"github.com/novelalex/shrimp/internal/context"
)

func Exit(shell context.Context, args []string) (output string, err error) {
	shell.SetIsRunning(false)
	return "", nil
}
