package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name string
		arg  *ListNode[int]
		want string
	}{
		{
			name: "general singly linked list",
			arg:  NewListNodeSequence(1, 3),
			want: "1 -> 2 -> 3",
		},
		{
			name: "single element linked list",
			arg:  NewListNodeWithValue(1),
			want: "1",
		},
		{
			name: "empty linked list",
			arg:  nil,
			want: "nil",
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
		head *ListNode[int]
		idx  int
	}
	tests := []struct {
		name string
		args arg
		want *ListNode[int]
	}{
		{
			name: "general case",
			args: arg{
				head: NewListNodeSequence(1, 5),
				idx:  2,
			},
			want: NewListNodeSequence(1, 2).Append(NewListNodeSequence(4, 2)),
		},
		{
			name: "remove the first node of linked list having one element",
			args: arg{
				head: &ListNode[int]{1, nil},
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

func TestNewListNodeFromSlice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want *ListNode[int]
	}{
		{
			name: "general case",
			args: []int{1, 2, 3},
			want: NewListNodeSequence(1, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNodeFromSlice(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNodeFromSlice() = %v want %v", got, tt.want)
			}
		})
	}
}

func ExampleListNode_Reverse() {
	l := NewListNodeSequence(1, 5)
	fmt.Println(l.Reverse())

	// Output:
	// 5 -> 4 -> 3 -> 2 -> 1
}

func ExampleListNode_String() {
	fmt.Println((*ListNode[any])(nil))
	fmt.Println(&ListNode[int]{Value: 4})
	fmt.Println(NewListNodeFromSlice([]int{5, 2, 3, 5, 9, 1}))

	// Output:
	// nil
	// 4
	// 5 -> 2 -> 3 -> 5 -> 9 -> 1
}

func TestNewDoublyLinkedListFromSlice(t *testing.T) {
	head := &DListNode[int]{
		Value: 1,
		Prev:  nil,
		Next:  nil,
	}
	node2 := &DListNode[int]{
		Value: 2,
		Prev:  head,
		Next:  nil,
	}
	head.Next = node2
	node3 := &DListNode[int]{
		Value: 3,
		Prev:  node2,
		Next:  nil,
	}
	node2.Next = node3

	tests := []struct {
		name string
		args []int
		want *DListNode[int]
	}{
		{
			name: "general case",
			args: []int{1, 2, 3},
			want: head,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := NewDoublyLinkedListFromSlice(test.args); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewDoublyLinkedListFromSlice() %s want: %s, but got: %s",
					test.name, test.want.String(), got.String())
			}
		})
	}
}

func ExampleDListNode_String() {
	fmt.Println((*DListNode[any])(nil))

	dl := NewDoublyLinkedListFromSlice([]string{"hello", "goland", "awesome"})
	fmt.Println(dl.String())

	// Output:
	// nil
	// hello <-> goland <-> awesome
}

func TestListNode_Clone(t *testing.T) {
	type testCase[T any] struct {
		name string
		n    *ListNode[T]
		want *ListNode[T]
	}
	tests := []testCase[int]{
		{
			name: "nil list",
			n:    nil,
			want: nil,
		},
		{
			name: "single node list",
			n:    NewListNodeWithValue(3),
			want: NewListNodeWithValue(3),
		},
		{
			name: "multi node list",
			n:    NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
			want: NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
