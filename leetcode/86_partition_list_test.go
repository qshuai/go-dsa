package leetcode

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

func Test_partitionList(t *testing.T) {
	type args[T constraints.Ordered] struct {
		head *types.ListNode[T]
		x    T
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
				head: types.NewListNodeFromSlice([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: types.NewListNodeFromSlice([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			name: "case-2",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{2, 1}),
				x:    2,
			},
			want: types.NewListNodeFromSlice([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionList(tt.args.head.Clone(), tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionList() = %v, want %v", got, tt.want)
			}

			if got := partitionList2(tt.args.head.Clone(), tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
