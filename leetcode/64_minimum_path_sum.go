package leetcode

import (
	math2 "math"

	"github.com/qshuai/go-dsa/utils"
)

// https://leetcode.com/problems/minimum-path-sum/

// minPathSum Given a m x n grid filled with non-negative numbers, find
// a path from top left to bottom right, which minimizes the sum of all
// numbers along its path.
// Note: You can only move either down or right at any point in time.
//
// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
func minPathSum(grid [][]int) int {
	// dp[i][j] 代表移动到i，j位置的最短路径
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

// minPathSum2 Solution-2 using dfs
func minPathSum2(grid [][]int) int {
	path, min := grid[0][0], math2.MaxInt
	dfsMinPath(grid, 0, 0, &path, &min)
	return min
}

func dfsMinPath(grid [][]int, i, j int, path, min *int) {
	if *path >= *min {
		// 优化：提前结束
		return
	}

	if i == len(grid)-1 && j == len(grid[0])-1 {
		// 到达终点
		if *path < *min {
			*min = *path
		}
		return
	}

	// 向下
	if i+1 < len(grid) {
		*path += grid[i+1][j]
		dfsMinPath(grid, i+1, j, path, min)
		*path -= grid[i+1][j]
	}

	// 向右
	if j+1 < len(grid[0]) {
		*path += grid[i][j+1]
		dfsMinPath(grid, i, j+1, path, min)
		*path -= grid[i][j+1]
	}
}
