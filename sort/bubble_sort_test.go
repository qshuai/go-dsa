package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_BubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		target   []int
		expected []int
	}{
		{
			name:     "case1(has sorted array)",
			target:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case2(reversed array)",
			target:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case3(truffle array)",
			target:   []int{2, 3, 1, 345, 2, 56, 86, 34, 0, -23},
			expected: []int{-23, 0, 1, 2, 2, 3, 34, 56, 86, 345},
		},
	}

	for _, test := range tests {
		sorted := bubbleSort(test.target)
		if !reflect.DeepEqual(sorted, test.expected) {
			t.Errorf("test: %s failed, expected: %v, but got: %v",
				test.name, test.expected, sorted)
		}
	}
}

func Test_ListBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode
		want *types.ListNode
	}{
		{
			name: "case-1",
			args: types.NewListNodeFromSlice([]int{5, 4, 3, 2, 1}),
			want: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case-2",
			args: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
			want: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := listBubbleSort(test.args); !reflect.DeepEqual(test.want, got) {
				t.Errorf("listBubbleSort() %s failed, expected: %s, but got %s",
					test.name, test.want, got)
			}
		})
	}
}

func Test_ListBubbleSortChangValue(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode
		want *types.ListNode
	}{
		{
			name: "case-1",
			args: types.NewListNodeFromSlice([]int{5, 4, 3, 2, 1}),
			want: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case-2",
			args: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
			want: types.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if listBubbleSortChangValue(test.args); !reflect.DeepEqual(test.want, test.args) {
				t.Errorf("listBubbleSortChangValue() %s failed, expected: %s, but got %s",
					test.name, test.want, test.args)
			}
		})
	}
}

func Test_DoublyListBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args *types.DListNode
		want *types.DListNode
	}{
		{
			name: "case-1",
			args: types.NewDoublyLinkedListFromSlice([]int{5, 4, 3, 2, 1}),
			want: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case-2",
			args: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			want: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := doublyListBubbleSort(test.args); !reflect.DeepEqual(test.want, got) {
				t.Errorf("DoublyListBubbleSort() %s failed, expected: %s, but got %s",
					test.name, test.want, got)
			}
		})
	}
}

func Test_DoublyListBubbleSortChangeValue(t *testing.T) {
	tests := []struct {
		name string
		args *types.DListNode
		want *types.DListNode
	}{
		{
			name: "case-1",
			args: types.NewDoublyLinkedListFromSlice([]int{5, 4, 3, 2, 1}),
			want: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "case-2",
			args: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			want: types.NewDoublyLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if doublyListBubbleSortChangeValue(test.args); !reflect.DeepEqual(test.want, test.args) {
				t.Errorf("DoublyListBubbleSortChangeValue() %s failed, expected: %s, but got %s",
					test.name, test.want, test.args)
			}
		})
	}
}
