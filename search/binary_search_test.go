package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{
			name:   "exist",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 6,
			expect: 5,
		},
		{
			name:   "not exist",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 11,
			expect: -1,
		},
		{
			name:   "exist(end subtracts start is 1 at last)",
			nums:   []int{1, 2, 3, 4, 5},
			target: 2,
			expect: 1,
		},
		{
			name:   "not exist(end subtracts start is 1 at last)",
			nums:   []int{1, 2, 4, 5, 6},
			target: 3,
			expect: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAny(tt.nums, tt.target); got != tt.expect {
				t.Errorf("FindAny() = %d, want: %d", got, tt.expect)
			}

			if got := FindAnyRecursive(tt.nums, tt.target); got != tt.expect {
				t.Errorf("FindAnyRecursive() = %d, want: %d", got, tt.expect)
			}
		})
	}
}

func TestFirstIndex(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		target        int
		firstIndex    int
		lastIndex     int
		firstGteIndex int
		lastLteIndex  int
	}{
		{
			name:          "empty container",
			nums:          nil,
			target:        1,
			firstIndex:    -1,
			lastIndex:     -1,
			firstGteIndex: -1,
			lastLteIndex:  -1,
		},
		{
			name:          "one element container",
			nums:          []int{5},
			target:        5,
			firstIndex:    0,
			lastIndex:     0,
			firstGteIndex: 0,
			lastLteIndex:  0,
		},
		{
			name:          "not contained",
			nums:          []int{1, 2, 3, 4, 5},
			target:        8,
			firstIndex:    -1,
			lastIndex:     -1,
			firstGteIndex: -1,
			lastLteIndex:  4,
		},
		{
			name:          "not duplicative element container",
			nums:          []int{1, 2, 3, 4, 5},
			target:        2,
			firstIndex:    1,
			lastIndex:     1,
			firstGteIndex: 1,
			lastLteIndex:  1,
		},
		{
			name:          "duplicative element container",
			nums:          []int{1, 2, 3, 3, 4, 5},
			target:        3,
			firstIndex:    2,
			lastIndex:     3,
			firstGteIndex: 2,
			lastLteIndex:  3,
		},
		{
			name:          "first gte(or last lte) example",
			nums:          []int{1, 3, 4, 5, 6},
			target:        2,
			firstIndex:    -1,
			lastIndex:     -1,
			firstGteIndex: 1,
			lastLteIndex:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstIndex(tt.nums, tt.target); got != tt.firstIndex {
				t.Errorf("FirstIndex() = %d, want: %d", got, tt.firstIndex)
			}

			if got := LastIndex(tt.nums, tt.target); got != tt.lastIndex {
				t.Errorf("LastIndex() = %d, want: %d", tt.lastIndex, got)
			}

			if got := FirstGteIndex(tt.nums, tt.target); got != tt.firstGteIndex {
				t.Errorf("FirstGteIndex() = %d, want: %d", tt.firstGteIndex, got)
			}

			if got := LastLteIndex(tt.nums, tt.target); got != tt.lastLteIndex {
				t.Errorf("LastLteIndex() = %d, want: %d", tt.lastLteIndex, got)
			}
		})
	}
}
