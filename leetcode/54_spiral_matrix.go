package leetcode

// https://leetcode.com/problems/spiral-matrix/

// spiralOrder Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Constraints:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return nil
	}

	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for left <= right && top <= bottom {
		// 从左向右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		// 从上往下
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// 关键点
		if left > right || top > bottom {
			break
		}

		// 从右向左
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		bottom--

		// 从下往上
		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
	}

	return ans
}
