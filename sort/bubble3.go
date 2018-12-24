package sort

// 时间复杂度O(n^2)
// 最好的情况: 如果已经排好序了，只需要进行n-1次比较，时间复杂度为O(n)
// 最差的情况: 1 + 2 +3 + ... + (n-1), 等于: n(n-1)/2，时间复杂度为O(n)
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
