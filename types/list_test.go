package types

import (
	"testing"
)

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name string
		arg  *ListNode
		want string
	}{
		{
			name: "general singly linked list",
			arg:  NewListNode(1, 3),
			want: "singly linked list: [1, 2, 3]",
		},
		{
			name: "single element linked list",
			arg:  NewListNode(1, 1),
			want: "singly linked list: [1]",
		},
		{
			name: "empty linked list",
			arg:  nil,
			want: "singly linked list: []",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
