package x

import (
	"fmt"
	"os"
	"runtime"
)

func Assert(ok bool, args ...any) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "ASSERT: %s:%d\n", file, line)
		} else {
			fmt.Fprintf(os.Stderr, "ASSERT: %s:%d: %s\n", file, line, fmt.Sprint(args...))
		}
		os.Exit(1)
	}
}
