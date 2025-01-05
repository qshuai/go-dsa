package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_deleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode[int]
		want *types.ListNode[int]
	}{
		{
			name: "general case",
			args: types.NewListNodeSequence(1, 7),
			want: types.NewListNodeSequence(1, 7).RemoveByPosition(3),
		},
		{
			name: "two elements linked list",
			args: types.NewListNodeSequence(1, 2),
			want: types.NewListNodeWithValue(1),
		},
		{
			name: "single element linked list",
			args: types.NewListNodeWithValue(1),
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
