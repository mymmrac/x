package x

import (
	"cmp"
)

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var s E
	for i := 0; i < len(x); i++ {
		s += x[i]
	}
	return s
}
