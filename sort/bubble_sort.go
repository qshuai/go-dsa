package sort

import (
	"math"

	"github.com/qshuai/go-dsa/types"
)

// bubbleSort 冒泡排序（稳定排序；时间复杂度O(n^2)）
// 基本思路：一次遍历每一个元素，如果该元素比下一个元素大，就交换两个元素的位置。这样的过程
// 执行n-1即可达到有序的状态
func bubbleSort(arr []int) []int {
	var sorted bool
	for i := 0; i < len(arr)-1 && !sorted; i++ {
		sorted = true
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				sorted = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
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

// listBubbleSortChangValue 单向链表的冒泡排序算法（元素值为int类型）
// Constraints: 可以修改元素字段值
func listBubbleSortChangValue(node *types.ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	cursor := node
	var swapCnt int
	for {
		if cursor.Next == nil {
			if swapCnt == 0 {
				break
			}

			cursor = node
			swapCnt = 0
		}

		if cursor.Value.(int) <= cursor.Next.Value.(int) {
			cursor = cursor.Next
			continue
		}

		swapCnt++
		cursor.Value, cursor.Next.Value = cursor.Next.Value, cursor.Value
	}
}

// DoublyListBubbleSort 双向链表的冒泡排序算法（元素值为int类型）
// Constraints: 不可以修改元素值
func doublyListBubbleSort(head *types.DListNode) *types.DListNode {
	if head == nil || head.Next == nil {
		return head
	}

	guard := &types.DListNode{
		Value: math.MinInt,
		Prev:  nil,
		Next:  head,
	}
	head.Prev = guard

	cursor := head
	prev := guard
	var swapCnt int
	for {
		if cursor.Next == nil {
			if swapCnt == 0 {
				break
			}

			cursor = guard.Next
			prev = guard
			swapCnt = 0
			continue
		}

		if cursor.Value.(int) <= cursor.Next.Value.(int) {
			cursor = cursor.Next
			prev = cursor
			continue
		}

		swapCnt++
		tmp := cursor.Next.Next
		cursor.Next.Next = cursor
		cursor.Next.Prev = prev
		prev.Next = cursor.Next
		prev = cursor.Next
		cursor.Prev = cursor.Next
		cursor.Next = tmp
		if tmp != nil {
			tmp.Prev = cursor
		}
	}

	guard.Next.Prev = nil
	return guard.Next
}

// DoublyListBubbleSortChangeValue 双向链表的冒泡排序算法（元素值为int类型）
// Constraints: 可以修改元素值
func doublyListBubbleSortChangeValue(head *types.DListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var swapCnt int
	cursor := head
	for {
		if cursor.Next == nil {
			if swapCnt == 0 {
				break
			}

			cursor = head
			swapCnt = 0
			continue
		}

		if cursor.Value.(int) <= cursor.Next.Value.(int) {
			cursor = cursor.Next
			continue
		}

		swapCnt++
		cursor.Value, cursor.Next.Value = cursor.Next.Value, cursor.Value
		cursor = cursor.Next
	}
}
