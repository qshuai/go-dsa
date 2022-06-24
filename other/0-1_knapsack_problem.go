package other

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
