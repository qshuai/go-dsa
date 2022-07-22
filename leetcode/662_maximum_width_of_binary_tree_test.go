package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_widthOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		args *types.Node
		want int
	}{
		{
			name: "case-1",
			args: &types.Node{
				Value: 1,
				Left: &types.Node{
					Value: 3,
					Left: &types.Node{
						Value: 5,
						Left: &types.Node{
							Value: 6,
						},
					},
				},
				Right: &types.Node{
					Value: 2,
					Right: &types.Node{
						Value: 9,
						Left: &types.Node{
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
