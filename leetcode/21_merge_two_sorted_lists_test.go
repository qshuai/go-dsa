package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_mergeSortedList(t *testing.T) {
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
				l1: types.NewListNodeFromSlice([]int{1, 2, 4}),
				l2: types.NewListNodeFromSlice([]int{1, 3, 4}),
			},
			want: types.NewListNodeFromSlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "case-2",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "case-3",
			args: args{
				l1: nil,
				l2: types.NewListNodeWithValue(0),
			},
			want: types.NewListNodeWithValue(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortedList(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
