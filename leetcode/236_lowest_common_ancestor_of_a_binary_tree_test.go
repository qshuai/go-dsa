package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *types.TreeNode
		p    *types.TreeNode
		q    *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "case-1",
			args: args{
				root: types.NewBinaryTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
				p:    &types.TreeNode{Value: 5},
				q:    &types.TreeNode{Value: 1},
			},
			want: 3,
		},
		{
			name: "case-2",
			args: args{
				root: types.NewBinaryTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
				p:    &types.TreeNode{Value: 5},
				q:    &types.TreeNode{Value: 4},
			},
			want: 5,
		},
		{
			name: "case-3",
			args: args{
				root: types.NewBinaryTree([]any{1, 2}),
				p:    &types.TreeNode{Value: 1},
				q:    &types.TreeNode{Value: 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Value, tt.want)
			}
		})
	}
}
