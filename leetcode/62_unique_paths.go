package leetcode

// https://leetcode.com/problems/unique-paths/description/
//
// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either
// down or right at any point in time.
// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach
// the bottom-right corner.
// The test cases are generated so that the answer will be less than or equal to 2 * 109.
//
// Constraints:
// 1 <= m, n <= 100
func uniquePaths(m int, n int) int {
	// dp[i][j]代表从[0,0]位置移动到[i,j]位置的路径数量
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)

		// 第一行统一设置为1
		if idx == 0 {
			for i := 0; i < n; i++ {
				dp[0][i] = 1
			}
		}

		// 	第一列统一设置为1
		dp[idx][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}
