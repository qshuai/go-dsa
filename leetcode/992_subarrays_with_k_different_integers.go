package leetcode

// https://leetcode.com/problems/subarrays-with-k-different-integers/
//
// Given an integer array nums and an integer k, return the number of good subarrays of nums.
// A good array is an array where the number of different integers in that array is exactly k.
// For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
// A subarray is a contiguous part of an array.
//
// Example 1:
// Input: nums = [1,2,1,2,3], k = 2
// Output: 7
// Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
//
// Example 2:
// Input: nums = [1,2,1,3,4], k = 3
// Output: 3
// Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
//
// Constraints:
// 1 <= nums.length <= 2 * 10^4
// 1 <= nums[i], k <= nums.length
func subarraysWithKDistinct(nums []int, k int) int {
	return numsOfMostKinds(nums, k) - numsOfMostKinds(nums, k-1)
}

func numsOfMostKinds(nums []int, k int) int {
	counts := make([]int, len(nums)+1)
	var numbers, res int

	for right, left := 0, 0; right < len(nums); right++ {
		if counts[nums[right]] == 0 {
			numbers++
		}
		counts[nums[right]]++

		for numbers > k {
			if counts[nums[left]] == 1 {
				numbers--
			}

			counts[nums[left]]--
			left++
		}

		res += right - left + 1
	}

	return res
}
