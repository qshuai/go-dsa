package leetcode

// https://leetcode.com/problems/permutations/

// permute returns all the possible permutations
//
// Constraints:
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.
func permute(nums []int) [][]int {
	size := 1
	for i := len(nums); i > 1; i-- {
		size *= i
	}

	visited := make([]bool, len(nums))
	collector := make([][]int, 0, size)
	path := make([]int, 0, len(nums))
	dfs(nums, &path, 0, visited, &collector)

	return collector
}

// dfs 深度优先搜索，经典的回溯算法，步骤为记录现场、递归、恢复现场
// 和二叉树的深度优先搜索相比，回溯算法需要记录遍历节点的数量和访问状态，并在终止节点处进行数据收集
func dfs(nums []int, path *[]int, depth int, visited []bool, collector *[][]int) {
	if depth == len(nums) {
		item := make([]int, len(nums))
		copy(item, *path)
		*collector = append(*collector, item)
		return
	}

	for idx := range nums {
		if visited[idx] {
			continue
		}

		// 记录现场
		*path = append(*path, nums[idx])
		visited[idx] = true
		// 递归搜索
		dfs(nums, path, depth+1, visited, collector)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
		visited[idx] = false
	}
}

// permute Solution 2: 该算法也是用回溯的思想，其状态是改变原始数组的元素排列。
// 相比通过额外的变量来记录状态实现回溯的算法，该算法不太容易理解
func permute2(nums []int) [][]int {
	size := 1
	for i := len(nums); i > 1; i-- {
		size *= i
	}

	collector := make([][]int, 0, size)
	permuteUsingSwap(nums, len(nums), &collector)
	return collector
}

func permuteUsingSwap(nums []int, k int, collector *[][]int) {
	if k == 1 {
		*collector = append(*collector, nums)
	}

	for i := 0; i < k; i++ {
		nums[i], nums[k-1] = nums[k-1], nums[i]
		permuteUsingSwap(nums, k-1, collector)
		nums[i], nums[k-1] = nums[k-1], nums[i]
	}
}
