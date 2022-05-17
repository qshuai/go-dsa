package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		args     []int
		copy     []int
		expected []int
	}{
		{
			name:     "case1(has sorted array)",
			args:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			copy:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case2(reversed array)",
			args:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			copy:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case3(truffle array)",
			args:     []int{2, 3, 1, 345, 2, 56, 86, 34, 0, -23},
			copy:     []int{2, 3, 1, 345, 2, 56, 86, 34, 0, -23},
			expected: []int{-23, 0, 1, 2, 2, 3, 34, 56, 86, 345},
		},
		{
			name:     "case4(one step sort)",
			args:     []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
			copy:     []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case5(normal case)",
			args:     []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			copy:     []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			expected: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args, PartitionHead)
			if !reflect.DeepEqual(tt.args, tt.expected) {
				t.Errorf("%s[PartitionHead] expected: %v but got: %v", tt.name, tt.expected, tt.args)
			}
			copy(tt.args, tt.copy) // args recovery

			QuickSort(tt.args, PartitionMiddle)
			if !reflect.DeepEqual(tt.args, tt.expected) {
				t.Errorf("%s[PartitionMiddle] expected: %v but got: %v", tt.name, tt.expected, tt.args)
			}
			copy(tt.args, tt.copy) // args recovery

			QuickSort(tt.args, PartitionTail)
			if !reflect.DeepEqual(tt.args, tt.expected) {
				t.Errorf("%s[PartitionTail] expected: %v but got: %v", tt.name, tt.expected, tt.args)
			}
			copy(tt.args, tt.copy) // args recovery

			QuickSort(tt.args, PartitionMedian)
			if !reflect.DeepEqual(tt.args, tt.expected) {
				t.Errorf("%s[PartitionMedian] expected: %v but got: %v", tt.name, tt.expected, tt.args)
			}
		})
	}
}
