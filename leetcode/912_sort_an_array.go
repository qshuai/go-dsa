package leetcode

// https://leetcode.com/problems/sort-an-array/
// 其中一个不错的Solution分析：https://leetcode.com/problems/sort-an-array/solutions/1401412/c-clean-code-solution-fastest-all-15-sorting-methods-detailed/
//
// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity
// and with the smallest space complexity possible.
//
// Example 1:
// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3),
// while the positions of other numbers are changed (for example, 1 and 5).
//
// Example 2:
// Input: nums = [5,1,1,2,0,0]
// Output: [0,0,1,1,2,5]
// Explanation: Note that the values of nums are not necessarily unique.
//
// Constraints:
// 1 <= nums.length <= 5 * 10^4
// -5 * 10^4 <= nums[i] <= 5 * 10^4
func sortArray(nums []int) []int {
	mergeSortArray(nums)
	return nums
}

func mergeSortArray(nums []int) {
	if len(nums) == 1 {
		return
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}

		return
	}

	pivot := len(nums) / 2
	mergeSortArray(nums[:pivot])
	mergeSortArray(nums[pivot:])

	mergeSortedArray(nums[:pivot], nums[pivot:])
}

func mergeSortedArray(left, right []int) {
	tmp := make([]int, len(left))
	copy(tmp, left)

	left = left[:0]
	var i, j int
	for i < len(tmp) && j < len(right) {
		if tmp[i] <= right[j] {
			left = append(left, tmp[i])
			i++
		} else {
			left = append(left, right[j])
			j++
		}
	}
	if i != len(tmp) {
		left = append(left, tmp[i:]...)
	}
	if j != len(right) {
		left = append(left, right[j:]...)
	}
}
