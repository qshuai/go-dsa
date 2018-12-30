package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		index  int
	}{
		{
			name:   "exist",
			array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 6,
			index:  5,
		},
		{
			name:   "not exist",
			array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 11,
			index:  -1,
		},
		{
			name:   "exist(end subtracts start is 1 at last)",
			array:  []int{1, 2, 3, 4, 5},
			target: 2,
			index:  1,
		},
		{
			name:   "not exist(end subtracts start is 1 at last)",
			array:  []int{1, 2, 4, 5, 6},
			target: 3,
			index:  -1,
		},
	}

	for _, test := range tests {
		index := binarySearchNoRecursion(test.array, test.target)
		if index != test.index {
			t.Errorf("target number: %d index, want: %d, but got: %d",
				test.target, test.index, index)
		}

		index2 := binarySearchWithRecursion(test.array, 0, len(test.array)-1, test.target)
		if index2 != test.index {
			t.Errorf("target number: %d index, want: %d, but got: %d",
				test.target, test.index, index2)
		}
	}
}
