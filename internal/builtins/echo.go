package builtins

import (
	"fmt"
	"os"

	"github.com/novelalex/shrimp/internal/api"
)

func Echo(shell api.Context, args []string) {
	for i, arg := range args {
		fmt.Fprint(os.Stdout, arg)

		// put spaces after all args except last
		if i < len(args)-1 {
			fmt.Fprint(os.Stdout, " ")
		}
	}
	fmt.Fprint(os.Stdout, "\n")
}
