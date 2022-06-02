package leetcode

import (
	"github.com/qshuai/go-dsa/utils"
	"testing"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "case-1",
			args: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args); !utils.EmbeddedSliceEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}

			if got := permute2(tt.args); !utils.EmbeddedSliceEqual(got, tt.want) {
				t.Errorf("permute2() = %v, want %v", got, tt.want)
			}
		})
	}
}
