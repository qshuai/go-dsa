package sort

// 时间复杂度O(n^2)
// 这个算法的优点是避免了不必要的元素交换
func selectSort(array []int) []int {
	var min int

	for i := 0; i < len(array); i++ {
		// 重置min
		min = i

		// 找到后面最小的元素，并记录下其索引
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}

		// 如果发现了一个小值元素，则进行交换
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}

	return array
}
