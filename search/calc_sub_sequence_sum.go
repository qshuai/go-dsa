package search

// CalcMaxSumForSubSequence 寻找一个序列的最大连续子序列的和（关键字： 连续， 最大和）
// eg:
//  [1, 2, 3, 4] => 1 = 2 + 3 + 4 = 10
//  [1, 2, 3, -40, 40] => 40 = 40
//  [10, -2, 12, -30, 2] => 10 - 2 + 12 = 20
func CalcMaxSumForSubSequence(arr []int) int {
	if arr == nil {
		return 0
	}

	sum, prevSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if prevSum < 0 {
			prevSum = arr[i]
			if prevSum > sum {
				sum = prevSum
			}
			continue
		}

		prevSum += arr[i]
		if prevSum > sum {
			sum = prevSum
		}
	}

	return sum
}
