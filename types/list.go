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

// NewListNodeFromSlice construct linked list from slice
func NewListNodeFromSlice(s []int) *ListNode {
	if len(s) <= 0 {
		return nil
	}

	head := &ListNode{Value: s[0]}
	cursor := head
	for i := 1; i < len(s); i++ {
		newNode := &ListNode{Value: s[i]}
		cursor.Next = newNode
		cursor = newNode
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

// RemoveByPosition remove the Nth using 0-based indexing node
// from begin of linked list
func (list *ListNode) RemoveByPosition(idx int) *ListNode {
	if list == nil {
		panic("nil linked list can not remove element")
	}
	if idx < 0 {
		panic("idx should be larger or equal than zero")
	}

	// 删除第一个元素
	if idx == 0 {
		return list.Next
	}

	head := list
	i := 1
	for list != nil {
		if i == idx {
			if list.Next == nil {
				panic("idx out of bounds")
			}

			list.Next = list.Next.Next
			return head
		}

		i++
		list = list.Next
	}

	return head
}

// ReverseSingleList 单链表反转
func ReverseSingleList(list *ListNode) *ListNode {
	cur := list
	var prev *ListNode
	var tmp *ListNode

	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
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

// NewDoublyLinkedListFromSlice 从数组生成双向链表
func NewDoublyLinkedListFromSlice(arr []int) *DListNode {
	if len(arr) <= 0 {
		return nil
	}

	head := &DListNode{
		Value: arr[0],
		Prev:  nil,
		Next:  nil,
	}
	prev := head
	for i := 1; i < len(arr); i++ {
		newNode := &DListNode{
			Value: arr[i],
			Prev:  prev,
			Next:  nil,
		}
		prev.Next = newNode

		prev = newNode
	}

	return head
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
