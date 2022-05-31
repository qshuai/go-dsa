package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/types"
)

func Test_BubbleSort(t *testing.T) {
	for _, test := range testdata.GetTestCases() {
		BubbleSort(test.Args)
		if !reflect.DeepEqual(test.Args, test.Expected) {
			t.Errorf("BubbleSort[%s] expected: %v, but got: %v", test.Name, test.Expected, test.Args)
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
			if got := ListBubbleSort(test.args); !reflect.DeepEqual(test.want, got) {
				t.Errorf("ListBubbleSort() %s failed, expected: %s, but got %s",
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
			if ListBubbleSortChangValue(test.args); !reflect.DeepEqual(test.want, test.args) {
				t.Errorf("ListBubbleSortChangValue() %s failed, expected: %s, but got %s",
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
			if got := DoublyListBubbleSort(test.args); !reflect.DeepEqual(test.want, got) {
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
			if DoublyListBubbleSortChangeValue(test.args); !reflect.DeepEqual(test.want, test.args) {
				t.Errorf("DoublyListBubbleSortChangeValue() %s failed, expected: %s, but got %s",
					test.name, test.want, test.args)
			}
		})
	}
}
