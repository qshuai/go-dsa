package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode
		want bool
	}{
		{
			name: "cycle linked list",
			args: buildCycleLinkedList(4, 1),
			want: true,
		},
		{
			name: "head-tail cycle linked list",
			args: buildCycleLinkedList(2, 0),
			want: true,
		},
		{
			name: "single element linked list",
			args: types.NewListNode(1, 1),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// buildCycleLinkedList 创建一个有环的单向链表
// params:
// 1. length: 链表元素的个数
// 2. pos: 尾节点和和链表相交的位置 0-based
func buildCycleLinkedList(length, pos int) *types.ListNode {
	if length <= 0 {
		return nil
	}
	if pos < 0 || pos >= length {
		panic("invalid pos")
	}

	head := types.NewListNode(0, 1)
	prev := head
	target := head
	for i := 1; i < length; i++ {
		newNode := types.NewListNode(i, 1)
		prev.Next = newNode
		prev = newNode

		if i == pos {
			target = newNode
		}
	}

	prev.Next = target
	return head
}
