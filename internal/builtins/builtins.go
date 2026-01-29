package builtins

import (
	"github.com/novelalex/shrimp/internal/context"
)

type BuiltInFunction func(shell context.Context, args []string) (output string, err error)

var BuiltIns = map[string]BuiltInFunction{
	"exit": Exit,
	"echo": Echo,
	"type": Type,
}

var BuiltInCmdMap = map[string]bool{}

func init() {
	for key := range BuiltIns {
		BuiltInCmdMap[key] = true
	}
}
