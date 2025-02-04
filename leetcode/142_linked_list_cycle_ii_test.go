package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_detectCycle(t *testing.T) {
	type testCase[T any] struct {
		name string
		head *types.ListNode[T]
		want *types.ListNode[T]
	}

	case1 := types.NewListNodeFromSlice([]int{3, 2, 0, -4})
	case1.Tail().Next = case1.Next

	case2 := types.NewListNodeFromSlice([]int{1, 2})
	case2.Next.Next = case2

	tests := []testCase[int]{
		{
			name: "case-1",
			head: case1,
			want: case1.Next,
		},
		{
			name: "case-2",
			head: case2,
			want: case2,
		},
		{
			name: "case-3",
			head: &types.ListNode[int]{
				Value: 1,
				Next:  nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[%s] detectCycle() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
