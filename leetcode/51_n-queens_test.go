package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name string
		args int
		want [][]string
	}{
		{
			name: "case-1",
			args: 4,
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "case-2",
			args: 1,
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args); !utils.EmbeddedSliceEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPos(t *testing.T) {
	type args struct {
		n   int
		row int
		col int
		pos []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{
				n:   4,
				row: 2,
				col: 2,
				pos: []int{0, 1, 4, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPos(tt.args.n, tt.args.row, tt.args.col, tt.args.pos); got != tt.want {
				t.Errorf("checkPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
