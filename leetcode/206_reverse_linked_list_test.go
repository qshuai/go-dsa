package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_reverseList(t *testing.T) {
	type args[T any] struct {
		head *types.ListNode[T]
	}
	type testCase[T any] struct {
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
			name: "example1",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
			},
			want: types.NewListNodeFromSlice([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "example2",
			args: args[int]{
				head: types.NewListNodeFromSlice([]int{1, 2}),
			},
			want: types.NewListNodeFromSlice([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
