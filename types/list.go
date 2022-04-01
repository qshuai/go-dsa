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
func NewListNode(start, length int) *ListNode {
	head := &ListNode{
		Value: start,
		Next:  nil,
	}

	prev := head
	for i := 1; i < length; i++ {
		start++
		newNode := &ListNode{
			Value: start,
			Next:  nil,
		}
		prev.Next = newNode
		prev = newNode
	}

	return head
}

// Append append the node to the tail. panic if the linked list is nil.
func (list *ListNode) Append(node *ListNode) *ListNode {
	if list == nil {
		panic("nil linked list can not append node")
	}

	head := list
	for list.Next != nil {
		list = list.Next
	}

	list.Next = node
	return head
}

func (list *ListNode) String() string {
	sb := strings.Builder{}
	sb.WriteString("singly linked list: [")
	for list != nil {
		if v, ok := list.Value.(fmt.Stringer); ok {
			sb.WriteString(fmt.Sprintf("%s", v))
		} else {
			sb.WriteString(fmt.Sprintf("%v", list.Value))
		}

		list = list.Next
		if list != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")

	return sb.String()
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
