package other

import (
	"math"
)

// maxWeightUsingDfs 已知给出各物品的重量weight，以及背包能够承受的最大重量maxWeight
// 在不超过背包最大承受重量的前提下，能够装载的最大重量
//
// Constraints:
// 1. 物品不可分割
func maxWeightUsingDfs(weight []int, maxWeight int) int {
	var max, cur int
	dfsWeight(weight, maxWeight, 0, cur, &max)
	return max
}

func dfsWeight(weight []int, maxWeight int, idx, cur int, max *int) {
	if idx >= len(weight) {
		return
	}

	// 不添加
	dfsWeight(weight, maxWeight, idx+1, cur, max)

	// 添加
	tmp := cur + weight[idx]
	if tmp > maxWeight {
		return
	}
	if tmp > *max {
		*max = tmp
	}
	if *max == maxWeight {
		// 已达到最优结果
		return
	}
	dfsWeight(weight, maxWeight, idx+1, tmp, max)
}

// maxWeightUsingMemoDfs maxWeightUsingDfs使用备忘录的版本
func maxWeightUsingMemoDfs(weight []int, maxWeight int) int {
	memo := make([][]bool, len(weight))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]bool, maxWeight+1)
	}

	var max, cur int
	dfsWeightWithMemo(weight, maxWeight, memo, 0, cur, &max)
	return max
}

func dfsWeightWithMemo(weight []int, maxWeight int, memo [][]bool, idx, cur int, max *int) {
	if idx >= len(weight) {
		return
	}

	if memo[idx][cur] {
		return
	}
	memo[idx][cur] = true

	// 不添加
	dfsWeightWithMemo(weight, maxWeight, memo, idx+1, cur, max)

	// 添加
	tmp := cur + weight[idx]
	if tmp > maxWeight {
		return
	}
	if tmp > *max {
		*max = tmp
	}
	if *max == maxWeight {
		// 已达到最优结果
		return
	}
	dfsWeight(weight, maxWeight, idx+1, tmp, max)
}

// maxWeightUsingDynamicPrograming 使用动态规划解决0/1背包问题
func maxWeightUsingDynamicPrograming(weight []int, maxWeight int) int {
	state := make([][]bool, len(weight))
	for i := 0; i < len(state); i++ {
		state[i] = make([]bool, maxWeight+1)
	}
	state[0][0] = true
	if weight[0] <= maxWeight {
		state[0][weight[0]] = true
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= maxWeight; j++ {
			if !state[i-1][j] {
				continue
			}

			// 不放
			state[i][j] = state[i-1][j]

			// 放
			w := j + weight[i]
			if w > maxWeight {
				continue
			}
			state[i][w] = true
		}
	}

	for i := maxWeight; i >= 0; i-- {
		if state[len(weight)-1][i] {
			return i
		}
	}

	return 0
}

// maxWeightUsingDynamicProgramingWithLessMemory 使用一维数组的方式实现动态规划解决方案
func maxWeightUsingDynamicProgramingWithLessMemory(weight []int, maxWeight int) int {
	state := make([]bool, maxWeight+1)
	state[0] = true
	if weight[0] <= maxWeight {
		state[weight[0]] = true
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= maxWeight; j++ {
			if !state[j] {
				continue
			}

			// 不放，不用改变原有状态

			// 放
			if j+weight[i] <= maxWeight {
				state[j+weight[i]] = true
			}
		}
	}

	for i := maxWeight; i >= 0; i-- {
		if state[i] {
			return i
		}
	}

	return 0
}

// maxValueUsingDfs 使用深度优先搜索实现价值最大化
func maxValueUsingDfs(weight, value []int, maxWeight int) int {
	var idx, max int
	dfsValue(weight, value, maxWeight, idx, 0, 0, &max)
	return max
}

func dfsValue(weight []int, value []int, maxWeight int, idx int, curWeight, curValue int, max *int) {
	if idx >= len(weight) {
		return
	}

	// 不放
	dfsValue(weight, value, maxWeight, idx+1, curWeight, curValue, max)

	// 放
	if curWeight+weight[idx] <= maxWeight {
		if curValue+value[idx] > *max {
			*max = curValue + value[idx]
		}
		dfsValue(weight, value, maxWeight, idx+1, curWeight+weight[idx], curValue+value[idx], max)
	}
}

// maxValueUsingDynamicPrograming 使用动态规划解决价值最大化的0/1背包问题
func maxValueUsingDynamicPrograming(weight, value []int, maxWeight int) int {
	state := make([]int, maxWeight+1)
	for i := 0; i < len(state); i++ {
		state[i] = -1
	}

	state[0] = 0
	if weight[0] <= maxWeight {
		state[weight[0]] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= maxWeight; j++ {
			if state[j] < 0 {
				continue
			}

			if j+weight[i] <= maxWeight {
				state[j+weight[i]] = state[j] + value[i]
			}
		}
	}

	var maxValue int
	for i := 0; i <= maxWeight; i++ {
		if state[i] > maxValue {
			maxValue = state[i]
		}
	}

	return maxValue
}

// minimumLimitValueUsingDynamicPrograming value代表各个商品的价值，索引为商品编号。limit为最低价值要求。
// 求出最近接近limit要求的商品组合(如果超过limit的3倍，则不在考虑范围内)
func minimumLimitValueUsingDynamicPrograming(value []int, limit int) []int {
	if len(value) <= 0 {
		return nil
	}

	state := make([][]bool, len(value))
	max := limit * 3
	state[0] = make([]bool, max)
	state[0][0] = true
	if value[0] < max {
		state[0][value[0]] = true
	}
	for i := 1; i < len(value); i++ {
		state[i] = make([]bool, max)
		for j := 0; j < max; j++ {
			if !state[i-1][j] {
				continue
			}
			state[i][j] = true

			if value[i]+j >= max {
				continue
			}

			state[i][value[i]+j] = true
		}
	}

	var target int
	for i := limit; i < max; i++ {
		if state[len(value)-1][i] {
			target = i
			break
		}
	}
	if target <= 0 {
		return nil
	}

	var items []int
	for i := len(value) - 1; i > 0; i-- {
		if target-value[i] >= 0 && state[i-1][target-value[i]] {
			items = append(items, value[i])
			target -= value[i]
		}
	}
	if target > 0 {
		items = append(items, value[0])
	}

	return items
}

// yanghuiTriangleUsingDynamicProgram 给出一个杨辉三角，节点值为该路径的长度。
// 要求从顶点出发，求出到达最底层的最短路径长度。
// image: ../images/yanghui_triangle.png
func yanghuiTriangleUsingDynamicProgram(matrix [][]int) int {
	if len(matrix) <= 0 {
		return 0
	}

	state := make([][]int, len(matrix))
	state[0] = make([]int, 1)
	state[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		state[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				// 首个元素只能从上一个节点走过来
				state[i][j] = state[i-1][j] + matrix[i][j]
			} else if j == len(matrix[i])-1 {
				// 最后一个元素只能沿着最后的边走过来
				state[i][j] = state[i-1][j-1] + matrix[i][j]
			} else {
				// 中间节点可以从两个方先过来，既然都可以达到这个节点，当然选择路径最短的
				state[i][j] = min(state[i-1][j-1], state[i-1][j]) + matrix[i][j]
			}
		}
	}

	minDistance := math.MaxInt
	for i := 0; i < len(state[len(state)-1]); i++ {
		if state[len(state)-1][i] < minDistance {
			minDistance = state[len(state)-1][i]
		}
	}

	return minDistance
}

// shortestPathInMatrixUsingDynamicProgram 给出一个矩阵，每个值代表移动到该位置的路径长度。
// 要求从左上角移动到右下角，给出移动的最短路径长度
func shortestPathInMatrixUsingDynamicProgram(matrix [][]int) int {
	if len(matrix) <= 0 {
		return 0
	}

	state := make([][]int, len(matrix))
	state[0] = make([]int, len(matrix[0]))
	state[0][0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		// 第一行元素上面没有元素，做特殊处理
		state[0][i] = state[0][i-1] + matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		state[i] = make([]int, len(matrix[i]))
		// 首列元素没有左边的元素，做特殊处理
		state[i][0] = state[i-1][0] + matrix[i][0]

		for j := 1; j < len(matrix[i]); j++ {
			state[i][j] = min(state[i-1][j], state[i][j-1]) + matrix[i][j]
		}
	}

	return state[len(state)-1][len(state[len(state)-1])-1]
}
