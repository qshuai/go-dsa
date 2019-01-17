package leetcode

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
