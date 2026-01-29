package builtins

import (
	"bytes"
	"fmt"

	"github.com/novelalex/shrimp/internal/context"
)

func Echo(shell context.Context, args []string) (output string, err error) {
	var buffer bytes.Buffer

	for i, arg := range args {
		fmt.Fprint(&buffer, arg)

		// put spaces after all args except last
		if i < len(args)-1 {
			fmt.Fprint(&buffer, " ")
		}
	}

	return buffer.String(), nil
}
