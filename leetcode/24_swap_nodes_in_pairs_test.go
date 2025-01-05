package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode[int]
		want *types.ListNode[int]
	}{
		{
			name: "linked list with even number elements",
			args: types.NewListNodeSequence(1, 4),
			want: types.NewListNodeFromSlice([]int{2, 1, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
