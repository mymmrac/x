package x

import (
	"fmt"
	"os"
)

func Assert(ok bool, args ...any) {
	if !ok {
		fmt.Fprintf(os.Stderr, "ASSERT: %s\n", fmt.Sprint(args...))
		os.Exit(1)
	}
}
