package leetcode

// https://leetcode.com/problems/first-missing-positive
//
// Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
// You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.
//
// Example 1:
// Input: nums = [1,2,0]
// Output: 3
// Explanation: The numbers in the range [1,2] are all in the array.
//
// Example 2:
// Input: nums = [3,4,-1,1]
// Output: 2
// Explanation: 1 is in the array but 2 is missing.
//
// Example 3:
// Input: nums = [7,8,9,11,12]
// Output: 1
// Explanation: The smallest positive integer 1 is missing.
//
// Constraints:
// 1 <= nums.length <= 10^5
// -2^31 <= nums[i] <= 2^31 - 1
func firstMissingPositive(nums []int) int {
	// 解题思路：使得[0, left)区域的数都满足nums[i] = i + 1，[r:len(nums)]的数都是“非法数据”。
	// 那么答案就是left + 1
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == left+1 {
			left++
			continue
		}

		// 非法数据：
		// 1. nums[left] <= left 超过最小有效范围的数据，eg. [1,2,1] 其第二个1就是这种情况；eg. [2,1,-1] 其中-1也属于这种情况
		// 2. nums[left] > right+1 超过最大有效范围的数据
		// 3. nums[left] == nums[nums[left]-1] 重复数据
		if nums[left] <= left || nums[left] > right+1 || nums[left] == nums[nums[left]-1] {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			continue
		}

		nums[left], nums[nums[left]-1] = nums[nums[left]-1], nums[left]
	}

	return left + 1
}
