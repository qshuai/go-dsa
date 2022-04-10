package sort

// 原理: 每个数都要和其他的所有数比较一次，如果不符合排序规则,这交换两个元素位置
// 时间复杂度O(n^2)
func simpleBubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

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

// 时间复杂度O(n^2)
// 最好的情况: 如果已经排好序了，只需要进行n-1次比较，时间复杂度为O(n)
// 最差的情况: 1 + 2 + 3 + ... + (n-1), 等于: n(n-1)/2，时间复杂度为O(n^2)
func optimizedBubbleSort(array []int) []int {
	// 每次外循环之后的结果是最小的数往前移动
	var sorted bool
	for i := 0; i < len(array) && !sorted; i++ {
		// 假设已经排序完成了
		sorted = true
		// 内循环是从后往前，小的往前移动
		for j := len(array) - 2; j >= i; j-- {
			if array[j] > array[j+1] {
				// 如果有元素交换，这说明可能还没有完全排好序
				sorted = false
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
