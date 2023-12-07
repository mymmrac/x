package x

import (
	"cmp"
)

func Convert[T1, T2 any](x []T1, conv func(item T1) T2) []T2 {
	y := make([]T2, len(x))
	for i := 0; i < len(x); i++ {
		y[i] = conv(x[i])
	}
	return y
}

func Sum[T cmp.Ordered](x []T) T {
	var s T
	for i := 0; i < len(x); i++ {
		s += x[i]
	}
	return s
}
