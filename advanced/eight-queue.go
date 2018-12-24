package main

import "fmt"

var arr [8][8]int
var n = 1

func main() {
	play(0)
}

func play(row int) {
	for i := 0; i < 8; i++ {
		if check(row, i) {
			arr[row][i] = 1
			if row == 7 {
				show()
			} else {
				play(row + 1)
			}
			arr[row][i] = 0
		}
	}
}

func show() {
	fmt.Println("第", n, "种解法：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Print("\n")
	}
	n++
}

func check(row, col int) bool {
	//上
	for i := row - 1; i >= 0; i-- {
		if arr[i][col] == 1 {
			return false
		}
	}

	for i := 0; i < 8; i++ {
		if row-i > -1 && col-i > -1 {
			if arr[row-i][col-i] == 1 {
				return false
			}
		}

	}

	for i := 0; i < 8; i++ {
		if row-i > -1 && col+i < 8 {
			if arr[row-i][col+i] == 1 {
				return false
			}
		}
	}

	return true
}
