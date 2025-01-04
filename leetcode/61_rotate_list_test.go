package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_rotateRight(t *testing.T) {
	type args[T any] struct {
		head *types.ListNode[T]
		k    int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *types.ListNode[T]
	}
	tests := []testCase[any]{
		{
			name: "empty list",
			args: args[any]{
				head: nil,
				k:    3,
			},
			want: nil,
		},
		{
			name: "single node",
			args: args[any]{
				head: &types.ListNode[any]{Value: 1},
				k:    3,
			},
			want: &types.ListNode[any]{Value: 1},
		},
		{
			name: "rotate according the number of node",
			args: args[any]{
				head: types.NewListNodeFromSlice([]any{1, 2, 3, 4, 5}),
				k:    5,
			},
			want: types.NewListNodeFromSlice([]any{1, 2, 3, 4, 5}),
		},
		{
			name: "rotate according the number of node",
			args: args[any]{
				head: types.NewListNodeFromSlice([]any{1, 2, 3, 4, 5}),
				k:    10,
			},
			want: types.NewListNodeFromSlice([]any{1, 2, 3, 4, 5}),
		},
		{
			name: "rotate-1",
			args: args[any]{
				head: types.NewListNodeFromSlice([]any{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: types.NewListNodeFromSlice([]any{3, 4, 5, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
