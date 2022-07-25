package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode
		want *types.ListNode
	}{
		{
			name: "case-1",
			args: types.NewListNodeFromSlice([]int{1, 2, 3, 3, 4, 4, 5}),
			want: types.NewListNodeFromSlice([]int{1, 2, 5}),
		},
		{
			name: "case-2",
			args: types.NewListNodeFromSlice([]int{1, 1, 1, 2, 3}),
			want: types.NewListNodeFromSlice([]int{2, 3}),
		},
		{
			name: "case-3",
			args: types.NewListNodeFromSlice([]int{1, 2, 2, 2}),
			want: types.NewListNodeFromSlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
