package search

import "testing"

func TestFindDuplicationElement(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
	}{
		{
			name:   "sorted args without duplicated element",
			array:  []int{0, 1, 2, 3, 4, 5},
			target: -1,
		},
		{
			name:   "shuffle args without duplicated element",
			array:  []int{3, 5, 0, 2, 1, 4},
			target: -1,
		},
		{
			name:   "duplicated element at expect 0",
			array:  []int{2, 0, 5, 2, 1, 3},
			target: 2,
		},
		{
			name:   "duplicated element at arbitrary expect",
			array:  []int{3, 4, 0, 1, 5, 4},
			target: 4,
		},
		{
			name:   "two duplicated elements",
			array:  []int{3, 4, 0, 5, 5, 4},
			target: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicationElement(tt.array); got != tt.target {
				t.Errorf("FindDuplicationElement() = %d, want: %d", got, tt.target)
			}
		})

	}
}
