package leetcode

// https://leetcode.com/problems/combination-sum/

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers
// sum to target. You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	data := make([]int, 0)
	sum(candidates, 0, target, data, &ret)
	return ret
}

func sum(candidates []int, idx, target int, data []int, collector *[][]int) {
	if target < 0 || idx > len(candidates) {
		return
	}
	if target == 0 {
		tmp := make([]int, len(data))
		copy(tmp, data)
		*collector = append(*collector, tmp)
	}

	for i := idx; i < len(candidates); i++ {
		data = append(data, candidates[i])
		sum(candidates, i, target-candidates[i], data, collector)
		data = data[:len(data)-1]
	}
}
