package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "case-1",
			args: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "name-2",
			args: []int{0},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args); !utils.EmbeddedSliceEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
