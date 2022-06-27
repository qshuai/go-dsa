package utils

import "golang.org/x/exp/constraints"

// Max returns the maximum item of the two inputs
func Max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	}

	return y
}

// Min returns the minimum item of the two inputs
func Min[T constraints.Ordered](x, y T) T {
	if x <= y {
		return x
	}

	return y
}
