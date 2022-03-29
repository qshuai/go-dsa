package types

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// NewListNode creates a singly linked list
func NewListNode(v interface{}, next *ListNode) *ListNode {
	return &ListNode{
		Value: v,
		Next:  next,
	}
}

// DListNode represents doubly linked list
type DListNode struct {
	Value interface{}
	Prev  *DListNode
	Next  *DListNode
}

// String return readable string representing doubly linked list
func (n *DListNode) String() string {
	if n == nil {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString("nil")
	for n != nil {
		sb.WriteString(fmt.Sprintf(" <=> %v", n.Value))
		n = n.Next
	}
	sb.WriteString(" <=> nil")

	return sb.String()
}
