package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		args *types.TreeNode
		want [][]int
	}{
		{
			name: "case-1",
			args: &types.TreeNode{
				Val: 3,
				Left: &types.TreeNode{
					Val: 9,
				},
				Right: &types.TreeNode{
					Val: 20,
					Left: &types.TreeNode{
						Val: 15,
					},
					Right: &types.TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
