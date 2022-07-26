package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		name string
		args *types.TreeNode
		want int
	}{
		{
			name: "case-1",
			args: types.NewBinaryTree([]any{1, 2, 3}),
			want: 25,
		},
		{
			name: "case-2",
			args: types.NewBinaryTree([]any{4, 9, 0, 5, 1}),
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
