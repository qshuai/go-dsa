package leetcode

// https://leetcode.com/problems/n-queens/

const (
	QueenPos  byte = 'Q'
	EmptyPost byte = '.'
)

// solveNQueens Given an integer n, return all distinct solutions to the n-queens puzzle.
// You may return the answer in any order.
func solveNQueens(n int) [][]string {
	collector := make([][]string, 0)
	pos := make([]int, n+1)

	blank := make([]byte, n)
	for i := 0; i < n; i++ {
		blank[i] = EmptyPost
	}

	dfsQueens(n, 0, pos, blank, &collector)
	return collector
}

func dfsQueens(n, row int, pos []int, blank []byte, collector *[][]string) {
	if row >= n {
		item := make([]string, 0, 4)
		for i := 1; i < len(pos); i++ {
			cp := make([]byte, n)
			copy(cp, blank)
			cp[pos[i]-1] = QueenPos
			item = append(item, string(cp))
		}
		*collector = append(*collector, item)

		return
	}

	for i := 1; i <= n; i++ {
		// 判断是否可以放在此位置
		if ok := checkPos(n, row, i, pos); !ok {
			continue
		}

		pos[row+1] = i
		dfsQueens(n, row+1, pos, blank, collector)
		pos[row+1] = 0
	}
}

func checkPos(n, row, col int, pos []int) bool {
	if row == 0 {
		return true
	}

	for i := 1; i <= n; i++ {
		if pos[i] <= 0 {
			break
		}

		if col == pos[i] {
			return false
		}
		// 处于对角线
		if col == pos[i]+(row-i+1) || col == pos[i]-(row-i+1) {
			return false
		}
	}

	return true
}
