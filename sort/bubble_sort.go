package sort

import (
	"math"

	"github.com/qshuai/go-dsa/types"
)

// bubbleSort 冒泡排序（稳定排序；时间复杂度O(n^2)）
func bubbleSort(arr []int) []int {
	var sorted bool
	for i := 0; i < len(arr)-1 && !sorted; i++ {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}

// listBubbleSort 单向链表的冒泡排序算法（元素值为int类型）
// Constraints: 不修改元素字段值
func listBubbleSort(node *types.ListNode) *types.ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	// 哨兵节点
	guard := &types.ListNode{
		Value: math.MinInt,
		Next:  node,
	}

	cursor := guard.Next
	prev := guard
	var swapCount int
	for {
		if cursor.Next == nil {
			if swapCount == 0 {
				break
			}

			// 最后一个节点了
			cursor = guard.Next
			prev = guard
			swapCount = 0
			continue
		}

		if cursor.Value.(int) <= cursor.Next.Value.(int) {
			// 下一个节点
			prev = cursor
			cursor = cursor.Next
			continue
		}

		swapCount++
		tmp := cursor.Next.Next
		cursor.Next.Next = cursor
		prev.Next = cursor.Next
		prev = cursor.Next
		cursor.Next = tmp
	}

	return guard.Next
}
