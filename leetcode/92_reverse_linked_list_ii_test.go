package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *types.ListNode[int]
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode[int]
	}{
		{
			name: "case-1",
			args: args{
				head:  types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: types.NewListNodeFromSlice([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "case-2",
			args: args{
				head:  &types.ListNode[int]{Value: 5},
				left:  1,
				right: 1,
			},
			want: &types.ListNode[int]{Value: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
