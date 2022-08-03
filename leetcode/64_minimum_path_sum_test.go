package leetcode

import (
	"testing"
)

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{
			name: "case-1",
			args: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1}},
			want: 7, // 1 -> 3 -> 1 -> 1 -> 1
		},
		{
			name: "case-2",
			args: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12, // 1 -> 2 -> 3 -> 6
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
