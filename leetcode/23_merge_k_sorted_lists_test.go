package leetcode

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

func Test_mergeKLists(t *testing.T) {
	type args[T constraints.Ordered] struct {
		lists []*types.ListNode[T]
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want *types.ListNode[T]
	}
	tests := []testCase[int]{
		{
			name: "case-1",
			args: args[int]{
				lists: []*types.ListNode[int]{
					types.NewListNodeFromSlice([]int{1, 4, 5}),
					types.NewListNodeFromSlice([]int{1, 3, 4}),
					types.NewListNodeFromSlice([]int{2, 6}),
				},
			},
			want: types.NewListNodeFromSlice([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "case-2",
			args: args[int]{
				lists: []*types.ListNode[int]{},
			},
			want: nil,
		},
		{
			name: "case-3",
			args: args[int]{
				lists: []*types.ListNode[int]{nil},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
