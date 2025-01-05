package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *types.ListNode[int]
		n    int
	}

	tests := []struct {
		name string
		args args
		want *types.ListNode[int]
	}{
		{
			name: "general case",
			args: args{
				head: types.NewListNodeSequence(1, 5),
				n:    2,
			},
			want: types.NewListNodeSequence(1, 3).Append(&types.ListNode[int]{
				Value: 5,
				Next:  nil,
			}),
		},
		{
			name: "single element linked list",
			args: args{
				head: types.NewListNodeWithValue(1),
				n:    1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head.Clone(), tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}

			if got := removeNthFromEnd2(tt.args.head.Clone(), tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd2() = %v, want %v", got, tt.want)
			}
		})
	}
}
