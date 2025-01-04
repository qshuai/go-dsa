package leetcode

import "github.com/qshuai/go-dsa/types"

// https://leetcode.com/problems/kth-largest-element-in-an-array/

// findKthLargest 构建k个元素的小顶堆，遍历将k到n-1的元素，如果比堆顶元素大就加入，否则忽略，
// 最终堆顶元素即为kth大的元素
func findKthLargest(nums []int, k int) int {
	(*types.Heap[int])(&nums).Heapify(types.HeapMinDir)

	for i := k; i < len(nums); i++ {
		if nums[i] <= nums[0] {
			continue
		}

		nums[0] = nums[i]
		(*types.Heap[int])(&nums).MinHeapify(k, 0)
	}

	return nums[0]
}

// findKthLargest2 Solution-2 从大到小，找到排名第K的元素(值相同的元素，
// 排名加1，也就是说排名是连续的[可以查看测试用例])
func findKthLargest2(nums []int, k int) int {
	if k > len(nums) || k <= 0 || len(nums) <= 0 {
		return -1
	}

	var offset int
	for {
		pivot := partition(nums)
		if pivot+1+offset == k {
			return nums[pivot]
		} else if pivot+1+offset > k {
			nums = nums[:pivot]
		} else {
			offset += pivot
			nums = nums[pivot:]
		}
	}
}

// 分区，大的元素放在前面（索引小的位置）
func partition(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	i, j := 0, len(nums)-1
	for {
		if i >= j {
			break
		}

		for i < j && nums[i] >= nums[j] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
		for i < j && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}
