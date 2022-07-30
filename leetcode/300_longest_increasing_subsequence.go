package leetcode

import "github.com/qshuai/go-dsa/utils"

// https://leetcode.com/problems/longest-increasing-subsequence/

// lengthOfLIS Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
// Constraints:
// 1 <= nums.length <= 2500
// -104 <= nums[i] <= 104
// 使用动态规划的思路
func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	// 假设索引为i，值为以nums[i]结尾的最长子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = utils.Max(dp[i], dp[j]+1)
			}
		}

		max = utils.Max(max, dp[i])
	}

	return max
}

// lengthOfLIS2 solution-2 使用贪心算法，在查找元素部分使用了二分查找(降低查找的时间复杂度)
func lengthOfLIS2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	d := make([]int, len(nums))
	d[0] = nums[0]
	var offset int
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[offset] {
			// 如果当前元素比offset出的元素大，直接追加（贪心）
			offset++
			d[offset] = nums[i]
		} else if nums[i] < d[offset] {
			// 之所以能使用二分查找，因为d元素是有序的
			// 将当前元素nums[i]插入到d中，替换掉原先的元素（第一个小于等于该元素的位置），保持数组的有序性
			begin, end := 0, offset
			for begin <= end {
				mid := begin + (end-begin)>>1
				if d[mid] < nums[i] {
					begin = mid + 1
				} else {
					end = mid - 1
				}
			}

			d[begin] = nums[i]
		}
	}

	return offset + 1
}
