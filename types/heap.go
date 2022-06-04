package types

import (
	"golang.org/x/exp/constraints"
)

// Heap represents a max-heap
type Heap[T constraints.Ordered] []T

type HeapDir byte

const (
	HeapMaxDir HeapDir = iota
	HeapMinDir
)

// Insert adds a new element to heap
func (h *Heap[T]) Insert(ele T) {
	// todo 这里为了简单实现，不做slice的自动扩容
	if len(*h) == cap(*h) {
		panic("slice not enough space")
	}

	// 将元素添加为最后一个叶子节点
	*h = append(*h, ele)

	i := len(*h) - 1
	father := (i - 1) >> 1
	for father >= 0 && (*h)[i] > (*h)[father] {
		// swap
		(*h)[i], (*h)[father] = (*h)[father], (*h)[i]

		i = father
		father = (i - 1) >> 1
	}
}

// DeleteHeapTop removes the root node of heap
func (h *Heap[T]) DeleteHeapTop() (ele T) {
	if len(*h) <= 0 {
		return
	}

	// 将最后一个叶子节点覆盖到堆顶位置
	ele = (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	// length - 1
	*h = (*h)[:len(*h)-1]

	h.MaxHeapify(len(*h)-1, 0)

	return
}

// Sort heap in natural order
func (h Heap[T]) Sort() {
	if len(h) <= 0 {
		return
	}

	for i := len(h) - 1; i > 0; i-- {
		h[i], h[0] = h[0], h[i]

		for j := 0; j < i<<1-1; j++ {
			h.MaxHeapify(i, 0)
		}
	}
}

// NewHeap return a buildHeap sequence
func NewHeap[T constraints.Ordered](nums []T) Heap[T] {
	if len(nums) <= 0 {
		return nil
	}

	heap := (Heap[T])(nums)
	heap.Heapify(HeapMaxDir)

	return heap
}

// NewHeapWithCapacity return a buildHeap sequence with capacity
func NewHeapWithCapacity[T constraints.Ordered](nums []T, capacity int) Heap[T] {
	if len(nums) <= 0 {
		return nil
	}
	if len(nums) > capacity {
		panic("capacity not enough")
	}

	// grow to capacity
	if capacity > len(nums) {
		temp := make([]T, 0, capacity)
		nums = append(temp, nums...)
	}

	heap := (Heap[T])(nums)
	heap.Heapify(HeapMaxDir)

	return heap
}

// Heapify according to dir
func (h Heap[T]) Heapify(dir HeapDir) {
	switch dir {
	case HeapMinDir:
		for i := len(h)>>1 - 1; i >= 0; i-- {
			h.MinHeapify(len(h), i)
		}
	case HeapMaxDir:
		for i := len(h)>>1 - 1; i >= 0; i-- {
			h.MaxHeapify(len(h), i)
		}
	default:
		panic("unrecognised heap dir")
	}
}

// MinHeapify maintains a min-heap property from top to bottom
func (h Heap[T]) MinHeapify(length int, index int) {
	for {
		minPos := index
		if index<<1+1 < length && h[index] > h[index<<1+1] {
			minPos = index<<1 + 1
		}
		if index<<1+2 < length && h[minPos] > h[index<<1+2] {
			minPos = index<<1 + 2
		}
		if minPos == index {
			break
		}

		h[index], h[minPos] = h[minPos], h[index]
		index = minPos
	}
}

// MaxHeapify maintains a max-heap property from top to bottom
func (h Heap[T]) MaxHeapify(length int, index int) {
	for {
		maxPos := index
		if index<<1+1 < length && h[index] < h[index<<1+1] {
			maxPos = index<<1 + 1
		}
		if index<<1+2 < length && h[maxPos] < h[index<<1+2] {
			maxPos = index<<1 + 2
		}
		if maxPos == index {
			break
		}

		h[index], h[maxPos] = h[maxPos], h[index]
		index = maxPos
	}
}
