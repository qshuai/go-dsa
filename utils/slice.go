package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Contain returns true if target is the element of container.
func Contain[T comparable](container []T, target T) bool {
	for i := 0; i < len(container); i++ {
		if container[i] == target {
			return true
		}
	}

	return false
}

// Copy returns the copy of given slice. Maybe shallow copy
// depending on the type T, eg: pointer or map
func Copy[T any](src []T) []T {
	if src == nil {
		return nil
	}
	if len(src) <= 0 {
		return []T{}
	}

	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

// ElementEqual 查看两个slice是否元素相同，不区分元素顺序
func ElementEqual[T constraints.Ordered](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// EmbeddedSliceEqual 判断内嵌的slice是否相同，不区分元素顺序
func EmbeddedSliceEqual[T constraints.Ordered](a [][]T, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) <= 0 {
		return true
	}

	passedMapping := make(map[int]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if _, ok := passedMapping[j]; ok {
				continue
			}

			if ElementEqual(a[i], b[j]) {
				passedMapping[j] = struct{}{}
			}
		}
	}

	return len(passedMapping) == len(b)
}

func ContainEqual[T constraints.Ordered](arrs [][]T, target []T) bool {
	for _, arr := range arrs {
		equal := ElementEqual(arr, target)
		if equal {
			return true
		}
	}

	return false
}
