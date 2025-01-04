package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/types"
	"github.com/qshuai/go-dsa/utils"
)

func Test_BubbleSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			if BubbleSort(arr); !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("BubbleSort() = %v, want: %v", arr, tt.Expected)
			}
		})
	}
}

func Test_ListBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode[int]
		want *types.ListNode[int]
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListBubbleSort(tt.args); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ListBubbleSort() = %s, want: %s", got, tt.want)
			}
		})
	}
}

func Test_ListBubbleSortChangValue(t *testing.T) {
	tests := []struct {
		name string
		args *types.ListNode[int]
		want *types.ListNode[int]
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ListBubbleSortChangValue(tt.args); !reflect.DeepEqual(tt.want, tt.args) {
				t.Errorf("ListBubbleSortChangValue() = %s, want: %s", tt.args, tt.want)
			}
		})
	}
}

func Test_DoublyListBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args *types.DListNode[int]
		want *types.DListNode[int]
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoublyListBubbleSort(tt.args); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("DoublyListBubbleSort() = %s, want: %s", got, tt.want)
			}
		})
	}
}

func Test_DoublyListBubbleSortChangeValue(t *testing.T) {
	tests := []struct {
		name string
		args *types.DListNode[int]
		want *types.DListNode[int]
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DoublyListBubbleSortChangeValue(tt.args); !reflect.DeepEqual(tt.want, tt.args) {
				t.Errorf("DoublyListBubbleSortChangeValue() = %s, want: %s", tt.args, tt.want)
			}
		})
	}
}
