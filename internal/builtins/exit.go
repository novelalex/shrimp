package builtins

import (
	"github.com/novelalex/shrimp/internal/api"
)

func Exit(shell api.Context, args []string) {
	shell.SetIsRunning(false)
}
