package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *types.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "general case",
			args: args{
				head: types.NewListNode(1, 5),
				k:    2,
			},
			want: types.NewListNodeFromSlice([]int{1, 4, 3, 2, 5}),
		},

		{
			name: "swap head and tail element",
			args: args{
				head: types.NewListNode(1, 3),
				k:    1,
			},
			want: types.NewListNodeFromSlice([]int{3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
