package leetcode

// https://leetcode.com/problems/sort-array-by-parity-ii/
//
// Given an array of integers nums, half of the integers in nums are odd, and the other half are even.
// Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.
// Return any answer array that satisfies this condition.
//
// Example 1:
// Input: nums = [4,2,5,7]
// Output: [4,5,2,7]
// Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
//
// Example 2:
// Input: nums = [2,3]
// Output: [2,3]
//
// Constraints:
// 2 <= nums.length <= 2 * 10^4
// nums.length is even.
// Half of the integers in nums are even.
// 0 <= nums[i] <= 1000
//
// Follow Up: Could you solve it in-place?
func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		for ; even < len(nums); even += 2 {
			if nums[even]&1 != 0 {
				break
			}
		}

		for ; odd < len(nums); odd += 2 {
			if nums[odd]&1 == 0 {
				break
			}
		}

		if even < len(nums) && odd < len(nums) {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}

	return nums
}

func sortArrayByParityII2(nums []int) []int {
	n := len(nums)

	for even, odd := 0, 1; even < n && odd < n; {
		if nums[n-1]&1 == 0 {
			// 偶数
			nums[even], nums[n-1] = nums[n-1], nums[even]
			even += 2
		} else {
			nums[odd], nums[n-1] = nums[n-1], nums[odd]
			odd += 2
		}
	}

	return nums
}
