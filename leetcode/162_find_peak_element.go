package leetcode

// https://leetcode.com/problems/find-peak-element/

// findPeakElement A peak element is an element that is strictly greater than its neighbors.
// Given a 0-indexed integer array nums, find a peak element, and return its
// index. If the array contains multiple peaks, return the index to any of the peaks.
// You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is
// always considered to be strictly greater than a neighbor that is outside the array.
// You must write an algorithm that runs in O(log n) time.
//
// Constraints:
// 1 <= nums.length <= 1000
// -2^31 <= nums[i] <= 2^31 - 1
// nums[i] != nums[i + 1] for all valid i.
func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if mid == 0 {
			left = 1
		} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] > nums[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
