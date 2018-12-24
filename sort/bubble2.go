package sort

// 时间复杂度O(n^2)
func mediumBubbleSort(array []int) []int {
	// 每次外循环之后的结果是最小的数往前移动
	for i := 0; i < len(array); i++ {
		// 内循环是从后往前，小的往前移动
		for j := len(array) - 2; j >= i; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
