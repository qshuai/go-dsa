package leetcode

// https://leetcode.com/problems/two-sum/

// twoSum Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[nums[i]]; ok {
			return []int{index, i}
		}

		m[target-nums[i]] = i
	}

	return nil
}

func twoSumN2(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
