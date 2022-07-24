package leetcode

// numIslands Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
//
// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfsIslands(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfsIslands(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		// 越界了
		return
	}
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfsIslands(grid, i-1, j) // 下一列（左）
	dfsIslands(grid, i+1, j) // 下一列（右）
	dfsIslands(grid, i, j-1) // 下一行（上）
	dfsIslands(grid, i, j+1) // 下一行（下）
}
