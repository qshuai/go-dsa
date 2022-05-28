package leetcode

import (
	"testing"
)

func TestSearchInRotatedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "case-1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "case-2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "case-3",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "case-4",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "case-5",
			nums:     []int{5, 1, 3},
			target:   3,
			expected: 2,
		},
		{
			name:     "case-6",
			nums:     []int{8, 9, 2, 3, 4},
			target:   9,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedArray(tt.nums, tt.target); got != tt.expected {
				t.Errorf("SearchInRotatedArray[%s] = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
