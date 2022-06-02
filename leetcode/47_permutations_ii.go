package leetcode

// https://leetcode.com/problems/permutations-ii/

// permuteUnique returns all possible unique permutations in any order.
//
// Constraints:
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
// nums might contain duplicates
func permuteUnique(nums []int) [][]int {
	collector := make([][]int, 0)
	path := make([]int, 0, len(nums))
	visited := make([]bool, len(nums))
	dfsWithClip(nums, 0, &path, visited, &collector)
	return collector
}

func dfsWithClip(nums []int, depth int, path *[]int, visited []bool, collector *[][]int) {
	if depth == len(nums) {
		item := make([]int, len(nums))
		copy(item, *path)
		*collector = append(*collector, item)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 该路径上已经访问的元素
		if visited[i] {
			continue
		}
		// 同一层去除重复元素[important]
		if i > 0 && !visited[i-1] && nums[i] == nums[i-1] {
			continue
		}

		*path = append(*path, nums[i])
		visited[i] = true
		dfsWithClip(nums, depth+1, path, visited, collector)
		*path = (*path)[:len(*path)-1]
		visited[i] = false
	}
}
