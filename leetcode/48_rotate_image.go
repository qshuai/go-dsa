package leetcode

// https://leetcode.com/problems/rotate-image/

// rotateMatrix You are given an n x n 2D matrix representing an image, rotate the image
// by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix
// directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Constraints:
// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
func rotateMatrix(matrix [][]int) {
	for len(matrix) <= 1 {
		return
	}

	// 上下翻转
	i, j := 0, len(matrix)-1
	for i < j {
		for k := 0; k < len(matrix[0]); k++ {
			matrix[i][k], matrix[j][k] = matrix[j][k], matrix[i][k]
		}
		i++
		j--
	}

	// 对角线翻转
	for i := 0; i < len(matrix)-1; i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
