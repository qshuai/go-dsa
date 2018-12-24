package sort

import (
	"testing"
)

func TestSimpleBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		target   []int
		expected []int
	}{
		{
			name:     "case1(has sorted array)",
			target:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case2(reversed array)",
			target:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case3(truffle array)",
			target:   []int{2, 3, 1, 345, 2, 56, 86, 34, 0, -23},
			expected: []int{-23, 0, 1, 2, 2, 3, 34, 56, 86, 345},
		},
	}

	for _, test := range tests {
		sorted := simpleBubbleSort(test.target)
		if !equal(sorted, test.expected) {
			t.Errorf("test: %s failed, expected: %v, but got: %v",
				test.name, test.expected, sorted)
		}
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
