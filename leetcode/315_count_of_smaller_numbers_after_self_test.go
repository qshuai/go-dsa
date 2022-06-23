package leetcode

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "case-1",
			args: []int{5, 2, 6, 1},
			want: []int{2, 1, 1, 0},
		},
		{
			name: "case-2",
			args: []int{-1},
			want: []int{0},
		},
		{
			name: "case-3",
			args: []int{-1, -1},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
