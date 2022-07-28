package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		args *types.TreeNode
		want bool
	}{
		{
			name: "case-1",
			args: types.NewBinaryTree([]any{1, 2, 2, 3, 4, 4, 3}),
			want: true,
		},
		{
			name: "case-2",
			args: types.NewBinaryTree([]any{1, 2, 2, nil, 3, nil, 3}),
			want: false,
		},
		{
			name: "case-3",
			args: types.NewBinaryTree([]any{1, 2, 2, nil, 3, 3, nil}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
