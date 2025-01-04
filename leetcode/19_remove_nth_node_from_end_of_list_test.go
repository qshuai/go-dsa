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
				head: types.NewListNode(1, 5),
				n:    2,
			},
			want: types.NewListNode(1, 3).Append(&types.ListNode[int]{
				Value: 5,
				Next:  nil,
			}),
		},
		{
			name: "single element linked list",
			args: args{
				head: types.NewListNode(1, 1),
				n:    1,
			},
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}

			// 恢复上面修改的原始参数
			if i == 0 {
				tt.args.head = types.NewListNode(1, 5)
			} else if i == 1 {
				tt.args.head = types.NewListNode(1, 1)
			}
			if got := removeNthFromEnd2(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd2() = %v, want %v", got, tt.want)
			}
		})
	}
}
