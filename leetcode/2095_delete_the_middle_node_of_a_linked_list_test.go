package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_deleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode
		want *types.ListNode
	}{
		{
			name: "general case",
			args: types.NewListNode(1, 7),
			want: types.NewListNode(1, 7).RemoveByPosition(3),
		},
		{
			name: "two elements linked list",
			args: types.NewListNode(1, 2),
			want: types.NewListNode(1, 1),
		},
		{
			name: "single element linked list",
			args: types.NewListNode(1, 1),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
