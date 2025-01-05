package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_deleteDuplicates(t *testing.T) {
	type args[T comparable] struct {
		head *types.ListNode[T]
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want *types.ListNode[T]
	}

	tests := []testCase[int]{
		{
			name: "empty list",
			args: args[int]{
				head: nil,
			},
			want: nil,
		},
		{
			name: "single node",
			args: args[int]{
				head: &types.ListNode[int]{Value: 1},
			},
			want: &types.ListNode[int]{Value: 1},
		},
		{
			name: "duplicates1",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{1, 1, 2}),
			},
			want: types.NewListNodeFromSlice([]int{1, 2}),
		},
		{
			name: "duplicates2",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{1, 1, 2, 3, 3}),
			},
			want: types.NewListNodeFromSlice([]int{1, 2, 3}),
		},
		{
			name: "duplicates3",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{0, 0, 0, 0, 0}),
			},
			want: types.NewListNodeWithValue(0),
		},
		{
			name: "duplicates4",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{1, 2, 2}),
			},
			want: types.NewListNodeFromSlice([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
