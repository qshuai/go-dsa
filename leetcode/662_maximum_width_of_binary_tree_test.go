package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_widthOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		args *types.TreeNode
		want int
	}{
		{
			name: "case-1",
			args: &types.TreeNode{
				Value: 1,
				Left: &types.TreeNode{
					Value: 3,
					Left: &types.TreeNode{
						Value: 5,
						Left: &types.TreeNode{
							Value: 6,
						},
					},
				},
				Right: &types.TreeNode{
					Value: 2,
					Right: &types.TreeNode{
						Value: 9,
						Left: &types.TreeNode{
							Value: 7,
						},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
