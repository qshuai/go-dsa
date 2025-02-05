package leetcode

import "github.com/qshuai/go-dsa/utils"

// https://leetcode.com/problems/set-mismatch
//
// You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error,
// one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and
// loss of another number.
// You are given an integer array nums representing the data status of this set after the error.
// Find the number that occurs twice and the number that is missing and return them in the form of an array.
//
// Example 1:
// Input: nums = [1,2,2,4]
// Output: [2,3]
//
// Example 2:
// Input: nums = [1,1]
// Output: [1,2]
//
// Constraints:
// 2 <= nums.length <= 10^4
// 1 <= nums[i] <= 10^4
func findErrorNums(nums []int) []int {
	set := make([]int, len(nums))
	res := make([]int, 2)
	for _, num := range nums {
		set[num-1]++
		if set[num-1] > 1 {
			res[0] = num
		}
	}

	for idx, count := range set {
		if count == 0 {
			res[1] = idx + 1
		}
	}

	return res
}

func findErrorNums2(nums []int) []int {
	var actualSum, uniqueSum int
	set := make([]int, len(nums))
	for _, num := range nums {
		actualSum += num

		set[num-1]++
		if set[num-1] == 1 {
			uniqueSum += num
		}
	}

	expectedSum := len(nums) * (1 + len(nums)) / 2
	return []int{actualSum - uniqueSum, expectedSum - uniqueSum}
}

func findErrorNums3(nums []int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		idx := utils.Abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = nums[idx] * -1
		} else {
			res[0] = utils.Abs(nums[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			// 该位置的数大于0，说明没有元素指向过该位置，使其变成负数；而指向是通过数值-1得到的，所以i+1就是缺失的数
			res[1] = i + 1
			return res
		}
	}

	return nil
}
