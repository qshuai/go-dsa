package leetcode

import (
	"math"
	"sort"

	"github.com/qshuai/go-dsa/utils"
)

// https://leetcode.com/problems/heaters
//
// Winter is coming! During the contest, your first job is to design a standard heater with a fixed warm radius to warm all the houses.
// Every house can be warmed, as long as the house is within the heater&#39;s warm radius range.
// Given the positions of houses and heaters on a horizontal line, return the minimum radius standard of heaters so that those heaters could cover all houses.
// Notice that all the heaters follow your radius standard, and the warm radius will the same.
//
// Example 1:
// Input: houses = [1,2,3], heaters = [2]
// Output: 1
// Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
//
// Example 2:
// Input: houses = [1,2,3,4], heaters = [1,4]
// Output: 1
// Explanation: The two heaters were placed at positions 1 and 4. We need to use a radius 1 standard, then all the houses can be warmed.
//
// Example 3:
// Input: houses = [1,5], heaters = [2]
// Output: 3
//
// Constraints:
// 1 <= houses.length, heaters.length <= 3 * 10^4
// 1 <= houses[i], heaters[i] <= 10^9
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	var res int
	for i, j := 0, 0; i < len(houses); {
		r1 := utils.Abs(houses[i] - heaters[j])
		r2 := math.MaxInt
		if j < len(heaters)-1 {
			r2 = utils.Abs(houses[i] - heaters[j+1])
		}
		if r1 < r2 {
			i++
			res = max(res, r1)
		} else {
			j++
		}
	}

	return res
}
