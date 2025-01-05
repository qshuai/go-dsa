package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *types.ListNode[int]
		l2 *types.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode[int]
	}{
		{
			name: "case-1",
			args: args{
				l1: types.NewListNodeFromSlice([]int{2, 4, 3}),
				l2: types.NewListNodeFromSlice([]int{5, 6, 4}),
			},
			want: types.NewListNodeFromSlice([]int{7, 0, 8}),
		},
		{
			name: "case-2",
			args: args{
				l1: types.NewListNodeWithValue(0),
				l2: types.NewListNodeWithValue(0),
			},
			want: types.NewListNodeWithValue(0),
		},
		{
			name: "case-3",
			args: args{
				l1: types.NewListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: types.NewListNodeFromSlice([]int{9, 9, 9, 9}),
			},
			want: types.NewListNodeFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
