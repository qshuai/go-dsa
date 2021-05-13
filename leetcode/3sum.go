package leetcode

import (
	"sort"
)

// Reference: https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}

	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			continue
		}

		head := i + 1
		tail := len(nums) - 1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] {
					head++
				}

				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}

				head++
				tail--
			} else if sum > 0 {
				tail--
			} else {
				head++
			}
		}
	}

	return result
}
