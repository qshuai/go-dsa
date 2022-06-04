package types

import (
	"golang.org/x/exp/constraints"
)

// Heap represents a max-heap
type Heap[T constraints.Ordered] []T

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

	var i int
	for {
		maxPos := i
		if i<<1+1 < len(*h) && (*h)[i<<1+1] > (*h)[i] {
			maxPos = i<<1 + 1
		}
		if i<<1+2 < len(*h) && (*h)[maxPos] < (*h)[i<<1+2] {
			maxPos = i<<1 + 2
		}
		if maxPos == i {
			break
		}

		(*h)[i], (*h)[maxPos] = (*h)[maxPos], (*h)[i]
		i = maxPos
	}

	return
}

// Sort heap in natural order
func (h Heap[T]) Sort() {
	if len(h) <= 0 {
		return
	}

	for i := len(h) - 1; i > 0; i-- {
		h[i], h[0] = h[0], h[i]

		h.heapifyToBottom(i)
	}
}

// NewHeap return a heapifyToTop sequence
func NewHeap[T constraints.Ordered](nums []T) Heap[T] {
	if len(nums) <= 0 {
		return nil
	}

	heap := (Heap[T])(nums)
	heap.heapifyToTop()
	return heap
}

// NewHeapWithCapacity return a heapifyToTop sequence with capacity
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
	heap.heapifyToTop()
	return heap
}

// heapifyToTop maintains heap property from bottom to top
func (h Heap[T]) heapifyToTop() {
	for i := len(h)>>1 - 1; i >= 0; i-- {
		tmp := i
		for {
			maxPos := i
			if i<<1+1 < len(h) && h[i<<1+1] > h[i] {
				maxPos = i<<1 + 1
			}
			if i<<1+2 < len(h) && h[i<<1+2] > h[maxPos] {
				maxPos = i<<1 + 2
			}
			if maxPos == i {
				break
			}

			h[i], h[maxPos] = h[maxPos], h[i]
			i = maxPos
		}
		i = tmp
	}
}

// heapifyToBottom maintains heap property from top to bottom
func (h Heap[T]) heapifyToBottom(length int) {
	if length <= 0 {
		return
	}

	for i := 0; i < length-1; i++ {
		tmp := i
		for {
			maxPos := i
			if i<<1+1 < length && h[i<<1+1] > h[i] {
				maxPos = i<<1 + 1
			}
			if i<<1+2 < length && h[i<<1+2] > h[maxPos] {
				maxPos = i<<1 + 2
			}
			if maxPos == i {
				break
			}

			// swap
			h[i], h[maxPos] = h[maxPos], h[i]
			i = maxPos
		}
		i = tmp
	}
}
