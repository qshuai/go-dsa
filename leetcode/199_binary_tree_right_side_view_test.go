package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name string
		args *types.TreeNode
		want []int
	}{
		{
			name: "case-1",
			args: types.NewBinaryTree([]any{1, 2, 3, nil, 5, nil, 4}),
			want: []int{1, 3, 4},
		},
		{
			name: "case-2",
			args: types.NewBinaryTree([]any{1, nil, 3}),
			want: []int{1, 3},
		},
		{
			name: "case-3",
			args: types.NewBinaryTree([]any{1, 2}),
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
