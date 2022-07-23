package leetcode

// https://leetcode.com/problems/subsets/

// subsets Given an integer array nums of unique elements, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Constraints:
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
func subsets(nums []int) [][]int {
	collector := make([][]int, 0, 1)
	dfsSubSets(nums, 0, &collector, new([]int))
	return collector
}

func dfsSubSets(nums []int, idx int, collector *[][]int, tmp *[]int) {
	if idx == len(nums) {
		cp := make([]int, len(*tmp))
		copy(cp, *tmp)
		*collector = append(*collector, cp)
		return
	}

	// 	不添加当前元素
	dfsSubSets(nums, idx+1, collector, tmp)

	// 添加当前元素
	*tmp = append(*tmp, nums[idx])
	dfsSubSets(nums, idx+1, collector, tmp)
	*tmp = (*tmp)[:len(*tmp)-1]
}
