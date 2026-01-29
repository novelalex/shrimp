package builtins

import (
	"fmt"

	"github.com/novelalex/shrimp/internal/context"
)

func Type(shell context.Context, args []string) (output string, err error) {
	usageString := "usage: type <command>"
	expectedArgCount := 1

	if len(args) != expectedArgCount {
		return usageString, fmt.Errorf("Incorrect number of args provided: expected %d, got %d", expectedArgCount, len(args))
	}

	cmd := args[0]

	if _, ok := BuiltInCmdMap[cmd]; ok {
		output = fmt.Sprintf("%s is a shell builtin", cmd)
	} else {
		output = fmt.Sprintf("%s: not found", cmd)
	}

	return output, nil
}
