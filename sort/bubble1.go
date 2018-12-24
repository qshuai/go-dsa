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
