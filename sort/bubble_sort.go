package sort

// bubbleSort 冒泡排序（稳定排序；时间复杂度O(n^2)）
func bubbleSort(arr []int) []int {
	var sorted bool
	for i := 0; i < len(arr)-1 && !sorted; i++ {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}
