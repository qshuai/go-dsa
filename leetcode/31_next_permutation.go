package leetcode

import "sort"

// https://leetcode.com/problems/next-permutation/

// nextPermutation find the next permutation of nums.
//
// Constraints:
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
// The replacement must be in place and use only constant extra memory.
func nextPermutation(nums []int) {
	// 从右向左遍历，为的是获取下一个比之前大一些的排列
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				// 将大的元素向前交换
				nums[i], nums[j] = nums[j], nums[i]
				// 并将之后的元素按从小到大排序，以获取贴近的下一个排列
				sort.Ints(nums[i+1:])
				return
			}
		}
	}

	// 如果没有找到，说明给出的序列数最大的一个，需要返回最小的排列，作为下一个排列
	sort.Ints(nums)
}
