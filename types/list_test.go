package types

import (
	"reflect"
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

func TestListNode_RemoveByPosition(t *testing.T) {
	type arg struct {
		head *ListNode
		idx  int
	}
	tests := []struct {
		name string
		args arg
		want *ListNode
	}{
		{
			name: "general case",
			args: arg{
				head: NewListNode(1, 5),
				idx:  2,
			},
			want: NewListNode(1, 2).Append(NewListNode(4, 2)),
		},
		{
			name: "remove the first node of linked list having one element",
			args: arg{
				head: &ListNode{1, nil},
				idx:  0,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.head.RemoveByPosition(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveByPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
