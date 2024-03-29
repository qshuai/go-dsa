package types

import (
	"reflect"
	"testing"
)

func TestHeap_Sort(t *testing.T) {
	tests := []struct {
		name string
		h    Heap[int]
		want Heap[int]
	}{
		{
			name: "case-1",
			h:    []int{9, 6, 3, 1, 5},
			want: []int{1, 3, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.h.Sort(); !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("Sort() = %v want %v", tt.h, tt.want)
			}
		})
	}
}

func TestNewHeap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want Heap[int]
	}{
		{
			name: "case-1",
			args: []int{8, 3, 9, 1, 2, 5},
			want: []int{9, 3, 8, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeap(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHeapWithCapacity(t *testing.T) {
	type args struct {
		nums     []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want Heap[int]
	}{
		{
			name: "case-1",
			args: args{
				nums:     []int{1, 4, 6, 3, 7, 9},
				capacity: 8,
			},
			want: []int{9, 7, 6, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHeapWithCapacity(tt.args.nums, tt.args.capacity)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeapWithCapacity() = %v, want %v", got, tt.want)
			}
			if cap(got) != tt.args.capacity {
				t.Errorf("NewHeapWithCapacity[capacity]() = %d want %d", cap(got), tt.args.capacity)
			}
		})
	}
}

func TestHeap_DeleteHeapTop(t *testing.T) {
	tests := []struct {
		name       string
		h          Heap[int]
		want       Heap[int]
		deletedEle int
	}{
		{
			name:       "case-1",
			h:          []int{9, 6, 7, 4, 3, 1},
			want:       []int{7, 6, 1, 4, 3},
			deletedEle: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEle := tt.h.DeleteHeapTop()
			if !reflect.DeepEqual(gotEle, tt.deletedEle) {
				t.Errorf("DeleteHeapTop() = %v, want %v", gotEle, tt.deletedEle)
			}
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("DeleteHeapTop[remaining]() = %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {
	tests := []struct {
		name string
		h    Heap[int]
		args int
		want Heap[int]
	}{
		{
			name: "case-1",
			h:    NewHeapWithCapacity([]int{9, 6, 7, 4}, 5),
			args: 5,
			want: []int{9, 6, 7, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.h.Insert(tt.args); !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("Insert() = %v want %v", tt.h, tt.want)
			}
		})
	}
}

func TestHeap_MaxHeapify(t *testing.T) {
	tests := []struct {
		name  string
		h     Heap[int]
		index int
		want  Heap[int]
	}{
		{
			name:  "case-1",
			h:     []int{0, 5, 6, 4, 2, 1},
			index: 0,
			want:  []int{6, 5, 1, 4, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.h.MaxHeapify(len(tt.h), tt.index); !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("heapify() = %v want %v", tt.h, tt.want)
			}
		})
	}
}

func TestHeap_MinHeapify(t *testing.T) {
	tests := []struct {
		name  string
		h     Heap[int]
		index int
		want  Heap[int]
	}{
		{
			name:  "case-1",
			h:     []int{9, 4, 5, 1, 0},
			index: 0,
			want:  []int{4, 0, 5, 1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.h.MinHeapify(len(tt.h), tt.index); !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("MinHeapify() = %v want %v", tt.h, tt.want)
			}
		})
	}
}
