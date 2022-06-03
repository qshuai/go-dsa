package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "case-1",
			args: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "case-2",
			args: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "case-3",
			args: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nextPermutation(tt.args); !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
