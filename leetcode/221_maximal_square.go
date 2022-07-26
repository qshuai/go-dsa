package leetcode

// https://leetcode.com/problems/maximal-square/

// maximalSquare Given an m x n binary matrix filled with 0's and 1's, find the largest square containing
// only 1's and return its area.
//
// Constraints:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] is '0' or '1'.
func maximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 {
		return 0
	}

	// dp[i][j]代表以matrix[i][j]为正方形右下角时的最大边长
	dp := make([][]int, len(matrix))
	var maxLength int
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				if maxLength == 0 {
					maxLength = 1
				}

				continue
			}

			// 地推公式
			dp[i][j] = minLength(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			if maxLength < dp[i][j] {
				maxLength = dp[i][j]
			}
		}
	}

	return maxLength * maxLength
}

func minLength(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	}

	if b < c {
		return b
	}

	return c
}
