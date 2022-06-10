package utils

import "golang.org/x/exp/constraints"

// Max returns the max item of the two inputs
func Max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	}

	return y
}
