package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want [][]int
	}{
		{
			name: "case-1",
			args: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "case-2",
			args: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateMatrix(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
