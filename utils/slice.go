package utils

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

// InInts 查询元素是否在切片中
func InInts(nums []int, target int) bool {
	if len(nums) <= 0 {
		return false
	}

	for _, num := range nums {
		if target == num {
			return true
		}
	}

	return false
}

// ElementEqual 查看两个slice是否元素相同，不区分元素顺序
func ElementEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) <= 0 {
		return true
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == a[j] {
				break
			}
		}
	}

	return true
}

// EmbeddedSliceEqual 判断内嵌的slice是否相同，不区分元素顺序
func EmbeddedSliceEqual[T comparable](a [][]T, b [][]T) bool {
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
