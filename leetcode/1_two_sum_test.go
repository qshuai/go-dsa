package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		target int
		array  []int
		ret    []int
	}{
		{
			name:   "not found",
			target: 10,
			array:  []int{1, 2, 3, 4, 5},
			ret:    nil,
		},
		{
			name:   "exits",
			target: 8,
			array:  []int{1, 2, 3, 4, 5},
			ret:    []int{2, 4},
		},
		{
			name:   "exits",
			target: 6,
			array:  []int{3, 2, 4},
			ret:    []int{1, 2},
		},
	}

	for _, test := range tests {
		result := twoSum(test.array, test.target)
		if !EqualArray(test.ret, result) {
			t.Errorf("test case(twoSum): %s, want: %v but got: %v", test.name, test.ret, result)
		}
	}

	for _, test := range tests {
		result := twoSumN2(test.array, test.target)
		if !EqualArray(test.ret, result) {
			t.Errorf("test case(twoSumN2): %s, want: %v but got: %v", test.name, test.ret, result)
		}
	}
}

func EqualArray(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
