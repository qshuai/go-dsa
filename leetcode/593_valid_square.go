package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/valid-square/

// validSquare Given the coordinates of four points in 2D space p1, p2, p3 and p4, return true if
// the four points construct a square.
// The coordinate of a point pi is represented as [xi, yi]. The input is not given in any order.
// A valid square has four equal sides with positive length and four equal angles (90-degree angles).
//
// Constraints:
// p1.length == p2.length == p3.length == p4.length == 2
// -10^4 <= xi, yi <= 10^4
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d1 := calcDistanceSquare(p1, p2)
	d2 := calcDistanceSquare(p1, p3)
	d3 := calcDistanceSquare(p1, p4)
	d4 := calcDistanceSquare(p2, p3)
	d5 := calcDistanceSquare(p2, p4)
	d6 := calcDistanceSquare(p3, p4)

	dists := []int{d1, d2, d3, d4, d5, d6}
	sort.Ints(dists)
	return dists[0] > 0 && // 防止为同一个点的情况
		dists[0] == dists[1] && dists[1] == dists[2] && dists[2] == dists[3] && // 四条边相等
		dists[4] == dists[5] // 对角线相等，和菱形进行区分
}

func calcDistanceSquare(p []int, q []int) int {
	x := p[0] - q[0]
	y := p[1] - q[1]

	return x*x + y*y
}
