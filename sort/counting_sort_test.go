package sort

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		nums []int
		min  int
		max  int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "case-1",
			args: args{
				nums: []int{-1, 1, 2, 1, 3},
				min:  -1,
				max:  3,
			},
			expected: []int{-1, 1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountingSort(tt.args.nums, tt.args.min, tt.args.max)
			if !reflect.DeepEqual(tt.expected, tt.args.nums) {
				t.Errorf("%s expected: %v, bug got: %v", tt.name, tt.expected, tt.args.nums)
			}
		})
	}
}
