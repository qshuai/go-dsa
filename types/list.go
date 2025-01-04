package types

import (
	"fmt"
	"strings"
)

type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

// NewListNode creates a singly linked list
func NewListNode[T Number](start T, length int) *ListNode[T] {
	head := &ListNode[T]{
		Value: start,
		Next:  nil,
	}

	prev := head
	for i := 1; i < length; i++ {
		start++
		newNode := &ListNode[T]{
			Value: start,
			Next:  nil,
		}
		prev.Next = newNode
		prev = newNode
	}

	return head
}

// NewListNodeFromSlice construct linked list from slice
func NewListNodeFromSlice[T any](s []T) *ListNode[T] {
	if len(s) <= 0 {
		return nil
	}

	head := &ListNode[T]{Value: s[0]}
	cursor := head
	for i := 1; i < len(s); i++ {
		newNode := &ListNode[T]{Value: s[i]}
		cursor.Next = newNode
		cursor = newNode
	}

	return head
}

// Append the node to the tail. panic if the linked list is nil.
func (n *ListNode[T]) Append(node *ListNode[T]) *ListNode[T] {
	if n == nil {
		panic("append to nil linked list")
	}

	head, cursor := n, n
	for cursor.Next != nil {
		cursor = cursor.Next
	}

	cursor.Next = node
	return head
}

// RemoveByPosition remove the Nth using 0-based indexing node
// from begin of linked list
func (n *ListNode[T]) RemoveByPosition(idx int) *ListNode[T] {
	if n == nil {
		panic("nil linked list can not remove element")
	}
	if idx < 0 {
		panic("idx should be larger or equal than zero")
	}

	// 删除第一个元素
	if idx == 0 {
		return n.Next
	}

	head, cursor := n, n
	i := 1
	for cursor != nil {
		if i == idx {
			if cursor.Next == nil {
				panic("index out of bound")
			}

			cursor.Next = cursor.Next.Next
			return head
		}

		i++
		cursor = cursor.Next
	}

	return head
}

// ReverseSingleList 单链表反转
func ReverseSingleList[T any](list *ListNode[T]) *ListNode[T] {
	cur := list
	var prev *ListNode[T]
	var tmp *ListNode[T]

	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

func (n *ListNode[T]) String() string {
	if n == nil {
		return "nil"
	}

	sb := strings.Builder{}
	cursor := n
	for cursor != nil {
		sb.WriteString(fmt.Sprintf("%v", cursor.Value))

		cursor = cursor.Next
		if cursor != nil {
			sb.WriteString(" -> ")
		}
	}

	return sb.String()
}

// DListNode represents doubly linked list
type DListNode[T any] struct {
	Value any
	Prev  *DListNode[T]
	Next  *DListNode[T]
}

// NewDoublyLinkedListFromSlice 从数组生成双向链表
func NewDoublyLinkedListFromSlice[T any](arr []T) *DListNode[T] {
	if len(arr) <= 0 {
		return nil
	}

	head := &DListNode[T]{
		Value: arr[0],
		Prev:  nil,
		Next:  nil,
	}
	prev := head
	for i := 1; i < len(arr); i++ {
		newNode := &DListNode[T]{
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
func (n *DListNode[T]) String() string {
	if n == nil {
		return "nil"
	}

	sb := strings.Builder{}
	cursor := n
	for cursor != nil {
		if v, ok := cursor.Value.(fmt.Stringer); ok {
			sb.WriteString(v.String())
		} else {
			sb.WriteString(fmt.Sprintf("%v", cursor.Value))
		}

		cursor = cursor.Next
		if cursor != nil {
			sb.WriteString(" <-> ")
		}
	}

	return sb.String()
}
