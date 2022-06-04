package search

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func TestContainLoopInList1(t *testing.T) {
	tests := []struct {
		name     string
		node     *types.ListNode
		loopFunc func(node *types.ListNode)
		want     bool
	}{
		{
			name:     "only one node",
			node:     types.NewListNode(1, 1),
			loopFunc: nil,
			want:     false,
		},
		{
			name:     "two node list",
			node:     types.NewListNode(1, 2),
			loopFunc: nil,
			want:     false,
		},
		{
			name:     "many node list without loop",
			node:     types.NewListNode(1, 4),
			loopFunc: nil,
			want:     false,
		},
		{
			name: "list with loop",
			node: types.NewListNode(1, 4),
			loopFunc: func(node *types.ListNode) {
				node.Next.Next.Next = node.Next
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.loopFunc != nil {
				tt.loopFunc(tt.node)
			}
			if got := ContainLoopInList(tt.node); got != tt.want {
				t.Errorf("ContainLoopInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
