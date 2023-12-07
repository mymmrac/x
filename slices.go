package x

import (
	"cmp"
)

func Convert[S1 ~[]E1, S2 ~[]E2, E1, E2 any](x S1, conv func(a E1) E2) S2 {
	y := make(S2, len(x))
	for i := 0; i < len(x); i++ {
		y[i] = conv(x[i])
	}
	return y
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var s E
	for i := 0; i < len(x); i++ {
		s += x[i]
	}
	return s
}
